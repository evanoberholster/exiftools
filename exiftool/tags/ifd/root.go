package ifd

import (
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
var RootIfd = exif.IfdItem{RootPath, IfdRootID, IfdRoot}

// RootIfdTags is a map of the the exif.TagID to exif.Tag for "IFD"
var RootIfdTags = map[exif.TagID]exif.Tag{
	ProcessingSoftware:          exif.NewTag("ProcessingSoftware", exif.TypeASCII),
	NewSubfileType:              exif.NewTag("NewSubfileType", exif.TypeLong),
	SubfileType:                 exif.NewTag("SubfileType", exif.TypeShort),
	ImageWidth:                  exif.NewTag("ImageWidth", exif.TypeLong),
	ImageLength:                 exif.NewTag("ImageLength", exif.TypeLong),
	BitsPerSample:               exif.NewTag("BitsPerSample", exif.TypeShort),
	Compression:                 exif.NewTag("Compression", exif.TypeShort),
	PhotometricInterpretation:   exif.NewTag("PhotometricInterpretation", exif.TypeShort),
	Thresholding:                exif.NewTag("Thresholding", exif.TypeShort),
	CellWidth:                   exif.NewTag("CellWidth", exif.TypeShort),
	CellLength:                  exif.NewTag("CellLength", exif.TypeShort),
	FillOrder:                   exif.NewTag("FillOrder", exif.TypeShort),
	DocumentName:                exif.NewTag("DocumentName", exif.TypeASCII),
	ImageDescription:            exif.NewTag("ImageDescription", exif.TypeASCII),
	Make:                        exif.NewTag("Make", exif.TypeASCII),
	Model:                       exif.NewTag("Model", exif.TypeASCII),
	StripOffsets:                exif.NewTag("StripOffsets", exif.TypeLong),
	Orientation:                 exif.NewTag("Orientation", exif.TypeShort),
	SamplesPerPixel:             exif.NewTag("SamplesPerPixel", exif.TypeShort),
	RowsPerStrip:                exif.NewTag("RowsPerStrip", exif.TypeLong),
	StripByteCounts:             exif.NewTag("StripByteCounts", exif.TypeLong),
	XResolution:                 exif.NewTag("XResolution", exif.TypeRational),
	YResolution:                 exif.NewTag("YResolution", exif.TypeRational),
	PlanarConfiguration:         exif.NewTag("PlanarConfiguration", exif.TypeShort),
	GrayResponseUnit:            exif.NewTag("GrayResponseUnit", exif.TypeShort),
	GrayResponseCurve:           exif.NewTag("GrayResponseCurve", exif.TypeShort),
	T4Options:                   exif.NewTag("T4Options", exif.TypeLong),
	T6Options:                   exif.NewTag("T6Options", exif.TypeLong),
	ResolutionUnit:              exif.NewTag("ResolutionUnit", exif.TypeShort),
	PageNumber:                  exif.NewTag("PageNumber", exif.TypeShort),
	TransferFunction:            exif.NewTag("TransferFunction", exif.TypeShort),
	Software:                    exif.NewTag("Software", exif.TypeASCII),
	DateTime:                    exif.NewTag("DateTime", exif.TypeASCII),
	Artist:                      exif.NewTag("Artist", exif.TypeASCII),
	HostComputer:                exif.NewTag("HostComputer", exif.TypeASCII),
	Predictor:                   exif.NewTag("Predictor", exif.TypeShort),
	WhitePoint:                  exif.NewTag("WhitePoint", exif.TypeRational),
	PrimaryChromaticities:       exif.NewTag("PrimaryChromaticities", exif.TypeRational),
	ColorMap:                    exif.NewTag("ColorMap", exif.TypeShort),
	HalftoneHints:               exif.NewTag("HalftoneHints", exif.TypeShort),
	TileWidth:                   exif.NewTag("TileWidth", exif.TypeShort),
	TileLength:                  exif.NewTag("TileLength", exif.TypeShort),
	TileOffsets:                 exif.NewTag("TileOffsets", exif.TypeShort),
	TileByteCounts:              exif.NewTag("TileByteCounts", exif.TypeShort),
	SubIFDs:                     exif.NewTag("SubIFDs", exif.TypeLong),
	InkSet:                      exif.NewTag("InkSet", exif.TypeShort),
	InkNames:                    exif.NewTag("InkNames", exif.TypeASCII),
	NumberOfInks:                exif.NewTag("NumberOfInks", exif.TypeShort),
	DotRange:                    exif.NewTag("DotRange", exif.TypeByte),
	TargetPrinter:               exif.NewTag("TargetPrinter", exif.TypeASCII),
	ExtraSamples:                exif.NewTag("ExtraSamples", exif.TypeShort),
	SampleFormat:                exif.NewTag("SampleFormat", exif.TypeShort),
	SMinSampleValue:             exif.NewTag("SMinSampleValue", exif.TypeShort),
	SMaxSampleValue:             exif.NewTag("SMaxSampleValue", exif.TypeShort),
	TransferRange:               exif.NewTag("TransferRange", exif.TypeShort),
	ClipPath:                    exif.NewTag("ClipPath", exif.TypeByte),
	XClipPathUnits:              exif.NewTag("XClipPathUnits", exif.TypeUndefined), // Unknown SSHORT
	YClipPathUnits:              exif.NewTag("YClipPathUnits", exif.TypeUndefined), // Unknown SSHORT
	Indexed:                     exif.NewTag("Indexed", exif.TypeShort),
	JPEGTables:                  exif.NewTag("JPEGTables", exif.TypeUndefined),
	OPIProxy:                    exif.NewTag("OPIProxy", exif.TypeShort),
	JPEGProc:                    exif.NewTag("JPEGProc", exif.TypeLong),
	JPEGInterchangeFormat:       exif.NewTag("JPEGInterchangeFormat", exif.TypeLong),
	JPEGInterchangeFormatLength: exif.NewTag("JPEGInterchangeFormatLength", exif.TypeLong),
	JPEGRestartInterval:         exif.NewTag("JPEGRestartInterval", exif.TypeShort),
	JPEGLosslessPredictors:      exif.NewTag("JPEGLosslessPredictors", exif.TypeShort),
	JPEGPointTransforms:         exif.NewTag("JPEGPointTransforms", exif.TypeShort),
	JPEGQTables:                 exif.NewTag("JPEGQTables", exif.TypeLong),
	JPEGDCTables:                exif.NewTag("JPEGDCTables", exif.TypeLong),
	JPEGACTables:                exif.NewTag("JPEGACTables", exif.TypeLong),
	YCbCrCoefficients:           exif.NewTag("YCbCrCoefficients", exif.TypeRational),
	YCbCrSubSampling:            exif.NewTag("YCbCrSubSampling", exif.TypeShort),
	YCbCrPositioning:            exif.NewTag("YCbCrPositioning", exif.TypeShort),
	ReferenceBlackWhite:         exif.NewTag("ReferenceBlackWhite", exif.TypeRational),
	XMLPacket:                   exif.NewTag("XMLPacket", exif.TypeByte),
	Rating:                      exif.NewTag("Rating", exif.TypeShort),
	RatingPercent:               exif.NewTag("RatingPercent", exif.TypeShort),
	ImageID:                     exif.NewTag("ImageID", exif.TypeASCII),
	CFARepeatPatternDim:         exif.NewTag("CFARepeatPatternDim", exif.TypeShort),
	CFAPattern:                  exif.NewTag("CFAPattern", exif.TypeByte),
	BatteryLevel:                exif.NewTag("BatteryLevel", exif.TypeRational),
	Copyright:                   exif.NewTag("Copyright", exif.TypeASCII),
	ExposureTime:                exif.NewTag("ExposureTime", exif.TypeRational),
	FNumber:                     exif.NewTag("FNumber", exif.TypeRational),
	IPTCNAA:                     exif.NewTag("IPTCNAA", exif.TypeLong),
	ImageResources:              exif.NewTag("ImageResources", exif.TypeByte),
	ExifTag:                     exif.NewTag("ExifTag", exif.TypeLong),
	InterColorProfile:           exif.NewTag("InterColorProfile", exif.TypeUndefined), // Unknown
	ExposureProgram:             exif.NewTag("ExposureProgram", exif.TypeShort),
	SpectralSensitivity:         exif.NewTag("SpectralSensitivity", exif.TypeASCII),
	GPSTag:                      exif.NewTag("GPSTag", exif.TypeLong),
	ISOSpeedRatings:             exif.NewTag("ISOSpeedRatings", exif.TypeShort),
	OECF:                        exif.NewTag("OECF", exif.TypeUndefined), // Unknown
	Interlace:                   exif.NewTag("Interlace", exif.TypeShort),
	TimeZoneOffset:              exif.NewTag("TimeZoneOffset", exif.TypeUndefined), // Unknown SSHORT
	SelfTimerMode:               exif.NewTag("SelfTimerMode", exif.TypeShort),
	DateTimeOriginal:            exif.NewTag("DateTimeOriginal", exif.TypeASCII),
	CompressedBitsPerPixel:      exif.NewTag("CompressedBitsPerPixel", exif.TypeRational),
	ShutterSpeedValue:           exif.NewTag("ShutterSpeedValue", exif.TypeSignedRational),
	ApertureValue:               exif.NewTag("ApertureValue", exif.TypeRational),
	BrightnessValue:             exif.NewTag("BrightnessValue", exif.TypeSignedRational),
	ExposureBiasValue:           exif.NewTag("ExposureBiasValue", exif.TypeSignedRational),
	MaxApertureValue:            exif.NewTag("MaxApertureValue", exif.TypeRational),
	SubjectDistance:             exif.NewTag("SubjectDistance", exif.TypeSignedRational),
	MeteringMode:                exif.NewTag("MeteringMode", exif.TypeShort),
	LightSource:                 exif.NewTag("LightSource", exif.TypeShort),
	Flash:                       exif.NewTag("Flash", exif.TypeShort),
	FocalLength:                 exif.NewTag("FocalLength", exif.TypeRational),
	FlashEnergy:                 exif.NewTag("FlashEnergy", exif.TypeRational),
	SpatialFrequencyResponse:    exif.NewTag("SpatialFrequencyResponse", exif.TypeUndefined), // Unknown
	Noise:                       exif.NewTag("Noise", exif.TypeUndefined),                    // Unknown
	FocalPlaneXResolution:       exif.NewTag("FocalPlaneXResolution", exif.TypeRational),
	FocalPlaneYResolution:       exif.NewTag("FocalPlaneYResolution", exif.TypeRational),
	FocalPlaneResolutionUnit:    exif.NewTag("FocalPlaneResolutionUnit", exif.TypeShort),
	ImageNumber:                 exif.NewTag("ImageNumber", exif.TypeLong),
	SecurityClassification:      exif.NewTag("SecurityClassification", exif.TypeASCII),
	ImageHistory:                exif.NewTag("ImageHistory", exif.TypeASCII),
	SubjectLocation:             exif.NewTag("SubjectLocation", exif.TypeShort),
	ExposureIndex:               exif.NewTag("ExposureIndex", exif.TypeRational),
	TIFFEPStandardID:            exif.NewTag("TIFFEPStandardID", exif.TypeByte),
	SensingMethod:               exif.NewTag("SensingMethod", exif.TypeShort),
	XPTitle:                     exif.NewTag("XPTitle", exif.TypeByte),
	XPComment:                   exif.NewTag("XPComment", exif.TypeByte),
	XPAuthor:                    exif.NewTag("XPAuthor", exif.TypeByte),
	XPKeywords:                  exif.NewTag("XPKeywords", exif.TypeByte),
	XPSubject:                   exif.NewTag("XPSubject", exif.TypeByte),
	PrintImageMatching:          exif.NewTag("PrintImageMatching", exif.TypeUndefined), // Unknown
	DNGVersion:                  exif.NewTag("DNGVersion", exif.TypeByte),
	DNGBackwardVersion:          exif.NewTag("DNGBackwardVersion", exif.TypeByte),
	UniqueCameraModel:           exif.NewTag("UniqueCameraModel", exif.TypeASCII),
	LocalizedCameraModel:        exif.NewTag("LocalizedCameraModel", exif.TypeByte),
	CFAPlaneColor:               exif.NewTag("CFAPlaneColor", exif.TypeByte),
	CFALayout:                   exif.NewTag("CFALayout", exif.TypeShort),
	LinearizationTable:          exif.NewTag("LinearizationTable", exif.TypeShort),
	BlackLevelRepeatDim:         exif.NewTag("BlackLevelRepeatDim", exif.TypeShort),
	BlackLevel:                  exif.NewTag("BlackLevel", exif.TypeRational),
	BlackLevelDeltaH:            exif.NewTag("BlackLevelDeltaH", exif.TypeSignedRational),
	BlackLevelDeltaV:            exif.NewTag("BlackLevelDeltaV", exif.TypeSignedRational),
	WhiteLevel:                  exif.NewTag("WhiteLevel", exif.TypeShort),
	DefaultScale:                exif.NewTag("DefaultScale", exif.TypeRational),
	DefaultCropOrigin:           exif.NewTag("DefaultCropOrigin", exif.TypeShort),
	DefaultCropSize:             exif.NewTag("DefaultCropSize", exif.TypeShort),
	ColorMatrix1:                exif.NewTag("ColorMatrix1", exif.TypeSignedRational),
	ColorMatrix2:                exif.NewTag("ColorMatrix2", exif.TypeSignedRational),
	CameraCalibration1:          exif.NewTag("CameraCalibration1", exif.TypeSignedRational),
	CameraCalibration2:          exif.NewTag("CameraCalibration2", exif.TypeSignedRational),
	ReductionMatrix1:            exif.NewTag("ReductionMatrix1", exif.TypeSignedRational),
	ReductionMatrix2:            exif.NewTag("ReductionMatrix2", exif.TypeSignedRational),
	AnalogBalance:               exif.NewTag("AnalogBalance", exif.TypeRational),
	AsShotNeutral:               exif.NewTag("AsShotNeutral", exif.TypeShort),
	AsShotWhiteXY:               exif.NewTag("AsShotWhiteXY", exif.TypeRational),
	BaselineExposure:            exif.NewTag("BaselineExposure", exif.TypeSignedRational),
	BaselineNoise:               exif.NewTag("BaselineNoise", exif.TypeRational),
	BaselineSharpness:           exif.NewTag("BaselineSharpness", exif.TypeRational),
	BayerGreenSplit:             exif.NewTag("BayerGreenSplit", exif.TypeLong),
	LinearResponseLimit:         exif.NewTag("LinearResponseLimit", exif.TypeRational),
	CameraSerialNumber:          exif.NewTag("CameraSerialNumber", exif.TypeASCII),
	LensInfo:                    exif.NewTag("LensInfo", exif.TypeRational),
	ChromaBlurRadius:            exif.NewTag("ChromaBlurRadius", exif.TypeRational),
	AntiAliasStrength:           exif.NewTag("AntiAliasStrength", exif.TypeRational),
	ShadowScale:                 exif.NewTag("ShadowScale", exif.TypeSignedRational),
	DNGPrivateData:              exif.NewTag("DNGPrivateData", exif.TypeByte),
	MakerNoteSafety:             exif.NewTag("MakerNoteSafety", exif.TypeShort),
	CalibrationIlluminant1:      exif.NewTag("CalibrationIlluminant1", exif.TypeShort),
	CalibrationIlluminant2:      exif.NewTag("CalibrationIlluminant2", exif.TypeShort),
	BestQualityScale:            exif.NewTag("BestQualityScale", exif.TypeRational),
	RawDataUniqueID:             exif.NewTag("RawDataUniqueID", exif.TypeByte),
	OriginalRawFileName:         exif.NewTag("OriginalRawFileName", exif.TypeByte),
	OriginalRawFileData:         exif.NewTag("OriginalRawFileData", exif.TypeUndefined), // Unknown
	ActiveArea:                  exif.NewTag("ActiveArea", exif.TypeShort),
	MaskedAreas:                 exif.NewTag("MaskedAreas", exif.TypeShort),
	AsShotICCProfile:            exif.NewTag("AsShotICCProfile", exif.TypeUndefined), // Unknown
	AsShotPreProfileMatrix:      exif.NewTag("AsShotPreProfileMatrix", exif.TypeSignedRational),
	CurrentICCProfile:           exif.NewTag("CurrentICCProfile", exif.TypeUndefined), // Unknown
	CurrentPreProfileMatrix:     exif.NewTag("CurrentPreProfileMatrix", exif.TypeSignedRational),
	ColorimetricReference:       exif.NewTag("ColorimetricReference", exif.TypeShort),
	CameraCalibrationSignature:  exif.NewTag("CameraCalibrationSignature", exif.TypeByte),
	ProfileCalibrationSignature: exif.NewTag("ProfileCalibrationSignature", exif.TypeByte),
	AsShotProfileName:           exif.NewTag("AsShotProfileName", exif.TypeByte),
	NoiseReductionApplied:       exif.NewTag("NoiseReductionApplied", exif.TypeRational),
	ProfileName:                 exif.NewTag("ProfileName", exif.TypeByte),
	ProfileHueSatMapDims:        exif.NewTag("ProfileHueSatMapDims", exif.TypeLong),
	ProfileHueSatMapData1:       exif.NewTag("ProfileHueSatMapData1", exif.TypeUndefined), // Unknown
	ProfileHueSatMapData2:       exif.NewTag("ProfileHueSatMapData2", exif.TypeUndefined), // Unknown
	ProfileToneCurve:            exif.NewTag("ProfileToneCurve", exif.TypeUndefined),      // Unknown
	ProfileEmbedPolicy:          exif.NewTag("ProfileEmbedPolicy", exif.TypeLong),
	ProfileCopyright:            exif.NewTag("ProfileCopyright", exif.TypeByte),
	ForwardMatrix1:              exif.NewTag("ForwardMatrix1", exif.TypeSignedRational),
	ForwardMatrix2:              exif.NewTag("ForwardMatrix2", exif.TypeSignedRational),
	PreviewApplicationName:      exif.NewTag("PreviewApplicationName", exif.TypeByte),
	PreviewApplicationVersion:   exif.NewTag("PreviewApplicationVersion", exif.TypeByte),
	PreviewSettingsName:         exif.NewTag("PreviewSettingsName", exif.TypeByte),
	PreviewSettingsDigest:       exif.NewTag("PreviewSettingsDigest", exif.TypeByte),
	PreviewColorSpace:           exif.NewTag("PreviewColorSpace", exif.TypeLong),
	PreviewDateTime:             exif.NewTag("PreviewDateTime", exif.TypeASCII),
	RawImageDigest:              exif.NewTag("RawImageDigest", exif.TypeUndefined),        // Unknown
	OriginalRawFileDigest:       exif.NewTag("OriginalRawFileDigest", exif.TypeUndefined), // Unknown
	SubTileBlockSize:            exif.NewTag("SubTileBlockSize", exif.TypeLong),
	RowInterleaveFactor:         exif.NewTag("RowInterleaveFactor", exif.TypeLong),
	ProfileLookTableDims:        exif.NewTag("ProfileLookTableDims", exif.TypeLong),
	ProfileLookTableData:        exif.NewTag("ProfileLookTableData", exif.TypeUndefined), // Unknown FLOAT
	OpcodeList1:                 exif.NewTag("OpcodeList1", exif.TypeUndefined),          // Unknown
	OpcodeList2:                 exif.NewTag("OpcodeList2", exif.TypeUndefined),          // Unknown
	OpcodeList3:                 exif.NewTag("OpcodeList3", exif.TypeUndefined),          // Unknown
	NoiseProfile:                exif.NewTag("NoiseProfile", exif.TypeUndefined),         // Unknown DOUBLE
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
