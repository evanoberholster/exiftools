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

// RootIfdTags is a map of the the exif.TagID to exif.Tag for "IFD"
var RootIfdTags = map[exif.TagID]exif.Tag{
	ProcessingSoftware:          {"ProcessingSoftware", exif.TypeASCII},
	NewSubfileType:              {"NewSubfileType", exif.TypeLong},
	SubfileType:                 {"SubfileType", exif.TypeShort},
	ImageWidth:                  {"ImageWidth", exif.TypeLong},
	ImageLength:                 {"ImageLength", exif.TypeLong},
	BitsPerSample:               {"BitsPerSample", exif.TypeShort},
	Compression:                 {"Compression", exif.TypeShort},
	PhotometricInterpretation:   {"PhotometricInterpretation", exif.TypeShort},
	Thresholding:                {"Thresholding", exif.TypeShort},
	CellWidth:                   {"CellWidth", exif.TypeShort},
	CellLength:                  {"CellLength", exif.TypeShort},
	FillOrder:                   {"FillOrder", exif.TypeShort},
	DocumentName:                {"DocumentName", exif.TypeASCII},
	ImageDescription:            {"ImageDescription", exif.TypeASCII},
	Make:                        {"Make", exif.TypeASCII},
	Model:                       {"Model", exif.TypeASCII},
	StripOffsets:                {"StripOffsets", exif.TypeLong},
	Orientation:                 {"Orientation", exif.TypeShort},
	SamplesPerPixel:             {"SamplesPerPixel", exif.TypeShort},
	RowsPerStrip:                {"RowsPerStrip", exif.TypeLong},
	StripByteCounts:             {"StripByteCounts", exif.TypeLong},
	XResolution:                 {"XResolution", exif.TypeRational},
	YResolution:                 {"YResolution", exif.TypeRational},
	PlanarConfiguration:         {"PlanarConfiguration", exif.TypeShort},
	GrayResponseUnit:            {"GrayResponseUnit", exif.TypeShort},
	GrayResponseCurve:           {"GrayResponseCurve", exif.TypeShort},
	T4Options:                   {"T4Options", exif.TypeLong},
	T6Options:                   {"T6Options", exif.TypeLong},
	ResolutionUnit:              {"ResolutionUnit", exif.TypeShort},
	PageNumber:                  {"PageNumber", exif.TypeShort},
	TransferFunction:            {"TransferFunction", exif.TypeShort},
	Software:                    {"Software", exif.TypeASCII},
	DateTime:                    {"DateTime", exif.TypeASCII},
	Artist:                      {"Artist", exif.TypeASCII},
	HostComputer:                {"HostComputer", exif.TypeASCII},
	Predictor:                   {"Predictor", exif.TypeShort},
	WhitePoint:                  {"WhitePoint", exif.TypeRational},
	PrimaryChromaticities:       {"PrimaryChromaticities", exif.TypeRational},
	ColorMap:                    {"ColorMap", exif.TypeShort},
	HalftoneHints:               {"HalftoneHints", exif.TypeShort},
	TileWidth:                   {"TileWidth", exif.TypeShort},
	TileLength:                  {"TileLength", exif.TypeShort},
	TileOffsets:                 {"TileOffsets", exif.TypeShort},
	TileByteCounts:              {"TileByteCounts", exif.TypeShort},
	SubIFDs:                     {"SubIFDs", exif.TypeLong},
	InkSet:                      {"InkSet", exif.TypeShort},
	InkNames:                    {"InkNames", exif.TypeASCII},
	NumberOfInks:                {"NumberOfInks", exif.TypeShort},
	DotRange:                    {"DotRange", exif.TypeByte},
	TargetPrinter:               {"TargetPrinter", exif.TypeASCII},
	ExtraSamples:                {"ExtraSamples", exif.TypeShort},
	SampleFormat:                {"SampleFormat", exif.TypeShort},
	SMinSampleValue:             {"SMinSampleValue", exif.TypeShort},
	SMaxSampleValue:             {"SMaxSampleValue", exif.TypeShort},
	TransferRange:               {"TransferRange", exif.TypeShort},
	ClipPath:                    {"ClipPath", exif.TypeByte},
	XClipPathUnits:              {"XClipPathUnits", exif.TypeUndefined}, // Unknown SSHORT
	YClipPathUnits:              {"YClipPathUnits", exif.TypeUndefined}, // Unknown SSHORT
	Indexed:                     {"Indexed", exif.TypeShort},
	JPEGTables:                  {"JPEGTables", exif.TypeUndefined},
	OPIProxy:                    {"OPIProxy", exif.TypeShort},
	JPEGProc:                    {"JPEGProc", exif.TypeLong},
	JPEGInterchangeFormat:       {"JPEGInterchangeFormat", exif.TypeLong},
	JPEGInterchangeFormatLength: {"JPEGInterchangeFormatLength", exif.TypeLong},
	JPEGRestartInterval:         {"JPEGRestartInterval", exif.TypeShort},
	JPEGLosslessPredictors:      {"JPEGLosslessPredictors", exif.TypeShort},
	JPEGPointTransforms:         {"JPEGPointTransforms", exif.TypeShort},
	JPEGQTables:                 {"JPEGQTables", exif.TypeLong},
	JPEGDCTables:                {"JPEGDCTables", exif.TypeLong},
	JPEGACTables:                {"JPEGACTables", exif.TypeLong},
	YCbCrCoefficients:           {"YCbCrCoefficients", exif.TypeRational},
	YCbCrSubSampling:            {"YCbCrSubSampling", exif.TypeShort},
	YCbCrPositioning:            {"YCbCrPositioning", exif.TypeShort},
	ReferenceBlackWhite:         {"ReferenceBlackWhite", exif.TypeRational},
	XMLPacket:                   {"XMLPacket", exif.TypeByte},
	Rating:                      {"Rating", exif.TypeShort},
	RatingPercent:               {"RatingPercent", exif.TypeShort},
	ImageID:                     {"ImageID", exif.TypeASCII},
	CFARepeatPatternDim:         {"CFARepeatPatternDim", exif.TypeShort},
	CFAPattern:                  {"CFAPattern", exif.TypeByte},
	BatteryLevel:                {"BatteryLevel", exif.TypeRational},
	Copyright:                   {"Copyright", exif.TypeASCII},
	ExposureTime:                {"ExposureTime", exif.TypeRational},
	FNumber:                     {"FNumber", exif.TypeRational},
	IPTCNAA:                     {"IPTCNAA", exif.TypeLong},
	ImageResources:              {"ImageResources", exif.TypeByte},
	ExifTag:                     {"ExifTag", exif.TypeLong},
	InterColorProfile:           {"InterColorProfile", exif.TypeUndefined}, // Unknown
	ExposureProgram:             {"ExposureProgram", exif.TypeShort},
	SpectralSensitivity:         {"SpectralSensitivity", exif.TypeASCII},
	GPSTag:                      {"GPSTag", exif.TypeLong},
	ISOSpeedRatings:             {"ISOSpeedRatings", exif.TypeShort},
	OECF:                        {"OECF", exif.TypeUndefined}, // Unknown
	Interlace:                   {"Interlace", exif.TypeShort},
	TimeZoneOffset:              {"TimeZoneOffset", exif.TypeUndefined}, // Unknown SSHORT
	SelfTimerMode:               {"SelfTimerMode", exif.TypeShort},
	DateTimeOriginal:            {"DateTimeOriginal", exif.TypeASCII},
	CompressedBitsPerPixel:      {"CompressedBitsPerPixel", exif.TypeRational},
	ShutterSpeedValue:           {"ShutterSpeedValue", exif.TypeSignedRational},
	ApertureValue:               {"ApertureValue", exif.TypeRational},
	BrightnessValue:             {"BrightnessValue", exif.TypeSignedRational},
	ExposureBiasValue:           {"ExposureBiasValue", exif.TypeSignedRational},
	MaxApertureValue:            {"MaxApertureValue", exif.TypeRational},
	SubjectDistance:             {"SubjectDistance", exif.TypeSignedRational},
	MeteringMode:                {"MeteringMode", exif.TypeShort},
	LightSource:                 {"LightSource", exif.TypeShort},
	Flash:                       {"Flash", exif.TypeShort},
	FocalLength:                 {"FocalLength", exif.TypeRational},
	FlashEnergy:                 {"FlashEnergy", exif.TypeRational},
	SpatialFrequencyResponse:    {"SpatialFrequencyResponse", exif.TypeUndefined}, // Unknown
	Noise:                       {"Noise", exif.TypeUndefined},                    // Unknown
	FocalPlaneXResolution:       {"FocalPlaneXResolution", exif.TypeRational},
	FocalPlaneYResolution:       {"FocalPlaneYResolution", exif.TypeRational},
	FocalPlaneResolutionUnit:    {"FocalPlaneResolutionUnit", exif.TypeShort},
	ImageNumber:                 {"ImageNumber", exif.TypeLong},
	SecurityClassification:      {"SecurityClassification", exif.TypeASCII},
	ImageHistory:                {"ImageHistory", exif.TypeASCII},
	SubjectLocation:             {"SubjectLocation", exif.TypeShort},
	ExposureIndex:               {"ExposureIndex", exif.TypeRational},
	TIFFEPStandardID:            {"TIFFEPStandardID", exif.TypeByte},
	SensingMethod:               {"SensingMethod", exif.TypeShort},
	XPTitle:                     {"XPTitle", exif.TypeByte},
	XPComment:                   {"XPComment", exif.TypeByte},
	XPAuthor:                    {"XPAuthor", exif.TypeByte},
	XPKeywords:                  {"XPKeywords", exif.TypeByte},
	XPSubject:                   {"XPSubject", exif.TypeByte},
	PrintImageMatching:          {"PrintImageMatching", exif.TypeUndefined}, // Unknown
	DNGVersion:                  {"DNGVersion", exif.TypeByte},
	DNGBackwardVersion:          {"DNGBackwardVersion", exif.TypeByte},
	UniqueCameraModel:           {"UniqueCameraModel", exif.TypeASCII},
	LocalizedCameraModel:        {"LocalizedCameraModel", exif.TypeByte},
	CFAPlaneColor:               {"CFAPlaneColor", exif.TypeByte},
	CFALayout:                   {"CFALayout", exif.TypeShort},
	LinearizationTable:          {"LinearizationTable", exif.TypeShort},
	BlackLevelRepeatDim:         {"BlackLevelRepeatDim", exif.TypeShort},
	BlackLevel:                  {"BlackLevel", exif.TypeRational},
	BlackLevelDeltaH:            {"BlackLevelDeltaH", exif.TypeSignedRational},
	BlackLevelDeltaV:            {"BlackLevelDeltaV", exif.TypeSignedRational},
	WhiteLevel:                  {"WhiteLevel", exif.TypeShort},
	DefaultScale:                {"DefaultScale", exif.TypeRational},
	DefaultCropOrigin:           {"DefaultCropOrigin", exif.TypeShort},
	DefaultCropSize:             {"DefaultCropSize", exif.TypeShort},
	ColorMatrix1:                {"ColorMatrix1", exif.TypeSignedRational},
	ColorMatrix2:                {"ColorMatrix2", exif.TypeSignedRational},
	CameraCalibration1:          {"CameraCalibration1", exif.TypeSignedRational},
	CameraCalibration2:          {"CameraCalibration2", exif.TypeSignedRational},
	ReductionMatrix1:            {"ReductionMatrix1", exif.TypeSignedRational},
	ReductionMatrix2:            {"ReductionMatrix2", exif.TypeSignedRational},
	AnalogBalance:               {"AnalogBalance", exif.TypeRational},
	AsShotNeutral:               {"AsShotNeutral", exif.TypeShort},
	AsShotWhiteXY:               {"AsShotWhiteXY", exif.TypeRational},
	BaselineExposure:            {"BaselineExposure", exif.TypeSignedRational},
	BaselineNoise:               {"BaselineNoise", exif.TypeRational},
	BaselineSharpness:           {"BaselineSharpness", exif.TypeRational},
	BayerGreenSplit:             {"BayerGreenSplit", exif.TypeLong},
	LinearResponseLimit:         {"LinearResponseLimit", exif.TypeRational},
	CameraSerialNumber:          {"CameraSerialNumber", exif.TypeASCII},
	LensInfo:                    {"LensInfo", exif.TypeRational},
	ChromaBlurRadius:            {"ChromaBlurRadius", exif.TypeRational},
	AntiAliasStrength:           {"AntiAliasStrength", exif.TypeRational},
	ShadowScale:                 {"ShadowScale", exif.TypeSignedRational},
	DNGPrivateData:              {"DNGPrivateData", exif.TypeByte},
	MakerNoteSafety:             {"MakerNoteSafety", exif.TypeShort},
	CalibrationIlluminant1:      {"CalibrationIlluminant1", exif.TypeShort},
	CalibrationIlluminant2:      {"CalibrationIlluminant2", exif.TypeShort},
	BestQualityScale:            {"BestQualityScale", exif.TypeRational},
	RawDataUniqueID:             {"RawDataUniqueID", exif.TypeByte},
	OriginalRawFileName:         {"OriginalRawFileName", exif.TypeByte},
	OriginalRawFileData:         {"OriginalRawFileData", exif.TypeUndefined}, // Unknown
	ActiveArea:                  {"ActiveArea", exif.TypeShort},
	MaskedAreas:                 {"MaskedAreas", exif.TypeShort},
	AsShotICCProfile:            {"AsShotICCProfile", exif.TypeUndefined}, // Unknown
	AsShotPreProfileMatrix:      {"AsShotPreProfileMatrix", exif.TypeSignedRational},
	CurrentICCProfile:           {"CurrentICCProfile", exif.TypeUndefined}, // Unknown
	CurrentPreProfileMatrix:     {"CurrentPreProfileMatrix", exif.TypeSignedRational},
	ColorimetricReference:       {"ColorimetricReference", exif.TypeShort},
	CameraCalibrationSignature:  {"CameraCalibrationSignature", exif.TypeByte},
	ProfileCalibrationSignature: {"ProfileCalibrationSignature", exif.TypeByte},
	AsShotProfileName:           {"AsShotProfileName", exif.TypeByte},
	NoiseReductionApplied:       {"NoiseReductionApplied", exif.TypeRational},
	ProfileName:                 {"ProfileName", exif.TypeByte},
	ProfileHueSatMapDims:        {"ProfileHueSatMapDims", exif.TypeLong},
	ProfileHueSatMapData1:       {"ProfileHueSatMapData1", exif.TypeUndefined}, // Unknown
	ProfileHueSatMapData2:       {"ProfileHueSatMapData2", exif.TypeUndefined}, // Unknown
	ProfileToneCurve:            {"ProfileToneCurve", exif.TypeUndefined},      // Unknown
	ProfileEmbedPolicy:          {"ProfileEmbedPolicy", exif.TypeLong},
	ProfileCopyright:            {"ProfileCopyright", exif.TypeByte},
	ForwardMatrix1:              {"ForwardMatrix1", exif.TypeSignedRational},
	ForwardMatrix2:              {"ForwardMatrix2", exif.TypeSignedRational},
	PreviewApplicationName:      {"PreviewApplicationName", exif.TypeByte},
	PreviewApplicationVersion:   {"PreviewApplicationVersion", exif.TypeByte},
	PreviewSettingsName:         {"PreviewSettingsName", exif.TypeByte},
	PreviewSettingsDigest:       {"PreviewSettingsDigest", exif.TypeByte},
	PreviewColorSpace:           {"PreviewColorSpace", exif.TypeLong},
	PreviewDateTime:             {"PreviewDateTime", exif.TypeASCII},
	RawImageDigest:              {"RawImageDigest", exif.TypeUndefined},        // Unknown
	OriginalRawFileDigest:       {"OriginalRawFileDigest", exif.TypeUndefined}, // Unknown
	SubTileBlockSize:            {"SubTileBlockSize", exif.TypeLong},
	RowInterleaveFactor:         {"RowInterleaveFactor", exif.TypeLong},
	ProfileLookTableDims:        {"ProfileLookTableDims", exif.TypeLong},
	ProfileLookTableData:        {"ProfileLookTableData", exif.TypeUndefined}, // Unknown FLOAT
	OpcodeList1:                 {"OpcodeList1", exif.TypeUndefined},          // Unknown
	OpcodeList2:                 {"OpcodeList2", exif.TypeUndefined},          // Unknown
	OpcodeList3:                 {"OpcodeList3", exif.TypeUndefined},          // Unknown
	NoiseProfile:                {"NoiseProfile", exif.TypeUndefined},         // Unknown DOUBLE
}

