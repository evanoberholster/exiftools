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
func (eh ExifHeader) Visit(rootIfdPointer exif.IfdPath, rootIfdName string, ifdMapping *IfdMapping, tagIndex *TagIndex, exifData []byte, visitor TagVisitorFn) (err error) {
	defer func() {
		if state := recover(); state != nil {
			err = state.(error)
		}
	}()
	// check to make sure that rawExif is valid
	if !eh.Valid() {
		return ErrExifNotValid
	}

	ie := NewIfdEnumerate(ifdMapping, tagIndex, exifData, eh.ByteOrder)

	return ie.Scan(rootIfdName, eh.FirstIfdOffset, visitor)
}

type ExifTagDecodeFn func(fqIfdPath string, ifdIndex int, ite *IfdTagEntry) (err error)

func (eh *ExifHeader) Decode(rootIfdName string, ifdMapping *IfdMapping, tagIndex *TagIndex, data []byte, decodeFn TagVisitorFn) (err error) {
	defer func() {
		if state := recover(); state != nil {
			err = state.(error)
		}
	}()
	// check to make sure that the ExifHeader is valid
	if !eh.Valid() {
		return ErrExifNotValid
	}
	// seek to begining of exif data
	//reader.Seek(eh.FirstIfdOffset, 0)
	er := NewExifReader(eh, data)
	//reader.Seek(int64(eh.foundAt)+4, 0)
	ie := newIfdEnumerate(er, ifdMapping, tagIndex, eh.ByteOrder)

	return ie.scan2(rootIfdName, eh.FirstIfdOffset, decodeFn)
}
