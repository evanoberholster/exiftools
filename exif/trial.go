package exif

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/evanoberholster/exif/tiff"
)

// DecodeWithParseHeader -
func DecodeWithParseHeader(r io.Reader) (x *Exif, err error) {
	defer func() {
		if state := recover(); state != nil {
			err = fmt.Errorf("Exif Error: %v", err)
		}
	}()
	r2 := io.LimitReader(r, int64(ExifLengthCutoff))
	data, _ := ioutil.ReadAll(r2)

	foundAt := -1
	for i := 0; i < len(data); i++ {
		if _, err = parseExifHeader(data[i:]); err == nil {
			foundAt = i
			break
		}
	}
	if err != nil {
		return
	}

	er := bytes.NewReader(data[foundAt:])
	tif, err := tiff.Decode(er)

	er.Seek(0, 0)
	raw, err := ioutil.ReadAll(er)
	if err != nil {
		return nil, decodeError{cause: err}
	}

	// build an exif structure from the tiff
	x = &Exif{
		main: map[FieldName]*tiff.Tag{},
		Tiff: tif,
		Raw:  raw,
	}

	for i, p := range parsers {
		if err := p.Parse(x); err != nil {
			if _, ok := err.(tiffErrors); ok {
				return x, err
			}
			// This should never happen, as Parse always returns a tiffError
			// for now, but that could change.
			return x, fmt.Errorf("exif: parser %v failed (%v)", i, err)
		}
	}

	return x, nil
}

// parseExifHeader -
func parseExifHeader(data []byte) (eh exifHeader, err error) {
	// Good reference:
	//
	//      CIPA DC-008-2016; JEITA CP-3451D
	//      -> http://www.cipa.jp/std/documents/e/DC-008-Translation-2016-E.pdf

	if len(data) < 2 {
		return eh, fmt.Errorf("Not enough data for EXIF header (1): (%d)", len(data))
	}

	byteOrderBytes := [2]byte{data[0], data[1]}

	byteOrder, found := byteOrderLookup[byteOrderBytes]
	if found == false {
		return eh, fmt.Errorf("EXIF byte-order not recognized: [%v]", byteOrderBytes)
	}

	if len(data) < 4 {
		return eh, fmt.Errorf("Not enough data for EXIF header (2): (%d)", len(data))
	}

	fixedBytes := [2]byte{data[2], data[3]}
	expectedFixedBytes := exifFixedBytesLookup[byteOrder]
	if fixedBytes != expectedFixedBytes {
		return eh, fmt.Errorf("EXIF header fixed-bytes should be [%v] but are: [%v]", expectedFixedBytes, fixedBytes)
	}

	if len(data) < 2 {
		return eh, fmt.Errorf("Not enough data for EXIF header (3): (%d)", len(data))
	}

	firstIfdOffset := byteOrder.Uint32(data[4:8])

	eh = exifHeader{
		ByteOrder:      byteOrder,
		FirstIfdOffset: firstIfdOffset,
	}

	return eh, nil
}

var (
	exifFixedBytesLookup = map[binary.ByteOrder][2]byte{
		binary.LittleEndian: [2]byte{0x2a, 0x00},
		binary.BigEndian:    [2]byte{0x00, 0x2a},
	}
	byteOrderLookup = map[[2]byte]binary.ByteOrder{
		bigEndianBoBytes:    binary.BigEndian,
		littleEndianBoBytes: binary.LittleEndian,
	}
	bigEndianBoBytes    = [2]byte{'M', 'M'}
	littleEndianBoBytes = [2]byte{'I', 'I'}
)

type exifHeader struct {
	ByteOrder      binary.ByteOrder
	FirstIfdOffset uint32
}
