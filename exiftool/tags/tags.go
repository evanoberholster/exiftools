package tags

import "github.com/evanoberholster/exiftools/exiftool/exif"

// IFD0 Tags
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
	ExposureTime                exif.TagID = 0x829a
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
	SubSecTime                  exif.TagID = 0x9290 // fractional seconds for ModifyDate
	SubSecTimeOriginal          exif.TagID = 0x9291 // fractional seconds for DateTimeOriginal
	SubSecTimeDigitized         exif.TagID = 0x9292 // fractional seconds for CreateDate
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

// RootIfdTagNames -
var RootIfdTagNames = map[exif.TagID]string{
	ProcessingSoftware:          "ProcessingSoftware",
	NewSubfileType:              "NewSubfileType",
	SubfileType:                 "SubfileType",
	ImageWidth:                  "ImageWidth",
	ImageLength:                 "ImageLength",
	BitsPerSample:               "BitsPerSample",
	Compression:                 "Compression",
	PhotometricInterpretation:   "PhotometricInterpretation",
	Thresholding:                "Thresholding",
	CellWidth:                   "CellWidth",
	CellLength:                  "CellLength",
	FillOrder:                   "FillOrder",
	DocumentName:                "DocumentName",
	ImageDescription:            "ImageDescription",
	Make:                        "Make",
	Model:                       "Model",
	StripOffsets:                "StripOffsets",
	Orientation:                 "Orientation",
	SamplesPerPixel:             "SamplesPerPixel",
	RowsPerStrip:                "RowsPerStrip",
	StripByteCounts:             "StripByteCounts",
	XResolution:                 "XResolution",
	YResolution:                 "YResolution",
	PlanarConfiguration:         "PlanarConfiguration",
	GrayResponseUnit:            "GrayResponseUnit",
	GrayResponseCurve:           "GrayResponseCurve",
	T4Options:                   "T4Options",
	T6Options:                   "T6Options",
	ResolutionUnit:              "ResolutionUnit",
	PageNumber:                  "PageNumber",
	TransferFunction:            "TransferFunction",
	Software:                    "Software",
	DateTime:                    "DateTime",
	Artist:                      "Artist",
	HostComputer:                "HostComputer",
	Predictor:                   "Predictor",
	WhitePoint:                  "WhitePoint",
	PrimaryChromaticities:       "PrimaryChromaticities",
	ColorMap:                    "ColorMap",
	HalftoneHints:               "HalftoneHints",
	TileWidth:                   "TileWidth",
	TileLength:                  "TileLength",
	TileOffsets:                 "TileOffsets",
	TileByteCounts:              "TileByteCounts",
	SubIFDs:                     "SubIFDs",
	InkSet:                      "InkSet",
	InkNames:                    "InkNames",
	NumberOfInks:                "NumberOfInks",
	DotRange:                    "DotRange",
	TargetPrinter:               "TargetPrinter",
	ExtraSamples:                "ExtraSamples",
	SampleFormat:                "SampleFormat",
	SMinSampleValue:             "SMinSampleValue",
	SMaxSampleValue:             "SMaxSampleValue",
	TransferRange:               "TransferRange",
	ClipPath:                    "ClipPath",
	XClipPathUnits:              "XClipPathUnits",
	YClipPathUnits:              "YClipPathUnits",
	Indexed:                     "Indexed",
	JPEGTables:                  "JPEGTables",
	OPIProxy:                    "OPIProxy",
	JPEGProc:                    "JPEGProc",
	JPEGInterchangeFormat:       "JPEGInterchangeFormat",
	JPEGInterchangeFormatLength: "JPEGInterchangeFormatLength",
	JPEGRestartInterval:         "JPEGRestartInterval",
	JPEGLosslessPredictors:      "JPEGLosslessPredictors",
	JPEGPointTransforms:         "JPEGPointTransforms",
	JPEGQTables:                 "JPEGQTables",
	JPEGDCTables:                "JPEGDCTables",
	JPEGACTables:                "JPEGACTables",
	YCbCrCoefficients:           "YCbCrCoefficients",
	YCbCrSubSampling:            "YCbCrSubSampling",
	YCbCrPositioning:            "YCbCrPositioning",
	ReferenceBlackWhite:         "ReferenceBlackWhite",
	XMLPacket:                   "XMLPacket",
	Rating:                      "Rating",
	RatingPercent:               "RatingPercent",
	ImageID:                     "ImageID",
	CFARepeatPatternDim:         "CFARepeatPatternDim",
	CFAPattern:                  "CFAPattern",
	BatteryLevel:                "BatteryLevel",
	Copyright:                   "Copyright",
	ExposureTime:                "ExposureTime",
	FNumber:                     "FNumber",
	IPTCNAA:                     "IPTCNAA",
	ImageResources:              "ImageResources",
	ExifTag:                     "ExifTag",
	InterColorProfile:           "InterColorProfile",
	ExposureProgram:             "ExposureProgram",
	SpectralSensitivity:         "SpectralSensitivity",
	GPSTag:                      "GPSTag",
	ISOSpeedRatings:             "ISOSpeedRatings",
	OECF:                        "OECF",
	Interlace:                   "Interlace",
	SensitivityType:             "SensitivityType",
	TimeZoneOffset:              "TimeZoneOffset",
	SelfTimerMode:               "SelfTimerMode",
	RecommendedExposureIndex:    "RecommendedExposureIndex",
	DateTimeOriginal:            "DateTimeOriginal",
	DateTimeDigitized:           "DateTimeDigitized",
	CompressedBitsPerPixel:      "CompressedBitsPerPixel",
	ShutterSpeedValue:           "ShutterSpeedValue",
	ApertureValue:               "ApertureValue",
	BrightnessValue:             "BrightnessValue",
	ExposureBiasValue:           "ExposureBiasValue",
	MaxApertureValue:            "MaxApertureValue",
	SubjectDistance:             "SubjectDistance",
	MeteringMode:                "MeteringMode",
	LightSource:                 "LightSource",
	Flash:                       "Flash",
	FocalLength:                 "FocalLength",
	FlashEnergy:                 "FlashEnergy",
	SpatialFrequencyResponse:    "SpatialFrequencyResponse",
	Noise:                       "Noise",
	FocalPlaneXResolution:       "FocalPlaneXResolution",
	FocalPlaneYResolution:       "FocalPlaneYResolution",
	FocalPlaneResolutionUnit:    "FocalPlaneResolutionUnit",
	ImageNumber:                 "ImageNumber",
	SecurityClassification:      "SecurityClassification",
	ImageHistory:                "ImageHistory",
	SubjectLocation:             "SubjectLocation",
	ExposureIndex:               "ExposureIndex",
	TIFFEPStandardID:            "TIFFEPStandardID",
	SensingMethod:               "SensingMethod",
	SubSecTime:                  "SubSecTime",
	SubSecTimeOriginal:          "SubSecTimeOriginal",
	SubSecTimeDigitized:         "SubSecTimeDigitized",
	XPTitle:                     "XPTitle",
	XPComment:                   "XPComment",
	XPAuthor:                    "XPAuthor",
	XPKeywords:                  "XPKeywords",
	XPSubject:                   "XPSubject",
	PrintImageMatching:          "PrintImageMatching",
	DNGVersion:                  "DNGVersion",
	DNGBackwardVersion:          "DNGBackwardVersion",
	UniqueCameraModel:           "UniqueCameraModel",
	LocalizedCameraModel:        "LocalizedCameraModel",
	CFAPlaneColor:               "CFAPlaneColor",
	CFALayout:                   "CFALayout",
	LinearizationTable:          "LinearizationTable",
	BlackLevelRepeatDim:         "BlackLevelRepeatDim",
	BlackLevel:                  "BlackLevel",
	BlackLevelDeltaH:            "BlackLevelDeltaH",
	BlackLevelDeltaV:            "BlackLevelDeltaV",
	WhiteLevel:                  "WhiteLevel",
	DefaultScale:                "DefaultScale",
	DefaultCropOrigin:           "DefaultCropOrigin",
	DefaultCropSize:             "DefaultCropSize",
	ColorMatrix1:                "ColorMatrix1",
	ColorMatrix2:                "ColorMatrix2",
	CameraCalibration1:          "CameraCalibration1",
	CameraCalibration2:          "CameraCalibration2",
	ReductionMatrix1:            "ReductionMatrix1",
	ReductionMatrix2:            "ReductionMatrix2",
	AnalogBalance:               "AnalogBalance",
	AsShotNeutral:               "AsShotNeutral",
	AsShotWhiteXY:               "AsShotWhiteXY",
	BaselineExposure:            "BaselineExposure",
	BaselineNoise:               "BaselineNoise",
	BaselineSharpness:           "BaselineSharpness",
	BayerGreenSplit:             "BayerGreenSplit",
	LinearResponseLimit:         "LinearResponseLimit",
	CameraSerialNumber:          "CameraSerialNumber",
	LensInfo:                    "LensInfo",
	ChromaBlurRadius:            "ChromaBlurRadius",
	AntiAliasStrength:           "AntiAliasStrength",
	ShadowScale:                 "ShadowScale",
	DNGPrivateData:              "DNGPrivateData",
	MakerNoteSafety:             "MakerNoteSafety",
	CalibrationIlluminant1:      "CalibrationIlluminant1",
	CalibrationIlluminant2:      "CalibrationIlluminant2",
	BestQualityScale:            "BestQualityScale",
	RawDataUniqueID:             "RawDataUniqueID",
	OriginalRawFileName:         "OriginalRawFileName",
	OriginalRawFileData:         "OriginalRawFileData",
	ActiveArea:                  "ActiveArea",
	MaskedAreas:                 "MaskedAreas",
	AsShotICCProfile:            "AsShotICCProfile",
	AsShotPreProfileMatrix:      "AsShotPreProfileMatrix",
	CurrentICCProfile:           "CurrentICCProfile",
	CurrentPreProfileMatrix:     "CurrentPreProfileMatrix",
	ColorimetricReference:       "ColorimetricReference",
	CameraCalibrationSignature:  "CameraCalibrationSignature",
	ProfileCalibrationSignature: "ProfileCalibrationSignature",
	AsShotProfileName:           "AsShotProfileName",
	NoiseReductionApplied:       "NoiseReductionApplied",
	ProfileName:                 "ProfileName",
	ProfileHueSatMapDims:        "ProfileHueSatMapDims",
	ProfileHueSatMapData1:       "ProfileHueSatMapData1",
	ProfileHueSatMapData2:       "ProfileHueSatMapData2",
	ProfileToneCurve:            "ProfileToneCurve",
	ProfileEmbedPolicy:          "ProfileEmbedPolicy",
	ProfileCopyright:            "ProfileCopyright",
	ForwardMatrix1:              "ForwardMatrix1",
	ForwardMatrix2:              "ForwardMatrix2",
	PreviewApplicationName:      "PreviewApplicationName",
	PreviewApplicationVersion:   "PreviewApplicationVersion",
	PreviewSettingsName:         "PreviewSettingsName",
	PreviewSettingsDigest:       "PreviewSettingsDigest",
	PreviewColorSpace:           "PreviewColorSpace",
	PreviewDateTime:             "PreviewDateTime",
	RawImageDigest:              "RawImageDigest",
	OriginalRawFileDigest:       "OriginalRawFileDigest",
	SubTileBlockSize:            "SubTileBlockSize",
	RowInterleaveFactor:         "RowInterleaveFactor",
	ProfileLookTableDims:        "ProfileLookTableDims",
	ProfileLookTableData:        "ProfileLookTableData",
	OpcodeList1:                 "OpcodeList1",
	OpcodeList2:                 "OpcodeList2",
	OpcodeList3:                 "OpcodeList3",
	NoiseProfile:                "NoiseProfile",
}
