package exiftool

import (
	"errors"

	"github.com/evanoberholster/exiftools/exiftool/exif"
)

const (
	// ThumbnailOffsetTagID is IFD0/IFD1 Thumbnail Offset TagID
	ThumbnailOffsetTagID = 0x0201
	// ThumbnailSizeTagID is IFD0/IFD1 Thumbnail Size TagID
	ThumbnailSizeTagID = 0x0202
)

// Errors
var (
	ErrIfdNotFound = errors.New("Ifd Not Found")
	ErrTagNotFound = errors.New("Tag Not Found")
)

// TagIndex contains Tags by IFD
type TagIndex map[string]map[exif.TagID]exif.Tag

// NewTagIndex - creates a new empty TagIndex
func NewTagIndex() TagIndex {
	return make(TagIndex)
}

// Add adds a TagMap to a fdIfdPath
func (ti TagIndex) Add(fqIfdPath string, tagMap exif.TagMap) {
	ti[fqIfdPath] = tagMap
}

// Get returns information about the non-IFD tag.
func (ti TagIndex) Get(fqIfdPath string, tagID exif.TagID) (tag exif.Tag, err error) {
	defer func() {
		if state := recover(); state != nil {
			err = state.(error)
		}
	}()

	family, ok := ti[fqIfdPath]
	if !ok {
		err = ErrIfdNotFound
		return
	}

	tag, ok = family[tagID]
	if !ok {
		err = ErrTagNotFound
		return
	}

	return tag, nil
}
