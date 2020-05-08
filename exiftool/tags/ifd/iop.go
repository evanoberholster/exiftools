package ifd

import (
	"github.com/evanoberholster/exiftools/exiftool"
	"github.com/evanoberholster/exiftools/exiftool/exif"
)

// IfdIop Name and TagID
const (
	IfdIop              = "Iop"
	IfdIopID exif.TagID = 0xA005
)

// IopPath is the IFD/Iop Ifd Path
var (
	IopPath = exif.IfdPath{IfdRootID}
)

// IopIfd is the IFD/Iop IFD for Interoperability Information
var IopIfd = exiftool.IfdItem{IopPath, IfdIopID, IfdIop}