// RootIFD TagIDs
const (
	ProcessingSoftware          exif.TagID = 0x000b
	NewSubfileType              exif.TagID = 0x00fe
	SubfileType                 exif.TagID = 0x00ff
	ImageWidth                  exif.TagID = 0x0100
	ImageLength                 exif.TagID = 0x0101
	BitsPerSample               exif.TagID = 0x0102
	Compression                 exif.TagID = 0x0103
	PhotometricInterpretation   exif.TagID = 0x0106
	Thresholding                exif.TagID = 0x0107
	CellWidth                   exif.TagID = 0x0108
	CellLength                  exif.TagID = 0x0109
	FillOrder                   exif.TagID = 0x010a
	DocumentName                exif.TagID = 0x010d
	ImageDescription            exif.TagID = 0x010e
	Make                        exif.TagID = 0x010f
	Model                       exif.TagID = 0x0110
	StripOffsets                exif.TagID = 0x0111
	Orientation                 exif.TagID = 0x0112
	SamplesPerPixel             exif.TagID = 0x0115
	RowsPerStrip                exif.TagID = 0x0116
	StripByteCounts             exif.TagID = 0x0117
	XResolution                 exif.TagID = 0x011a
	YResolution                 exif.TagID = 0x011b
	PlanarConfiguration         exif.TagID = 0x011c
	GrayResponseUnit            exif.TagID = 0x0122
	GrayResponseCurve           exif.TagID = 0x0123
	T4Options                   exif.TagID = 0x0124
	T6Options                   exif.TagID = 0x0125
	ResolutionUnit              exif.TagID = 0x0128
	PageNumber                  exif.TagID = 0x0129
	TransferFunction            exif.TagID = 0x012d
	Software                    exif.TagID = 0x0131
	DateTime                    exif.TagID = 0x0132
	Artist                      exif.TagID = 0x013b
	HostComputer                exif.TagID = 0x013c
	Predictor                   exif.TagID = 0x013d
	WhitePoint                  exif.TagID = 0x013e
	PrimaryChromaticities       exif.TagID = 0x013f
	ColorMap                    exif.TagID = 0x0140
	HalftoneHints               exif.TagID = 0x0141
	TileWidth                   exif.TagID = 0x0142
	TileLength                  exif.TagID = 0x0143
	TileOffsets                 exif.TagID = 0x0144
	TileByteCounts              exif.TagID = 0x0145
	SubIFDs                     exif.TagID = 0x014a
	InkSet                      exif.TagID = 0x014c
	InkNames                    exif.TagID = 0x014d
	NumberOfInks                exif.TagID = 0x014e
	DotRange                    exif.TagID = 0x0150
	TargetPrinter               exif.TagID = 0x0151
	ExtraSamples                exif.TagID = 0x0152
	SampleFormat                exif.TagID = 0x0153
	SMinSampleValue             exif.TagID = 0x0154
	SMaxSampleValue             exif.TagID = 0x0155
	TransferRange               exif.TagID = 0x0156
	ClipPath                    exif.TagID = 0x0157
	XClipPathUnits              exif.TagID = 0x0158
	YClipPathUnits              exif.TagID = 0x0159
	Indexed                     exif.TagID = 0x015a
	JPEGTables                  exif.TagID = 0x015b
	OPIProxy                    exif.TagID = 0x015f
	JPEGProc                    exif.TagID = 0x0200
	JPEGInterchangeFormat       exif.TagID = 0x0201
	JPEGInterchangeFormatLength exif.TagID = 0x0202
	JPEGRestartInterval         exif.TagID = 0x0203
	JPEGLosslessPredictors      exif.TagID = 0x0205
	JPEGPointTransforms         exif.TagID = 0x0206
	JPEGQTables                 exif.TagID = 0x0207
	JPEGDCTables                exif.TagID = 0x0208
	JPEGACTables                exif.TagID = 0x0209
	YCbCrCoefficients           exif.TagID = 0x0211
	YCbCrSubSampling            exif.TagID = 0x0212
	YCbCrPositioning            exif.TagID = 0x0213
	ReferenceBlackWhite         exif.TagID = 0x0214
	XMLPacket                   exif.TagID = 0x02bc
	Rating                      exif.TagID = 0x4746
	RatingPercent               exif.TagID = 0x4749
	ImageID                     exif.TagID = 0x800d
	CFARepeatPatternDim         exif.TagID = 0x828d
	CFAPattern                  exif.TagID = 0x828e
	BatteryLevel                exif.TagID = 0x828f
	Copyright                   exif.TagID = 0x8298
	ExposureTime                exif.TagID = 0x829a // IFD/EXIF and IFD
	FNumber                     exif.TagID = 0x829d
	IPTCNAA                     exif.TagID = 0x83bb
	ImageResources              exif.TagID = 0x8649
	ExifTag                     exif.TagID = 0x8769
	InterColorProfile           exif.TagID = 0x8773
	ExposureProgram             exif.TagID = 0x8822
	SpectralSensitivity         exif.TagID = 0x8824
	GPSTag                      exif.TagID = 0x8825
	ISOSpeedRatings             exif.TagID = 0x8827
	OECF                        exif.TagID = 0x8828
	Interlace                   exif.TagID = 0x8829
	SensitivityType             exif.TagID = 0x8830
	TimeZoneOffset              exif.TagID = 0x882a
	SelfTimerMode               exif.TagID = 0x882b
	RecommendedExposureIndex    exif.TagID = 0x8832
	DateTimeOriginal            exif.TagID = 0x9003
	DateTimeDigitized           exif.TagID = 0x9004
	CompressedBitsPerPixel      exif.TagID = 0x9102
	ShutterSpeedValue           exif.TagID = 0x9201
	ApertureValue               exif.TagID = 0x9202
	BrightnessValue             exif.TagID = 0x9203
	ExposureBiasValue           exif.TagID = 0x9204
	MaxApertureValue            exif.TagID = 0x9205
	SubjectDistance             exif.TagID = 0x9206
	MeteringMode                exif.TagID = 0x9207
	LightSource                 exif.TagID = 0x9208
	Flash                       exif.TagID = 0x9209
	FocalLength                 exif.TagID = 0x920a
	FlashEnergy                 exif.TagID = 0x920b
	SpatialFrequencyResponse    exif.TagID = 0x920c
	Noise                       exif.TagID = 0x920d
	FocalPlaneXResolution       exif.TagID = 0x920e
	FocalPlaneYResolution       exif.TagID = 0x920f
	FocalPlaneResolutionUnit    exif.TagID = 0x9210
	ImageNumber                 exif.TagID = 0x9211
	SecurityClassification      exif.TagID = 0x9212
	ImageHistory                exif.TagID = 0x9213
	SubjectLocation             exif.TagID = 0x9214
	ExposureIndex               exif.TagID = 0x9215
	TIFFEPStandardID            exif.TagID = 0x9216
	SensingMethod               exif.TagID = 0x9217
	XPTitle                     exif.TagID = 0x9c9b
	XPComment                   exif.TagID = 0x9c9c
	XPAuthor                    exif.TagID = 0x9c9d
	XPKeywords                  exif.TagID = 0x9c9e
	XPSubject                   exif.TagID = 0x9c9f
	PrintImageMatching          exif.TagID = 0xc4a5
	DNGVersion                  exif.TagID = 0xc612
	DNGBackwardVersion          exif.TagID = 0xc613
	UniqueCameraModel           exif.TagID = 0xc614
	LocalizedCameraModel        exif.TagID = 0xc615
	CFAPlaneColor               exif.TagID = 0xc616
	CFALayout                   exif.TagID = 0xc617
	LinearizationTable          exif.TagID = 0xc618
	BlackLevelRepeatDim         exif.TagID = 0xc619
	BlackLevel                  exif.TagID = 0xc61a
	BlackLevelDeltaH            exif.TagID = 0xc61b
	BlackLevelDeltaV            exif.TagID = 0xc61c
	WhiteLevel                  exif.TagID = 0xc61d
	DefaultScale                exif.TagID = 0xc61e
	DefaultCropOrigin           exif.TagID = 0xc61f
	DefaultCropSize             exif.TagID = 0xc620
	ColorMatrix1                exif.TagID = 0xc621
	ColorMatrix2                exif.TagID = 0xc622
	CameraCalibration1          exif.TagID = 0xc623
	CameraCalibration2          exif.TagID = 0xc624
	ReductionMatrix1            exif.TagID = 0xc625
	ReductionMatrix2            exif.TagID = 0xc626
	AnalogBalance               exif.TagID = 0xc627
	AsShotNeutral               exif.TagID = 0xc628
	AsShotWhiteXY               exif.TagID = 0xc629
	BaselineExposure            exif.TagID = 0xc62a
	BaselineNoise               exif.TagID = 0xc62b
	BaselineSharpness           exif.TagID = 0xc62c
	BayerGreenSplit             exif.TagID = 0xc62d
	LinearResponseLimit         exif.TagID = 0xc62e
	CameraSerialNumber          exif.TagID = 0xc62f
	LensInfo                    exif.TagID = 0xc630
	ChromaBlurRadius            exif.TagID = 0xc631
	AntiAliasStrength           exif.TagID = 0xc632
	ShadowScale                 exif.TagID = 0xc633
	DNGPrivateData              exif.TagID = 0xc634
	MakerNoteSafety             exif.TagID = 0xc635
	CalibrationIlluminant1      exif.TagID = 0xc65a
	CalibrationIlluminant2      exif.TagID = 0xc65b
	BestQualityScale            exif.TagID = 0xc65c
	RawDataUniqueID             exif.TagID = 0xc65d
	OriginalRawFileName         exif.TagID = 0xc68b
	OriginalRawFileData         exif.TagID = 0xc68c
	ActiveArea                  exif.TagID = 0xc68d
	MaskedAreas                 exif.TagID = 0xc68e
	AsShotICCProfile            exif.TagID = 0xc68f
	AsShotPreProfileMatrix      exif.TagID = 0xc690
	CurrentICCProfile           exif.TagID = 0xc691
	CurrentPreProfileMatrix     exif.TagID = 0xc692
	ColorimetricReference       exif.TagID = 0xc6bf
	CameraCalibrationSignature  exif.TagID = 0xc6f3
	ProfileCalibrationSignature exif.TagID = 0xc6f4
	AsShotProfileName           exif.TagID = 0xc6f6
	NoiseReductionApplied       exif.TagID = 0xc6f7
	ProfileName                 exif.TagID = 0xc6f8
	ProfileHueSatMapDims        exif.TagID = 0xc6f9
	ProfileHueSatMapData1       exif.TagID = 0xc6fa
	ProfileHueSatMapData2       exif.TagID = 0xc6fb
	ProfileToneCurve            exif.TagID = 0xc6fc
	ProfileEmbedPolicy          exif.TagID = 0xc6fd
	ProfileCopyright            exif.TagID = 0xc6fe
	ForwardMatrix1              exif.TagID = 0xc714
	ForwardMatrix2              exif.TagID = 0xc715
	PreviewApplicationName      exif.TagID = 0xc716
	PreviewApplicationVersion   exif.TagID = 0xc717
	PreviewSettingsName         exif.TagID = 0xc718
	PreviewSettingsDigest       exif.TagID = 0xc719
	PreviewColorSpace           exif.TagID = 0xc71a
	PreviewDateTime             exif.TagID = 0xc71b
	RawImageDigest              exif.TagID = 0xc71c
	OriginalRawFileDigest       exif.TagID = 0xc71d
	SubTileBlockSize            exif.TagID = 0xc71e
	RowInterleaveFactor         exif.TagID = 0xc71f
	ProfileLookTableDims        exif.TagID = 0xc725
	ProfileLookTableData        exif.TagID = 0xc726
	OpcodeList1                 exif.TagID = 0xc740
	OpcodeList2                 exif.TagID = 0xc741
	OpcodeList3                 exif.TagID = 0xc74e
	NoiseProfile                exif.TagID = 0xc761
)
