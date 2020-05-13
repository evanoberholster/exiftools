package ifdexif

import (
	"github.com/evanoberholster/exiftools/exiftool/exif"
	"github.com/evanoberholster/exiftools/exiftool/tags/ifd"
)

// IfdExif Name and TagID
const (
	IfdExif              = "Exif"
	FqIfdExif            = "IFD/Exif"
	IfdExifID exif.TagID = 0x8769
)

// ExifPath is the IFD/Exif Ifd Path
var (
	ExifPath = exif.IfdPath{ifd.IfdRootID}
)

// ExifIfd is the IFD/Exif IFD for ExifData
var ExifIfd = exif.IfdItem{ExifPath, IfdExifID, IfdExif}

// ExifIfdTags is a map of the the exif.TagID to exif.Tag for "IFD/Exif"
var ExifIfdTags = map[exif.TagID]exif.Tag{
	ExposureTime:              exif.NewTag("ExposureTime", exif.TypeRational),
	FNumber:                   exif.NewTag("FNumber", exif.TypeRational),
	ExposureProgram:           exif.NewTag("ExposureProgram", exif.TypeShort),
	SpectralSensitivity:       exif.NewTag("SpectralSensitivity", exif.TypeASCII),
	ISOSpeedRatings:           exif.NewTag("ISOSpeedRatings", exif.TypeShort),
	OECF:                      exif.NewTag("OECF", exif.TypeUndefined), // Unknown
	SensitivityType:           exif.NewTag("SensitivityType", exif.TypeShort),
	StandardOutputSensitivity: exif.NewTag("StandardOutputSensitivity", exif.TypeLong),
	RecommendedExposureIndex:  exif.NewTag("RecommendedExposureIndex", exif.TypeLong),
	ISOSpeed:                  exif.NewTag("ISOSpeed", exif.TypeLong),
	ISOSpeedLatitudeyyy:       exif.NewTag("ISOSpeedLatitudeyyy", exif.TypeLong),
	ISOSpeedLatitudezzz:       exif.NewTag("ISOSpeedLatitudezzz", exif.TypeLong),
	ExifVersion:               exif.NewTag("ExifVersion", exif.TypeUndefined), // Unknown
	DateTimeOriginal:          exif.NewTag("DateTimeOriginal", exif.TypeASCII),
	DateTimeDigitized:         exif.NewTag("DateTimeDigitized", exif.TypeASCII),
	ComponentsConfiguration:   exif.NewTag("ComponentsConfiguration", exif.TypeUndefined), // Unknown
	CompressedBitsPerPixel:    exif.NewTag("CompressedBitsPerPixel", exif.TypeRational),
	ShutterSpeedValue:         exif.NewTag("ShutterSpeedValue", exif.TypeSignedRational),
	ApertureValue:             exif.NewTag("ApertureValue", exif.TypeRational),
	BrightnessValue:           exif.NewTag("BrightnessValue", exif.TypeSignedRational),
	ExposureBiasValue:         exif.NewTag("ExposureBiasValue", exif.TypeSignedRational),
	MaxApertureValue:          exif.NewTag("MaxApertureValue", exif.TypeRational),
	SubjectDistance:           exif.NewTag("SubjectDistance", exif.TypeRational),
	MeteringMode:              exif.NewTag("MeteringMode", exif.TypeShort),
	LightSource:               exif.NewTag("LightSource", exif.TypeShort),
	Flash:                     exif.NewTag("Flash", exif.TypeShort),
	FocalLength:               exif.NewTag("FocalLength", exif.TypeRational),
	SubjectArea:               exif.NewTag("SubjectArea", exif.TypeShort),
	MakerNote:                 exif.NewTag("MakerNote", exif.TypeUndefined),   // Unknown
	UserComment:               exif.NewTag("UserComment", exif.TypeUndefined), // Unknown
	SubSecTime:                exif.NewTag("SubSecTime", exif.TypeASCII),
	SubSecTimeOriginal:        exif.NewTag("SubSecTimeOriginal", exif.TypeASCII),
	SubSecTimeDigitized:       exif.NewTag("SubSecTimeDigitized", exif.TypeASCII),
	FlashpixVersion:           exif.NewTag("FlashpixVersion", exif.TypeUndefined), // Unknown
	ColorSpace:                exif.NewTag("ColorSpace", exif.TypeShort),
	PixelXDimension:           exif.NewTag("PixelXDimension", exif.TypeLong),
	PixelYDimension:           exif.NewTag("PixelYDimension", exif.TypeLong),
	RelatedSoundFile:          exif.NewTag("RelatedSoundFile", exif.TypeASCII),
	InteroperabilityTag:       exif.NewTag("InteroperabilityTag", exif.TypeLong),
	FlashEnergy:               exif.NewTag("FlashEnergy", exif.TypeRational),
	SpatialFrequencyResponse:  exif.NewTag("SpatialFrequencyResponse", exif.TypeUndefined), // Unknown
	FocalPlaneXResolution:     exif.NewTag("FocalPlaneXResolution", exif.TypeRational),
	FocalPlaneYResolution:     exif.NewTag("FocalPlaneYResolution", exif.TypeRational),
	FocalPlaneResolutionUnit:  exif.NewTag("FocalPlaneResolutionUnit", exif.TypeShort),
	SubjectLocation:           exif.NewTag("SubjectLocation", exif.TypeShort),
	ExposureIndex:             exif.NewTag("ExposureIndex", exif.TypeRational),
	SensingMethod:             exif.NewTag("SensingMethod", exif.TypeShort),
	FileSource:                exif.NewTag("FileSource", exif.TypeUndefined), // Unknown
	SceneType:                 exif.NewTag("SceneType", exif.TypeUndefined),  // Unknown
	CFAPattern:                exif.NewTag("CFAPattern", exif.TypeUndefined), // Unknown
	CustomRendered:            exif.NewTag("CustomRendered", exif.TypeShort),
	ExposureMode:              exif.NewTag("ExposureMode", exif.TypeShort),
	WhiteBalance:              exif.NewTag("WhiteBalance", exif.TypeShort),
	DigitalZoomRatio:          exif.NewTag("DigitalZoomRatio", exif.TypeRational),
	FocalLengthIn35mmFilm:     exif.NewTag("FocalLengthIn35mmFilm", exif.TypeShort),
	SceneCaptureType:          exif.NewTag("SceneCaptureType", exif.TypeShort),
	GainControl:               exif.NewTag("GainControl", exif.TypeShort),
	Contrast:                  exif.NewTag("Contrast", exif.TypeShort),
	Saturation:                exif.NewTag("Saturation", exif.TypeShort),
	Sharpness:                 exif.NewTag("Sharpness", exif.TypeShort),
	DeviceSettingDescription:  exif.NewTag("DeviceSettingDescription", exif.TypeUndefined), // Unknown
	SubjectDistanceRange:      exif.NewTag("SubjectDistanceRange", exif.TypeShort),
	ImageUniqueID:             exif.NewTag("ImageUniqueID", exif.TypeASCII),
	CameraOwnerName:           exif.NewTag("CameraOwnerName", exif.TypeASCII),
	BodySerialNumber:          exif.NewTag("BodySerialNumber", exif.TypeASCII),
	LensSpecification:         exif.NewTag("LensSpecification", exif.TypeRational),
	LensMake:                  exif.NewTag("LensMake", exif.TypeASCII),
	LensModel:                 exif.NewTag("LensModel", exif.TypeASCII),
	LensSerialNumber:          exif.NewTag("LensSerialNumber", exif.TypeASCII),
}

