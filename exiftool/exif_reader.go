package exiftool

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
)

// Reader interface
type Reader interface {
	io.ReaderAt
	io.Reader
}

// ExifReader -
type ExifReader struct {
	reader Reader

	// previously ExifHeader
	byteOrder      binary.ByteOrder
	exifOffset     int64
	firstIfdOffset uint32
	//
	index  int64
	offset int64
}

func (er *ExifReader) ByteOrder() binary.ByteOrder {
	return er.byteOrder
}

// SubReader returns an ExifReader with its offest set to IfdOffset and
// an empty index
func (er *ExifReader) SubReader(ifdOffset int64) *ExifReader {
	return &ExifReader{
		reader:     er.reader,
		byteOrder:  er.byteOrder,
		exifOffset: er.exifOffset,
		offset:     ifdOffset,
	}
}

// Read reads from ExifReafer and moves the offset marker
func (er *ExifReader) Read(p []byte) (n int, err error) {
	// Buffer is empty
	if len(p) == 0 {
		return 0, nil
	}
	n, err = er.reader.ReadAt(p, er.exifOffset+er.offset)
	er.offset += int64(n)

	return n, err
}

// ReadAt reads from ExifReader at the given offset
func (er *ExifReader) ReadAt(p []byte, off int64) (n int, err error) {
	if off < 0 {
		return 0, errors.New("ExifReader.ReadAt: negative offset")
	}
	return er.reader.ReadAt(p, er.exifOffset+off)
}

// NewExifReader returns a new ExifReader. It reads from reader according to byteOrder from exifOffset
func NewExifReader(reader Reader, byteOrder binary.ByteOrder, firstIfdOffset uint32, exifOffset int64) *ExifReader {
	return &ExifReader{
		reader:         reader,
		byteOrder:      byteOrder,
		exifOffset:     exifOffset,
		firstIfdOffset: firstIfdOffset,
	}
}

// NewExifReader creates a New ExifReader from an ExifHeader with the following reader
func (eh ExifHeader) NewExifReader(reader Reader) (*ExifReader, error) {
	if !eh.Valid() {
		return nil, ErrExifNotValid
	}
	return &ExifReader{
		reader:         reader,
		byteOrder:      eh.byteOrder,
		exifOffset:     eh.exifOffset,
		firstIfdOffset: eh.firstIfdOffset,
	}, nil
}

// Valid returns true for a valid ExifHeader
func (eh ExifHeader) Valid() bool {
	return eh.exifOffset >= 0
}

// ExifHeader is an Exif Header
type ExifHeader struct {
	byteOrder      binary.ByteOrder
	firstIfdOffset uint32
	exifOffset     int64
}

// ParseExif2 parses a an io.Reader and returns an ExifReader. ErrNoExif is returned
// as an error if No Exif byte Header is found.
func ParseExif2(r Reader) (er *ExifReader, err error) {
	defer func() {
		if state := recover(); state != nil {
			err = fmt.Errorf("Exif Error: %v", state.(error))
		}
	}()

	eh, err := parseExifHeader2(r)
	if err != nil || !eh.Valid() {
		return nil, ErrNoExif
	}
	return eh.NewExifReader(r)
}

// ParseExifHeader2 parses the bytes at the very top of the header.
// Benchmarked
func parseExifHeader2(r io.Reader) (*ExifHeader, error) {
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
			return nil, ErrNoExif
		} else if err != nil {
			panic(err)
		}
		j += a
		for i := 0; i < len(data)-8; i++ {

			// If BigEndian Header is found
			if bytes.Equal([]byte{data[i], data[i+1], data[i+2], data[i+3]}, bigEndianHeaderBytes[:]) {
				return &ExifHeader{
					byteOrder:      binary.BigEndian,
					firstIfdOffset: binary.BigEndian.Uint32(data[i+4 : i+8]),
					exifOffset:     int64((j + i) - bufSize),
				}, nil
			}

			// If LittleEndian Header is found
			if bytes.Equal([]byte{data[i], data[i+1], data[i+2], data[i+3]}, littleEndianHeaderBytes[:]) {
				return &ExifHeader{
					byteOrder:      binary.LittleEndian,
					firstIfdOffset: binary.LittleEndian.Uint32(data[i+4 : i+8]),
					exifOffset:     int64((j + i) - bufSize),
				}, nil
			}
		}
	}
}

// ParseExif parses a an io.Reader and returns an ExifHeader. ErrNoExif is returned
// as an error if No Exif byte Header is found.
func ParseExif(r io.Reader) (eh *ExifHeader, err error) {
	defer func() {
		if state := recover(); state != nil {
			err = fmt.Errorf("Exif Error: %v", state.(error))
		}
	}()

	eh, err = parseExifHeader(r)
	if err != nil || !eh.Valid() {
		return nil, ErrNoExif
	}
	return eh, nil
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
			return nil, ErrNoExif
		} else if err != nil {
			panic(err)
		}
		j += a
		for i := 0; i < len(data)-8; i++ {

			// If BigEndian Header is found
			if bytes.Equal([]byte{data[i], data[i+1], data[i+2], data[i+3]}, bigEndianHeaderBytes[:]) {
				return &ExifHeader{
					byteOrder:      binary.BigEndian,
					firstIfdOffset: binary.BigEndian.Uint32(data[i+4 : i+8]),
					exifOffset:     int64((j + i) - bufSize),
				}, nil
			}

			// If LittleEndian Header is found
			if bytes.Equal([]byte{data[i], data[i+1], data[i+2], data[i+3]}, littleEndianHeaderBytes[:]) {
				return &ExifHeader{
					byteOrder:      binary.LittleEndian,
					firstIfdOffset: binary.LittleEndian.Uint32(data[i+4 : i+8]),
					exifOffset:     int64((j + i) - bufSize),
				}, nil
			}
		}
	}
}

var (
	bigEndianHeaderBytes    = [4]byte{'M', 'M', 0x00, 0x2a}
	littleEndianHeaderBytes = [4]byte{'I', 'I', 0x2a, 0x00}
)
