package ifdexif

import (
	"github.com/evanoberholster/exiftools/exiftool"
	"github.com/evanoberholster/exiftools/exiftool/exif"
	"github.com/evanoberholster/exiftools/exiftool/tags/ifd"
)

// IfdExif Name and TagID
const (
	IfdExif              = "Exif"
	IfdExifID exif.TagID = 0x8769
)

// ExifPath is the IFD/Exif Ifd Path
var (
	ExifPath = exif.IfdPath{ifd.IfdRootID}
)

// ExifIfd is the IFD/Exif IFD for ExifData
var ExifIfd = exiftool.IfdItem{ExifPath, IfdExifID, IfdExif}

// ExifIfdTags is a map of the the exif.TagID to exif.Tag for "IFD/Exif"
var ExifIfdTags = map[exif.TagID]exif.Tag{
	ExposureTime:              {"ExposureTime", exif.TypeRational},
	FNumber:                   {"FNumber", exif.TypeRational},
	ExposureProgram:           {"ExposureProgram", exif.TypeShort},
	SpectralSensitivity:       {"SpectralSensitivity", exif.TypeASCII},
	ISOSpeedRatings:           {"ISOSpeedRatings", exif.TypeShort},
	OECF:                      {"OECF", exif.TypeUndefined}, // Unknown
	SensitivityType:           {"SensitivityType", exif.TypeShort},
	StandardOutputSensitivity: {"StandardOutputSensitivity", exif.TypeLong},
	RecommendedExposureIndex:  {"RecommendedExposureIndex", exif.TypeLong},
	ISOSpeed:                  {"ISOSpeed", exif.TypeLong},
	ISOSpeedLatitudeyyy:       {"ISOSpeedLatitudeyyy", exif.TypeLong},
	ISOSpeedLatitudezzz:       {"ISOSpeedLatitudezzz", exif.TypeLong},
	ExifVersion:               {"ExifVersion", exif.TypeUndefined}, // Unknown
	DateTimeOriginal:          {"DateTimeOriginal", exif.TypeASCII},
	DateTimeDigitized:         {"DateTimeDigitized", exif.TypeASCII},
	ComponentsConfiguration:   {"ComponentsConfiguration", exif.TypeUndefined}, // Unknown
	CompressedBitsPerPixel:    {"CompressedBitsPerPixel", exif.TypeRational},
	ShutterSpeedValue:         {"ShutterSpeedValue", exif.TypeSignedRational},
	ApertureValue:             {"ApertureValue", exif.TypeRational},
	BrightnessValue:           {"BrightnessValue", exif.TypeSignedRational},
	ExposureBiasValue:         {"ExposureBiasValue", exif.TypeSignedRational},
	MaxApertureValue:          {"MaxApertureValue", exif.TypeRational},
	SubjectDistance:           {"SubjectDistance", exif.TypeRational},
	MeteringMode:              {"MeteringMode", exif.TypeShort},
	LightSource:               {"LightSource", exif.TypeShort},
	Flash:                     {"Flash", exif.TypeShort},
	FocalLength:               {"FocalLength", exif.TypeRational},
	SubjectArea:               {"SubjectArea", exif.TypeShort},
	MakerNote:                 {"MakerNote", exif.TypeUndefined},   // Unknown
	UserComment:               {"UserComment", exif.TypeUndefined}, // Unknown
	SubSecTime:                {"SubSecTime", exif.TypeASCII},
	SubSecTimeOriginal:        {"SubSecTimeOriginal", exif.TypeASCII},
	SubSecTimeDigitized:       {"SubSecTimeDigitized", exif.TypeASCII},
	FlashpixVersion:           {"FlashpixVersion", exif.TypeUndefined}, // Unknown
	ColorSpace:                {"ColorSpace", exif.TypeShort},
	PixelXDimension:           {"PixelXDimension", exif.TypeLong},
	PixelYDimension:           {"PixelYDimension", exif.TypeLong},
	RelatedSoundFile:          {"RelatedSoundFile", exif.TypeASCII},
	InteroperabilityTag:       {"InteroperabilityTag", exif.TypeLong},
	FlashEnergy:               {"FlashEnergy", exif.TypeRational},
	SpatialFrequencyResponse:  {"SpatialFrequencyResponse", exif.TypeUndefined}, // Unknown
	FocalPlaneXResolution:     {"FocalPlaneXResolution", exif.TypeRational},
	FocalPlaneYResolution:     {"FocalPlaneYResolution", exif.TypeRational},
	FocalPlaneResolutionUnit:  {"FocalPlaneResolutionUnit", exif.TypeShort},
	SubjectLocation:           {"SubjectLocation", exif.TypeShort},
	ExposureIndex:             {"ExposureIndex", exif.TypeRational},
	SensingMethod:             {"SensingMethod", exif.TypeShort},
	FileSource:                {"FileSource", exif.TypeUndefined}, // Unknown
	SceneType:                 {"SceneType", exif.TypeUndefined},  // Unknown
	CFAPattern:                {"CFAPattern", exif.TypeUndefined}, // Unknown
	CustomRendered:            {"CustomRendered", exif.TypeShort},
	ExposureMode:              {"ExposureMode", exif.TypeShort},
	WhiteBalance:              {"WhiteBalance", exif.TypeShort},
	DigitalZoomRatio:          {"DigitalZoomRatio", exif.TypeRational},
	FocalLengthIn35mmFilm:     {"FocalLengthIn35mmFilm", exif.TypeShort},
	SceneCaptureType:          {"SceneCaptureType", exif.TypeShort},
	GainControl:               {"GainControl", exif.TypeShort},
	Contrast:                  {"Contrast", exif.TypeShort},
	Saturation:                {"Saturation", exif.TypeShort},
	Sharpness:                 {"Sharpness", exif.TypeShort},
	DeviceSettingDescription:  {"DeviceSettingDescription", exif.TypeUndefined}, // Unknown
	SubjectDistanceRange:      {"SubjectDistanceRange", exif.TypeShort},
	ImageUniqueID:             {"ImageUniqueID", exif.TypeASCII},
	CameraOwnerName:           {"CameraOwnerName", exif.TypeASCII},
	BodySerialNumber:          {"BodySerialNumber", exif.TypeASCII},
	LensSpecification:         {"LensSpecification", exif.TypeRational},
	LensMake:                  {"LensMake", exif.TypeASCII},
	LensModel:                 {"LensModel", exif.TypeASCII},
	LensSerialNumber:          {"LensSerialNumber", exif.TypeASCII},
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
