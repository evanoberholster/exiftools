package exiftool

import (
	"errors"

	"github.com/evanoberholster/exiftools/exiftool/exif"
)

const (
	// IFD1

	ThumbnailOffsetTagId = 0x0201
	ThumbnailSizeTagId   = 0x0202
)

var (
	ErrIfdNotFound = errors.New("Ifd Not Found")
	ErrTagNotFound = errors.New("Tag Not Found")
)

type TagIndex struct {
	//tagFn      map[string]map[exif.TagID]func()
	tagsByIfd map[string]map[exif.TagID]exif.Tag
	//tagsByIfd  map[string]map[uint16]*IndexedTag
	//tagsByIfdR map[string]map[string]*IndexedTag
}

type IndexedTag struct {
	Id   exif.TagID
	Name string
	//IfdPath string
	Type exif.TagType
}

func NewTagIndex() *TagIndex {
	ti := new(TagIndex)
	// map[IFD]map[TagID]IndexedTag
	ti.tagsByIfd = make(map[string]map[exif.TagID]exif.Tag)
	//ti.tagFn = make(map[string]map[exif.TagID]func())
	//ti.tagsByIfd = make(map[string]map[uint16]*IndexedTag)
	//ti.tagsByIfdR = make(map[string]map[string]*IndexedTag)
	return ti
}

func (ti *TagIndex) Add(ifdPath string, tagMap exif.TagMap) {
	ti.tagsByIfd[ifdPath] = tagMap
}

// Get returns information about the non-IFD tag.
func (ti *TagIndex) Get(ifdPath string, tagID exif.TagID) (tag exif.Tag, err error) {
	defer func() {
		if state := recover(); state != nil {
			err = state.(error)
		}
	}()

	//if len(ti.tagsByIfd) == 0 {
	//	err := LoadStandardTags(ti)
	//	log.PanicIf(err)
	//}

	family, ok := ti.tagsByIfd[ifdPath]
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
