package exif

import (
	"encoding/binary"
	"fmt"
)

// ValueContext embeds all of the parameters required to find and extract the
// actual tag value.
type ValueContext struct {
	unitCount       uint32
	valueOffset     uint32
	rawValueOffset  []byte
	addressableData []byte

	tagType   TagType
	byteOrder binary.ByteOrder

	// undefinedValueTagType is the effective type to use if this is an
	// "undefined" value.
	undefinedValueTagType TagType

	ifdPath string
	tagID   TagID
}

// NewValueContext returns a new ValueContext struct.
func NewValueContext(ifdPath string, tagID TagID, unitCount, valueOffset uint32, rawValueOffset, addressableData []byte, tagType TagType, byteOrder binary.ByteOrder) *ValueContext {
	return &ValueContext{
		unitCount:       unitCount,
		valueOffset:     valueOffset,
		rawValueOffset:  rawValueOffset,
		addressableData: addressableData,

		tagType:   tagType,
		byteOrder: byteOrder,

		ifdPath: ifdPath,
		tagID:   tagID,
	}
}

// effectiveValueType returns the effective type of the unknown-type tag or, if
// not unknown, the actual type.
func (vc *ValueContext) effectiveValueType() (tagType TagType) {
	if vc.tagType == TypeUndefined {
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
	case TypeByte:
		return vc.ReadBytes()
	case TypeASCII:
		return vc.ReadASCII()
	case TypeASCIINoNul:
		return vc.ReadASCIINoNul()
	case TypeLong:
		return vc.ReadLongs()
	case TypeShort:
		return vc.ReadShorts()
	case TypeRational:
		return vc.ReadRationals()
	case TypeSignedLong:
		return vc.ReadSignedLongs()
	case TypeSignedRational:
		return vc.ReadSignedRationals()
	case TypeUndefined:
		return nil, fmt.Errorf("Will not parse undefined-type value")
	default:
		return nil, fmt.Errorf("Value of type [%s] is unparseable", vc.tagType)
	}

}
