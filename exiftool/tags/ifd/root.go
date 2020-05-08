package ifd

import (
	"github.com/evanoberholster/exiftools/exiftool"
	"github.com/evanoberholster/exiftools/exiftool/exif"
)

// IfdRoot Name and TagID
const (
	IfdRoot              = "IFD"
	IfdRootID exif.TagID = 0x0000
)

// RootPath is the Root Ifd Path
var (
	RootPath = exif.IfdPath{}
)

// RootIfd is the Root IFD "IFD0" for ExifData
var RootIfd = exiftool.IfdItem{RootPath, IfdRootID, IfdRoot}

// RootIfdTags is a map of the the exif.TagID to exiftool.IndexedTag
var RootIfdTags = map[exif.TagID]exiftool.IndexedTag{
	ProcessingSoftware: {ProcessingSoftware, "ProcessingSoftware", exif.TypeASCII},
}

const (
	ProcessingSoftware        exif.TagID = 0x000b
	NewSubfileType            exif.TagID = 0x00fe
	SubfileType               exif.TagID = 0x00ff
	ImageWidth                exif.TagID = 0x0100
	ImageLength               exif.TagID = 0x0101
	BitsPerSample             exif.TagID = 0x0102
	Compression               exif.TagID = 0x0103
	PhotometricInterpretation exif.TagID = 0x0106
	Thresholding              exif.TagID = 0x0107
	CellWidth                 exif.TagID = 0x0108
	CellLength                exif.TagID = 0x0109
	FillOrder                 exif.TagID = 0x010a
	DocumentName              exif.TagID = 0x010d
	ImageDescription          exif.TagID = 0x010e
	Make                      exif.TagID = 0x010f
	Model                     exif.TagID = 0x0110
)
