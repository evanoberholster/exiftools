package ifd

import (
	"github.com/evanoberholster/exiftools/exiftool"
	"github.com/evanoberholster/exiftools/exiftool/exif"
)

// IfdExif Name and TagID
const (
	IfdExif              = "Exif"
	IfdExifID exif.TagID = 0x8769
)

// ExifPath is the IFD/Exif Ifd Path
var (
	ExifPath = exif.IfdPath{IfdRootID}
)

// ExifIfd is the IFD/Exif IFD for ExifData
var ExifIfd = exiftool.IfdItem{ExifPath, IfdExifID, IfdExif}
