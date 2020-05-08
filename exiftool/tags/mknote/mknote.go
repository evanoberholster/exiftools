package mknote

import (
	"github.com/evanoberholster/exiftools/exiftool"
	"github.com/evanoberholster/exiftools/exiftool/exif"
	"github.com/evanoberholster/exiftools/exiftool/tags/ifd"
	"github.com/evanoberholster/exiftools/exiftool/tags/ifdexif"
)

// Makernote Name and TagID
const (
	IfdMakernote              = "Makernote"
	IfdMakernoteID exif.TagID = 0x927c
)

// MakernotePath is the MakernotePath Ifd Path
var (
	MakernotePath = exif.IfdPath{ifd.IfdRootID, ifdexif.IfdExifID}
)

// MakernoteIfd is the Makernote IFD "IFD/MakerNote" for Makernote
var MakernoteIfd = exiftool.IfdItem{MakernotePath, IfdMakernoteID, IfdMakernote}

func LoadMakerNotes(make string) *exiftool.IfdItem {
	return &CanonMakernoteIfd
}
