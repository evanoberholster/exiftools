package mknote

import (
	"github.com/evanoberholster/exif/exif"
)

var (
	// Canon-specific fiends
	Canon_CameraSettings exif.FieldName = "Canon.CameraSettings" // A sub-IFD
	Canon_ShotInfo       exif.FieldName = "Canon.ShotInfo"       // A sub-IFD
	Canon_AFInfo         exif.FieldName = "Canon.AFInfo"
	Canon_TimeInfo       exif.FieldName = "Canon.TimeInfo"
	Canon_0x0000         exif.FieldName = "Canon.0x0000"
	Canon_0x0003         exif.FieldName = "Canon.0x0003"
	Canon_0x00b5         exif.FieldName = "Canon.0x00b5"
	Canon_0x00c0         exif.FieldName = "Canon.0x00c0"
	Canon_0x00c1         exif.FieldName = "Canon.0x00c1"
)

var makerNoteCanonFields = map[uint16]exif.FieldName{
	0x0000: Canon_0x0000,
	0x0001: Canon_CameraSettings,
	0x0002: exif.FocalLength,
	0x0003: Canon_0x0003,
	0x0004: Canon_ShotInfo,
	0x0005: Panorama,
	0x0006: ImageType,
	0x0007: FirmwareVersion,
	0x0008: FileNumber,
	0x0009: OwnerName,
	0x000c: SerialNumber,
	0x000d: CameraInfo,
	0x000f: CustomFunctions,
	0x0010: ModelID,
	0x0012: PictureInfo,
	0x0013: ThumbnailImageValidArea,
	0x0015: SerialNumberFormat,
	0x001a: SuperMacro,
	0x0026: Canon_AFInfo,
	0x0035: Canon_TimeInfo,
	0x0083: OriginalDecisionDataOffset,
	0x00a4: WhiteBalanceTable,
	0x0095: LensModel,
	0x0096: InternalSerialNumber,
	0x0097: DustRemovalData,
	0x0099: CustomFunctions,
	0x00a0: ProcessingInfo,
	0x00aa: MeasuredColor,
	0x00b4: exif.ColorSpace,
	0x00b5: Canon_0x00b5,
	0x00c0: Canon_0x00c0,
	0x00c1: Canon_0x00c1,
	0x00d0: VRDOffset,
	0x00e0: SensorInfo,
	0x4001: ColorData,
}

// CanonRaw - Raw Image from a Canon Camera
type CanonRaw struct{}

// RawCameraSettings - Get Canon camera Settings
func (cr *CanonRaw) RawCameraSettings(x *exif.Exif) (CameraSettings, error) {
	c := CameraSettings{}
	tag, err := x.Get(Canon_CameraSettings)
	if err != nil {
		return c, err
	}

	c.ContinuousDrive = ProcessCameraSettingsFields(tag, CanonContinuousDrive)
	c.RecordMode = ProcessCameraSettingsFields(tag, CanonRecordMode)
	c.FocusMode = ProcessCameraSettingsFields(tag, CanonFocusMode)
	c.ExposureMode = ProcessCameraSettingsFields(tag, CanonExposureMode)
	c.MeteringMode = ProcessCameraSettingsFields(tag, CanonMeteringMode)

	return c, nil
}

// CanonContinuousDriveValues -
var CanonContinuousDriveValues = map[int]string{
	0:  "Single",
	1:  "Continuous",
	2:  "Movie",
	3:  "Continuous, Speed Priority",
	4:  "Continuous, Low",
	5:  "Continuous, High",
	6:  "Silent Single",
	9:  "Single, Silent",
	10: "Continuous, Silent",
}

// CanonFocusModeValues -
var CanonFocusModeValues = map[int]string{
	0:   "One-shot AF",
	1:   "AI Servo AF",
	2:   "AI Focus AF",
	3:   "Manual Focus (3)",
	4:   "Single",
	5:   "Continuous",
	6:   "Manual Focus (6)",
	16:  "Pan Focus",
	256: "AF + MF",
	512: "Movie Snap Focus",
	519: "Movie Servo AF",
}

// CanonExposureModeValues -
var CanonExposureModeValues = map[int]string{
	0: "Easy",
	1: "Program AE",
	2: "Shutter speed priority AE",
	3: "Aperture-priority AE",
	4: "Manual",
	5: "Depth-of-field AE",
	6: "M-Dep",
	7: "Bulb",
	8: "Flexible-priority AE",
}

// CanonRecordModeValues -
var CanonRecordModeValues = map[int]string{
	1:  "JPEG",
	2:  "CRW+THM",
	3:  "AVI+THM",
	4:  "TIF",
	5:  "TIF+JPEG",
	6:  "CR2",
	7:  "CR2+JPEG",
	9:  "MOV",
	10: "MP4",
	11: "CRM",
	12: "CR3",
	13: "CR3+JPEG",
}

// CanonMeteringModeValues -
var CanonMeteringModeValues = map[int]string{
	0: "Default",
	1: "Spot",
	2: "Average",
	3: "Evaluative",
	4: "Partial",
	5: "Center-weighted average",
}

// CanonAESettingValues -
var CanonAESettingValues = map[int]string{
	0: "Normal AE",
	1: "Exposure Compensation",
	2: "AE Lock",
	3: "AE Lock + Exposure Compensation",
	4: "No AE",
}

// Canon Exif information constants
const (
	CanonContinuousDrive int = 5
	CanonFocusMode       int = 7
	CanonRecordMode      int = 9
	CanonMeteringMode    int = 17
	CanonExposureMode    int = 20
	CanonAESetting       int = 33
)

// CanonCameraSettingsFields -
var CanonCameraSettingsFields = map[int]CameraSettingsField{
	CanonContinuousDrive: CanonContinuousDriveValues,
	CanonFocusMode:       CanonFocusModeValues,
	CanonRecordMode:      CanonRecordModeValues,
	CanonMeteringMode:    CanonMeteringModeValues,
	CanonExposureMode:    CanonExposureModeValues,
	CanonAESetting:       CanonAESettingValues,
}
