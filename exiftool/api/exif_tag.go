package api

import (
	"errors"
	"fmt"

	"github.com/evanoberholster/exiftools/exiftool/exif"
)

// API Errors
var (
	ErrEmptyTag = errors.New("Error empty tag")
	ErrParseTag = fmt.Errorf("Error parsing tag")
	ErrTagType  = errors.New("Error wrong tag Type")

	// ErrGpsCoordsNotValid means that some part of the geographic data were unparseable.
	ErrGpsCoordsNotValid = errors.New("GPS coordinates not valid")
)

// ExifTag -
type ExifTag struct {
	TagID   exif.TagID
	tagName string
	tagType exif.TagType

	// Needed to create ValueContext

	// Old
	value interface{}
}

// String returns an ExifTag as a string
func (et *ExifTag) String() (string, error) {
	if et == nil {
		return "", ErrEmptyTag
	}
	switch et.tagType {
	case exif.TypeByte:
		return string(et.value.([]byte)), nil
	case exif.TypeASCII, exif.TypeASCIINoNul:
		return et.value.(string), nil
	case exif.TypeRational:
		r := et.value.([]exif.Rational)
		return fmt.Sprintf("%d/%d", r[0].Numerator, r[0].Denominator), nil
	case exif.TypeSignedRational:
		r := et.value.([]exif.SignedRational)
		return fmt.Sprintf("%d/%d", r[0].Numerator, r[0].Denominator), nil
	}
	return "", ErrTagType
}

// Rational returns an ExifTag as a []Rational
// returns ErrTagType if ExifTag is not of TypeRational
func (et *ExifTag) Rational() ([]exif.Rational, error) {
	if et == nil {
		return nil, ErrEmptyTag
	}
	switch et.tagType {
	case exif.TypeRational:
		return et.value.([]exif.Rational), nil
	}
	return nil, ErrTagType
}

// Short returns an ExifTag as a []uint16
// returns ErrTagType if ExifTag is not of TypeShort
func (et *ExifTag) Short() ([]uint16, error) {
	if et == nil {
		return nil, ErrEmptyTag
	}
	switch et.tagType {
	case exif.TypeShort:
		return et.value.([]uint16), nil
	}
	return nil, ErrTagType
}

// SignedRational returns an ExifTag as a []SignedRational
// returns ErrTagType if ExifTag is not of TypeSignedRational
func (et *ExifTag) SignedRational() ([]exif.SignedRational, error) {
	if et == nil {
		return nil, ErrEmptyTag
	}
	switch et.tagType {
	case exif.TypeSignedRational:
		return et.value.([]exif.SignedRational), nil
	}
	return nil, ErrTagType
}

// Int returns the first item of an ExifTag of TypeShort
// or TypeLong as an int.
// returns ErrTagType if ExifTag is not of TypeShort or TypeLong
func (et *ExifTag) Int() (int, error) {
	if et == nil {
		return 0, ErrEmptyTag
	}
	switch et.tagType {
	case exif.TypeShort:
		return int(et.value.([]uint16)[0]), nil
	case exif.TypeLong:
		return int(et.value.([]uint32)[0]), nil
	}
	return 0, ErrTagType
}
