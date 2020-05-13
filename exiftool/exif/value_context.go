package exif

import (
	"encoding/binary"
	"io"
)

// ValueContext embeds all of the parameters required to find and extract the
// actual tag value.
type ValueContext struct {
	unitCount      uint32
	valueOffset    uint32
	rawValueOffset []byte

	exifReader io.ReaderAt
	//addressableData []byte

	tagType   TagType
	byteOrder binary.ByteOrder

	// undefinedValueTagType is the effective type to use if this is an
	// "undefined" value.
	undefinedValueTagType TagType

	ifdPath string
	tagID   TagID
}

// NewValueContext returns a new ValueContext struct.
//func NewValueContext(ifdPath string, tagID TagID, unitCount, valueOffset uint32, rawValueOffset []byte, exifReader io.ReaderAt, tagType TagType, byteOrder binary.ByteOrder) *ValueContext {
//	return &ValueContext{
//		unitCount:      unitCount,
//		valueOffset:    valueOffset,
//		rawValueOffset: rawValueOffset,
//		exifReader:     exifReader,
//
//		tagType:   tagType,
//		byteOrder: byteOrder,
//
//		ifdPath: ifdPath,
//		tagID:   tagID,
//	}
//}
