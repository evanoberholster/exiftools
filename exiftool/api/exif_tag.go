package api

import (
	"errors"
	"fmt"

	"github.com/evanoberholster/exiftools/exiftool/exif"
)

var (
	ErrEmptyTag = errors.New("Error empty tag")
)

// ExifTag -
type ExifTag struct {
	TagID   exif.TagID
	tagName string
	tagType exif.TagType
	value   interface{} `json:"value"`
}

// String
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
	return "", nil
}

// Rational -
func (et ExifTag) Rational() []exif.Rational {
	switch et.tagType {
	case exif.TypeRational:
		return et.value.([]exif.Rational)
	}
	return nil
}

// Int -
func (et *ExifTag) Int() (int, error) {
	if et == nil {
		return 0, ErrEmptyTag
	}
	switch et.tagType {
	case exif.TypeShort:
		return int(et.value.([]uint16)[0]), nil
	}
	return 0, nil
}
