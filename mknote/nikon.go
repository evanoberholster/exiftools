package mknote

import "github.com/evanoberholster/exif/exif"

// Nikon-specific Maker Note fields
var (
	NikonVersion        exif.FieldName = "Nikon.Version"
	NikonWhiteBalance   exif.FieldName = "Nikon.WhiteBalance"
	NikonColorSpace     exif.FieldName = "Nikon.ColorSpace"
	NikonLightSource    exif.FieldName = "Nikon.LightSource"
	NikonSaturation     exif.FieldName = "Nikon_Saturation"
	NikonShotInfo       exif.FieldName = "Nikon.ShotInfo"       // A sub-IFD
	NikonVRInfo         exif.FieldName = "Nikon.VRInfo"         // A sub-IFD
	NikonPictureControl exif.FieldName = "Nikon.PictureControl" // A sub-IFD
	NikonWorldTime      exif.FieldName = "Nikon.WorldTime"      // A sub-IFD
	NikonISOInfo        exif.FieldName = "Nikon.ISOInfo"        // A sub-IFD
	NikonAFInfo         exif.FieldName = "Nikon.AFInfo"         // A sub-IFD
	NikonColorBalance   exif.FieldName = "Nikon.ColorBalance"   // A sub-IFD
	NikonLensData       exif.FieldName = "Nikon.LensData"       // A sub-IFD
	NikonSerialNO       exif.FieldName = "Nikon.SerialNO"       // usually starts with "NO="
	NikonFlashInfo      exif.FieldName = "Nikon.FlashInfo"      // A sub-IFD
	NikonMultiExposure  exif.FieldName = "Nikon.MultiExposure"  // A sub-IFD
	NikonAFInfo2        exif.FieldName = "Nikon.AFInfo2"        // A sub-IFD
	NikonFileInfo       exif.FieldName = "Nikon.FileInfo"       // A sub-IFD
	NikonAFTune         exif.FieldName = "Nikon.AFTune"         // A sub-IFD
	Nikon3_0x000a       exif.FieldName = "Nikon3.0x000a"
	Nikon3_0x009b       exif.FieldName = "Nikon3.0x009b"
	Nikon3_0x009f       exif.FieldName = "Nikon3.0x009f"
	Nikon3_0x00a3       exif.FieldName = "Nikon3.0x00a3"
)

// Nikon version 3 Maker Notes fields (used by E5400, SQ, D2H, D70, and newer)
var makerNoteNikon3Fields = map[uint16]exif.FieldName{
	0x0001: NikonVersion,
	0x0002: ISOSpeed,
	0x0003: ColorMode,
	0x0004: Quality,
	0x0005: NikonWhiteBalance,
	0x0006: Sharpening,
	0x0007: Focus,
	0x0008: FlashSetting,
	0x0009: FlashDevice,
	0x000a: Nikon3_0x000a,
	0x000b: WhiteBalanceBias,
	0x000c: WBRBLevels,
	0x000d: ProgramShift,
	0x000e: ExposureDiff,
	0x000f: ISOSelection,
	0x0010: DataDump,
	0x0011: Preview,
	0x0012: FlashComp,
	0x0013: ISOSettings,
	0x0016: ImageBoundary,
	0x0017: FlashExposureComp,
	0x0018: FlashBracketComp,
	0x0019: ExposureBracketComp,
	0x001a: ImageProcessing,
	0x001b: CropHiSpeed,
	0x001c: ExposureTuning,
	0x001d: SerialNumber,
	0x001e: NikonColorSpace,
	0x001f: NikonVRInfo,
	0x0020: ImageAuthentication,
	0x0022: ActiveDLighting,
	0x0023: NikonPictureControl,
	0x0024: NikonWorldTime,
	0x0025: NikonISOInfo,
	0x002a: VignetteControl,
	0x0080: ImageAdjustment,
	0x0081: ToneComp,
	0x0082: AuxiliaryLens,
	0x0083: LensType,
	0x0084: Lens,
	0x0085: FocusDistance,
	0x0086: DigitalZoom,
	0x0087: FlashMode,
	0x0088: NikonAFInfo,
	0x0089: ShootingMode,
	0x008a: AutoBracketRelease,
	0x008b: LensFStops,
	0x008c: ContrastCurve,
	0x008d: ColorHue,
	0x008f: SceneMode,
	0x0090: NikonLightSource,
	0x0091: NikonShotInfo,
	0x0092: HueAdjustment,
	0x0093: NEFCompression,
	0x0094: NikonSaturation,
	0x0095: NoiseReduction,
	0x0096: LinearizationTable,
	0x0097: NikonColorBalance,
	0x0098: NikonLensData,
	0x0099: RawImageCenter,
	0x009a: SensorPixelSize,
	0x009b: Nikon3_0x009b,
	0x009c: SceneAssist,
	0x009e: RetouchHistory,
	0x009f: Nikon3_0x009f,
	0x00a0: NikonSerialNO,
	0x00a2: ImageDataSize,
	0x00a3: Nikon3_0x00a3,
	0x00a5: ImageCount,
	0x00a6: DeletedImageCount,
	0x00a7: ShutterCount,
	0x00a8: NikonFlashInfo,
	0x00a9: ImageOptimization,
	0x00aa: SaturationText,
	0x00ab: VariProgram,
	0x00ac: ImageStabilization,
	0x00ad: AFResponse,
	0x00b0: NikonMultiExposure,
	0x00b1: HighISONoiseReduction,
	0x00b3: ToningEffect,
	0x00b7: NikonAFInfo2,
	0x00b8: NikonFileInfo,
	0x00b9: NikonAFTune,
	0x0e00: PrintIM,
	0x0e01: CaptureData,
	0x0e09: CaptureVersion,
	0x0e0e: CaptureOffsets,
	0x0e10: ScanIFD,
	0x0e1d: ICCProfile,
	0x0e1e: CaptureOutput,
}

// NikonRaw - Raw Image from a Canon Camera
type NikonRaw struct{}

// RawCameraSettings -
func (nr *NikonRaw) RawCameraSettings(x *exif.Exif) (CameraSettings, error) {
	return CameraSettings{}, nil
}
