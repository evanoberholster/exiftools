package exiftool

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"io/ioutil"
)

// Valid returns true for a valid ExifHeader
func (eh ExifHeader) Valid() bool {
	return eh.foundAt >= 0
}

// ExifHeader - The Exif header
type ExifHeader struct {
	ByteOrder      binary.ByteOrder
	FirstIfdOffset uint32
	foundAt        int
}

// ParseExif parses a an io.ReadSeeker and returns an ExifHeader. ErrNoExif is returned
// as an error if No Exif byte Header is found.
func ParseExif(r io.ReadSeeker) (eh *ExifHeader, err error) {
	defer func() {
		if state := recover(); state != nil {
			err = fmt.Errorf("Exif Error: %v", state.(error))
		}
	}()

	eh, err = parseExifHeader(r)
	if !eh.Valid() {
		return nil, ErrNoExif
	}
	return eh, nil
}

// ParseExif2 parses possible Exif from an io.Reader
func ParseExif2(r io.ReadSeeker) (eh *ExifHeader, err error) {
	defer func() {
		if state := recover(); state != nil {
			err = fmt.Errorf("Exif Error: %v", err)
		}
	}()

	eh, err = parseExifHeader2(r)
	return eh, err
}

// ParseExifHeader parses the bytes at the very top of the header.
// Benchmarked
func parseExifHeader(r io.Reader) (*ExifHeader, error) {
	// Good reference:
	//
	//      CIPA DC-008-2016; JEITA CP-3451D
	//      -> http://www.cipa.jp/std/documents/e/DC-008-Translation-2016-E.pdf

	bufSize := 1024 // 1kB Buffer
	data := make([]byte, bufSize)
	var err error
	var a, j int
	for {
		// Header length is 8bytes
		// Search for the beginning of the EXIF information. The EXIF is near the
		// beginning of our/most JPEGs, so this has a very low cost if Exif information is present
		copy(data[:8], data[len(data)-8:])
		a, err = r.Read(data[8:])
		if err == io.EOF {
			panic(ErrNoExif)
			//return &ExifHeader{foundAt: -1}, ErrNoExif
		} else if err != nil {
			panic(err)
		}
		j += a
		for i := 0; i < len(data)-8; i++ {

			// If BigEndian Header is found
			if bytes.Equal([]byte{data[i], data[i+1], data[i+2], data[i+3]}, bigEndianHeaderBytes[:]) {
				return &ExifHeader{
					ByteOrder:      binary.BigEndian,
					FirstIfdOffset: binary.BigEndian.Uint32(data[i+4 : i+8]),
					foundAt:        (j + i) - bufSize,
				}, nil
			}

			// If LittleEndian Header is found
			if bytes.Equal([]byte{data[i], data[i+1], data[i+2], data[i+3]}, littleEndianHeaderBytes[:]) {
				return &ExifHeader{
					ByteOrder:      binary.LittleEndian,
					FirstIfdOffset: binary.LittleEndian.Uint32(data[i+4 : i+8]),
					foundAt:        (j + i) - bufSize,
				}, nil
			}
		}
	}
}

// Search for the beginning of the EXIF information. The EXIF is near the
// beginning of our/most JPEGs, so this has a very low cost.

// ParseExifHeader parses the bytes at the very top of the header.
func parseExifHeader2(r io.Reader) (*ExifHeader, error) {
	eh := &ExifHeader{
		foundAt: -1,
	}

	data, err := ioutil.ReadAll(r)
	if err != nil {
		return eh, err
	}
	for i := 0; i < len(data)-8; i++ {
		if bytes.Equal([]byte{data[i], data[i+1], data[i+2], data[i+3]}, bigEndianHeaderBytes[:]) {
			firstIfdOffset := binary.BigEndian.Uint32(data[i+4 : i+8])
			return &ExifHeader{
				ByteOrder:      binary.BigEndian,
				FirstIfdOffset: firstIfdOffset,
				foundAt:        i,
			}, nil
		}
		if bytes.Equal([]byte{data[i], data[i+1], data[i+2], data[i+3]}, littleEndianHeaderBytes[:]) {
			firstIfdOffset := binary.LittleEndian.Uint32(data[i+4 : i+8])
			return &ExifHeader{
				ByteOrder:      binary.LittleEndian,
				FirstIfdOffset: firstIfdOffset,
				foundAt:        i,
			}, nil
		}
	}

	return eh, ErrNoExif
}

var (
	bigEndianHeaderBytes    = [4]byte{'M', 'M', 0x00, 0x2a}
	littleEndianHeaderBytes = [4]byte{'I', 'I', 0x2a, 0x00}
)