// ExifIFD TagIDs
const (
	ExposureTime              exif.TagID = 0x829a
	FNumber                   exif.TagID = 0x829d
	ExposureProgram           exif.TagID = 0x8822
	SpectralSensitivity       exif.TagID = 0x8824
	ISOSpeedRatings           exif.TagID = 0x8827
	OECF                      exif.TagID = 0x8828
	SensitivityType           exif.TagID = 0x8830
	StandardOutputSensitivity exif.TagID = 0x8831
	RecommendedExposureIndex  exif.TagID = 0x8832
	ISOSpeed                  exif.TagID = 0x8833
	ISOSpeedLatitudeyyy       exif.TagID = 0x8834
	ISOSpeedLatitudezzz       exif.TagID = 0x8835
	ExifVersion               exif.TagID = 0x9000
	DateTimeOriginal          exif.TagID = 0x9003
	DateTimeDigitized         exif.TagID = 0x9004
	ComponentsConfiguration   exif.TagID = 0x9101
	CompressedBitsPerPixel    exif.TagID = 0x9102
	ShutterSpeedValue         exif.TagID = 0x9201
	ApertureValue             exif.TagID = 0x9202
	BrightnessValue           exif.TagID = 0x9203
	ExposureBiasValue         exif.TagID = 0x9204
	MaxApertureValue          exif.TagID = 0x9205
	SubjectDistance           exif.TagID = 0x9206
	MeteringMode              exif.TagID = 0x9207
	LightSource               exif.TagID = 0x9208
	Flash                     exif.TagID = 0x9209
	FocalLength               exif.TagID = 0x920a
	SubjectArea               exif.TagID = 0x9214
	MakerNote                 exif.TagID = 0x927c
	UserComment               exif.TagID = 0x9286
	SubSecTime                exif.TagID = 0x9290 // fractional seconds for ModifyDate
	SubSecTimeOriginal        exif.TagID = 0x9291 // fractional seconds for DateTimeOriginal
	SubSecTimeDigitized       exif.TagID = 0x9292 // fractional seconds for CreateDate
	FlashpixVersion           exif.TagID = 0xa000
	ColorSpace                exif.TagID = 0xa001
	PixelXDimension           exif.TagID = 0xa002
	PixelYDimension           exif.TagID = 0xa003
	RelatedSoundFile          exif.TagID = 0xa004
	InteroperabilityTag       exif.TagID = 0xa005
	FlashEnergy               exif.TagID = 0xa20b
	SpatialFrequencyResponse  exif.TagID = 0xa20c
	FocalPlaneXResolution     exif.TagID = 0xa20e
	FocalPlaneYResolution     exif.TagID = 0xa20f
	FocalPlaneResolutionUnit  exif.TagID = 0xa210
	SubjectLocation           exif.TagID = 0xa214
	ExposureIndex             exif.TagID = 0xa215
	SensingMethod             exif.TagID = 0xa217
	FileSource                exif.TagID = 0xa300
	SceneType                 exif.TagID = 0xa301
	CFAPattern                exif.TagID = 0xa302
	CustomRendered            exif.TagID = 0xa401
	ExposureMode              exif.TagID = 0xa402
	WhiteBalance              exif.TagID = 0xa403
	DigitalZoomRatio          exif.TagID = 0xa404
	FocalLengthIn35mmFilm     exif.TagID = 0xa405
	SceneCaptureType          exif.TagID = 0xa406
	GainControl               exif.TagID = 0xa407
	Contrast                  exif.TagID = 0xa408
	Saturation                exif.TagID = 0xa409
	Sharpness                 exif.TagID = 0xa40a
	DeviceSettingDescription  exif.TagID = 0xa40b
	SubjectDistanceRange      exif.TagID = 0xa40c
	ImageUniqueID             exif.TagID = 0xa420
	CameraOwnerName           exif.TagID = 0xa430
	BodySerialNumber          exif.TagID = 0xa431
	LensSpecification         exif.TagID = 0xa432
	LensMake                  exif.TagID = 0xa433
	LensModel                 exif.TagID = 0xa434
	LensSerialNumber          exif.TagID = 0xa435
)
