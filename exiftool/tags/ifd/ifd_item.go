package ifd

import "github.com/evanoberholster/exiftools/exiftool/exif"

// IfdPath is an array of TagID representing an IFD
type IfdPath []exif.TagID

// IfdItem is a container to define each IFD
type IfdItem struct {
	IfdPath IfdPath
	TagID   exif.TagID
	Name    string
}

// Valid returns true if the IfdItem has a Name
func (ifd IfdItem) Valid() bool {
	return ifd.Name != ""
}
