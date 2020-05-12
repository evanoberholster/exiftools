package exif

import (
	"fmt"
)

// TagType -
type TagType uint8

// TagTypes defined
// Copied from dsoprea/go-exif
const (
	// TypeByte describes an encoded list of bytes.
	TypeByte TagType = 1

	// TypeASCII describes an encoded list of characters that is terminated
	// with a NUL in its encoded form.
	TypeASCII TagType = 2

	// TypeShort describes an encoded list of shorts.
	TypeShort TagType = 3

	// TypeLong describes an encoded list of longs.
	TypeLong TagType = 4

	// TypeRational describes an encoded list of rationals.
	TypeRational TagType = 5

	// TypeUndefined describes an encoded value that has a complex/non-clearcut
	// interpretation.
	TypeUndefined TagType = 7

	// TypeSignedLong describes an encoded list of signed longs.
	TypeSignedLong TagType = 9

	// TypeSignedRational describes an encoded list of signed rationals.
	TypeSignedRational TagType = 10

	// TypeASCIINoNul is just a pseudo-type, for our own purposes.
	TypeASCIINoNul TagType = 0xf0
)

// IsValid returns true if tagType is a valid type.
func (tagType TagType) IsValid() bool {
	return tagType == TypeByte ||
		tagType == TypeASCII ||
		tagType == TypeASCIINoNul ||
		tagType == TypeShort ||
		tagType == TypeLong ||
		tagType == TypeRational ||
		tagType == TypeSignedLong ||
		tagType == TypeSignedRational ||
		tagType == TypeUndefined
}

// Tag sizes
const (
	TypeByteSize           = 1
	TypeASCIISize          = 1
	TypeASCIINoNulSize     = 1
	TypeShortSize          = 2
	TypeLongSize           = 4
	TypeRationalSize       = 8
	TypeSignedLongSize     = 4
	TypeSignedRationalSize = 8
)

// Size returns the size of one atomic unit of the type.
func (tagType TagType) Size() int {
	switch tagType {
	case TypeByte:
		return TypeByteSize
	case TypeASCII, TypeASCIINoNul:
		return TypeASCIISize
	case TypeShort:
		return TypeShortSize
	case TypeLong:
		return TypeLongSize
	case TypeRational:
		return TypeRationalSize
	case TypeSignedLong:
		return TypeSignedLongSize
	case TypeSignedRational:
		return TypeSignedRationalSize
	default:
		panic(fmt.Errorf("Can not determine tag-value size for type (%d): [%s]", tagType, tagType.String()))
		return 0
	}
}

// String returns the name of the Tag Type
func (tagType TagType) String() string {
	return tagTypeNames[tagType]
}

var (
	tagTypeNames = map[TagType]string{
		TypeByte:           "BYTE",
		TypeASCII:          "ASCII",
		TypeShort:          "SHORT",
		TypeLong:           "LONG",
		TypeRational:       "RATIONAL",
		TypeUndefined:      "UNDEFINED",
		TypeSignedLong:     "SLONG",
		TypeSignedRational: "SRATIONAL",

		TypeASCIINoNul: "_ASCII_NO_NUL",
	}

	tagTypeNamesR = map[string]TagType{}
)

// Rational -
type Rational struct {
	Numerator   uint32
	Denominator uint32
}

// SignedRational -
type SignedRational struct {
	Numerator   int32
	Denominator int32
}
