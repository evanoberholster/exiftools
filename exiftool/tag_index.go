package exiftool

import "github.com/evanoberholster/exiftools/exiftool/exif"

const (
	// IFD1

	ThumbnailOffsetTagId = 0x0201
	ThumbnailSizeTagId   = 0x0202
)

type TagIndex struct {
	//tagFn      map[string]map[exif.TagID]func()
	tagsByID map[string]map[exif.TagID]IndexedTag
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
	ti.tagsByID = make(map[string]map[exif.TagID]IndexedTag)
	//ti.tagFn = make(map[string]map[exif.TagID]func())
	//ti.tagsByIfd = make(map[string]map[uint16]*IndexedTag)
	//ti.tagsByIfdR = make(map[string]map[string]*IndexedTag)

	return ti
}

func (ti *TagIndex) Add(path string, ifdMap map[exif.TagID]IndexedTag) {
	ti.tagsByID[path] = ifdMap
}
