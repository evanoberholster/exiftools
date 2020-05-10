package mknote

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
	MakernotePath = exif.IfdPath{ifd.IfdRootID, ifdexif.IfdExifID}
)

// MakernoteIfd is the Makernote IFD "IFD/MakerNote" for Makernote
var MakernoteIfd = exif.IfdItem{MakernotePath, IfdMakernoteID, IfdMakernote}

func LoadMakerNotesIfd(make string) exif.IfdItem {
	if mknoteIfd, ok := mknoteIfdMap[make]; ok {
		return mknoteIfd
	}
	return exif.IfdItem{}
}

// Makernote registry
var mknoteIfdMap = map[string]exif.IfdItem{
	"Canon": CanonMakernoteIfd,
}
