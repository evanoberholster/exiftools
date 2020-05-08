package exiftool

import "github.com/evanoberholster/exiftools/exiftool/exif"

// IfdItem is a container to define each IFD
type IfdItem struct {
	IfdPath exif.IfdPath
	TagID   exif.TagID
	Name    string
}

// Valid returns true if the IfdItem has a Name
func (ifd IfdItem) Valid() bool {
	return ifd.Name != ""
}

// NewIfdMapping creates an empty IfdMapping from the rootNode
func NewIfdMapping() (ifdMapping *IfdMapping) {
	rootNode := &MappedIfd{
		Path:     make([]string, 0),
		Children: make(map[exif.TagID]*MappedIfd),
	}

	return &IfdMapping{
		rootNode: rootNode,
	}
}

type MappedIfd struct {
	ParentTagID exif.TagID
	Placement   exif.IfdPath
	Path        []string

	Name     string
	TagID    exif.TagID
	Children map[exif.TagID]*MappedIfd
}
