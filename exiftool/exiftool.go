package exiftool

import (
	"errors"

	"github.com/evanoberholster/exiftools/exiftool/exif"
)

var ifdTagNames = make(map[string]TagNameMap)

type TagNameMap map[exif.TagID]string

func (tnm TagNameMap) Name(tagID exif.TagID) string {
	return tnm[tagID]
}

// Options
const (
	MaxExifSize int64 = 4 * 1024 * 1024
)

// Errors
var (
	ErrNoExif       error = errors.New("Error no Exif")
	ErrExifNotValid       = errors.New("Error Exif Not Valid")
)

// TagVisitorFn is called for each tag when enumerating through the EXIF.
type TagVisitorFn func(fqIfdPath string, ifdIndex int, ite *IfdTagEntry) (err error)

// Visit recursively invokes a callback for every tag.
func (er ExifReader) Visit(rootIfdName string, ifdMapping *IfdMapping, tagIndex *TagIndex, visitor TagVisitorFn) (err error) {
	defer func() {
		if state := recover(); state != nil {
			err = state.(error)
		}
	}()
	// check to make sure that the ExifReader is valid
	//if !eh.Valid() {
	//	return ErrExifNotValid
	//}

	ie := er.getIfdEnumerate(ifdMapping, tagIndex)

	return ie.Scan(rootIfdName, er.firstIfdOffset, visitor)
}
