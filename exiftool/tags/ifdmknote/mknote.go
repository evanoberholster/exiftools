package ifdmknote

import (
	"github.com/evanoberholster/exiftools/exiftool/exif"
	"github.com/evanoberholster/exiftools/exiftool/tags/ifd"
	"github.com/evanoberholster/exiftools/exiftool/tags/ifdexif"
)

// Makernote Name and TagID
const (
	IfdMakernote              = "Makernotes"
	IfdMakernoteID exif.TagID = 0x927c
)

// MakernotePath is the MakernotePath Ifd Path
var (
	MakernotePath = ifd.IfdPath{ifd.IfdRootID, ifdexif.IfdExifID}
)

// MakernoteIfd is the Makernote IFD "IFD/MakerNote" for Makernote
var MakernoteIfd = ifd.IfdItem{MakernotePath, IfdMakernoteID, IfdMakernote}

// LoadMakernotesIfd - returns the exif.IfdItem for the Make of the Camera
func LoadMakernotesIfd(make string) ifd.IfdItem {
	if mknoteIfd, ok := mknoteRegistry[make]; ok {
		return mknoteIfd
	}
	return ifd.IfdItem{}
}

// Makernote registry
var mknoteRegistry = map[string]ifd.IfdItem{
	"Canon": CanonMakernoteIfd,
}
