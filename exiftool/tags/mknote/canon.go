package mknote

import (
	"github.com/evanoberholster/exiftools/exiftool"
	"github.com/evanoberholster/exiftools/exiftool/exif"
)

// Makernote Name
const (
	IfdCanonMakernote = "CanonMKnote"
)

// CanonMakernoteIfd is the Makernote IFD "IFD/CanonMKnote" for Canon Cameras
var CanonMakernoteIfd = exiftool.IfdItem{MakernotePath, IfdMakernoteID, IfdCanonMakernote}

// CanonIfdTags is a map of the the exif.TagID to exif.Tag for "IFD/Exif/CanonMKnote"
// Source: https://exiftool.org/TagNames/Canon.html on 8/05/2020
var CanonIfdTags = map[exif.TagID]exif.Tag{
	CanonCameraSettings:     {"CanonCameraSettings", exif.TypeShort},
	CanonFocalLength:        {"CanonFocalLength", exif.TypeShort},
	CanonFlashInfo:          {"CanonFlashInfo", exif.TypeShort},
	CanonShotInfo:           {"CanonShotInfo", exif.TypeShort},
	CanonPanorama:           {"CanonPanorama", exif.TypeShort},
	CanonImageType:          {"CanonImageType", exif.TypeASCII},
	CanonFirmwareVersion:    {"CanonFirmwareVersion", exif.TypeASCII},
	FileNumber:              {"FileNumber", exif.TypeLong},
	OwnerName:               {"OwnerName", exif.TypeASCII},
	SerialNumber:            {"SerialNumber", exif.TypeLong},
	CanonModelID:            {"CanonModelID", exif.TypeLong},
	CanonAFInfo:             {"CanonAFInfo", exif.TypeShort},
	ThumbnailImageValidArea: {"ThumbnailImageValidArea", exif.TypeShort},
	SuperMacro:              {"SuperMacro", exif.TypeShort},
	DateStampMode:           {"DateStampMode", exif.TypeShort},
	CanonAFInfo2:            {"CanonAFInfo2", exif.TypeShort},
	ImageUniqueID:           {"ImageUniqueID", exif.TypeByte},
	TimeInfo:                {"TimeInfo", exif.TypeLong},
	CanonFileInfo:           {"CanonFileInfo", exif.TypeShort},
	LensModel:               {"LensModel", exif.TypeASCII},
}

// CanonMKnoteIFD TagIDs
// Source: https://exiftool.org/TagNames/Canon.html on 8/05/2020
const (
	CanonCameraSettings        exif.TagID = 0x0001
	CanonFocalLength           exif.TagID = 0x0002
	CanonFlashInfo             exif.TagID = 0x0003
	CanonShotInfo              exif.TagID = 0x0004
	CanonPanorama              exif.TagID = 0x0005
	CanonImageType             exif.TagID = 0x0006
	CanonFirmwareVersion       exif.TagID = 0x0007
	FileNumber                 exif.TagID = 0x0008
	OwnerName                  exif.TagID = 0x0009
	UnknownD30                 exif.TagID = 0x000a
	SerialNumber               exif.TagID = 0x000c
	CanonCameraInfo            exif.TagID = 0x000d // WIP
	CanonFileLength            exif.TagID = 0x000e // WIP
	CustomFunctions            exif.TagID = 0x000f // WIP
	CanonModelID               exif.TagID = 0x0010
	MovieInfo                  exif.TagID = 0x0011 // WIP
	CanonAFInfo                exif.TagID = 0x0012
	ThumbnailImageValidArea    exif.TagID = 0x0013 // WIP
	SerialNumberFormat         exif.TagID = 0x0015 // WIP
	SuperMacro                 exif.TagID = 0x001a // WIP
	DateStampMode              exif.TagID = 0x001c // WIP
	MyColors                   exif.TagID = 0x001d // WIP
	FirmwareRevision           exif.TagID = 0x001e // WIP
	Categories                 exif.TagID = 0x0023 // WIP
	FaceDetect1                exif.TagID = 0x0024 // WIP
	FaceDetect2                exif.TagID = 0x0025 // WIP
	CanonAFInfo2               exif.TagID = 0x0026
	ContrastInfo               exif.TagID = 0x0027 // WIP
	ImageUniqueID              exif.TagID = 0x0028 // WIP
	WBInfo                     exif.TagID = 0x0029 // WIP
	FaceDetect3                exif.TagID = 0x002f // WIP
	TimeInfo                   exif.TagID = 0x0035
	BatteryType                exif.TagID = 0x0038 // WIP
	AFInfo3                    exif.TagID = 0x003c // WIP
	RawDataOffset              exif.TagID = 0x0081 // WIP
	OriginalDecisionDataOffset exif.TagID = 0x0083 // WIP
	CustomFunctions1D          exif.TagID = 0x0090 // WIP
	PersonalFunctions          exif.TagID = 0x0091 // WIP
	PersonalFunctionValues     exif.TagID = 0x0092 // WIP
	CanonFileInfo              exif.TagID = 0x0093
	AFPointsInFocus1D          exif.TagID = 0x0094 // WIP
	LensModel                  exif.TagID = 0x0095
)