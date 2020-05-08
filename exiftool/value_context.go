package exiftool

import (
	"encoding/binary"
	"fmt"

	"github.com/evanoberholster/exiftools/exiftool/exif"
)

// ValueContext embeds all of the parameters required to find and extract the
// actual tag value.
type ValueContext struct {
	unitCount      uint32
	valueOffset    uint32
	rawValueOffset []byte

	exifReader      *ExifReader
	addressableData []byte

	tagType   exif.TagType
	byteOrder binary.ByteOrder

	// undefinedValueTagType is the effective type to use if this is an
	// "undefined" value.
	undefinedValueTagType exif.TagType

	ifdPath string
	tagID   exif.TagID
}

// NewValueContext returns a new ValueContext struct.
func NewValueContext(ifdPath string, tagID exif.TagID, unitCount, valueOffset uint32, rawValueOffset []byte, exifReader *ExifReader, tagType exif.TagType, byteOrder binary.ByteOrder) *ValueContext {
	return &ValueContext{
		unitCount:      unitCount,
		valueOffset:    valueOffset,
		rawValueOffset: rawValueOffset,
		exifReader:     exifReader,

		tagType:   tagType,
		byteOrder: byteOrder,

		ifdPath: ifdPath,
		tagID:   tagID,
	}
}

// effectiveValueType returns the effective type of the unknown-type tag or, if
// not unknown, the actual type.
func (vc *ValueContext) effectiveValueType() (tagType exif.TagType) {
	if vc.tagType == exif.TypeUndefined {
		tagType = vc.undefinedValueTagType

		if tagType == 0 {
			panic(fmt.Errorf("undefined-value type not set"))
		}
	} else {
		tagType = vc.tagType
	}

	return tagType
}

// Values knows how to resolve the given value. This value is always a list
// (undefined-values aside), so we're named accordingly.
//
// Since this method lacks the information to process unknown-type tags (e.g.
// byte-order, tag-ID, IFD type), it will return an error if attempted. See
// `Undefined()`.
func (vc *ValueContext) Values() (values interface{}, err error) {

	switch vc.tagType {
	case exif.TypeByte:
		return vc.ReadBytes()
	case exif.TypeASCII:
		return vc.ReadASCII()
	case exif.TypeASCIINoNul:
		return vc.ReadASCIINoNul()
	case exif.TypeLong:
		return vc.ReadLongs()
	case exif.TypeShort:
		return vc.ReadShorts()
	case exif.TypeRational:
		return vc.ReadRationals()
	case exif.TypeSignedLong:
		return vc.ReadSignedLongs()
	case exif.TypeSignedRational:
		return vc.ReadSignedRationals()
	case exif.TypeUndefined:
		return nil, fmt.Errorf("Will not parse undefined-type value")
	default:
		return nil, fmt.Errorf("Value of type [%s] is unparseable", vc.tagType)
	}

}
