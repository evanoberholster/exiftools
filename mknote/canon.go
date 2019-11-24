package mknote

import (
	"github.com/evanoberholster/exiftools/exif"
	"github.com/evanoberholster/exiftools/models"
)

// Canon-specific fields
var (
	CanonCameraSettings exif.FieldName = "Canon.CameraSettings" // A sub-IFD
	CanonShotInfo       exif.FieldName = "Canon.ShotInfo"       // A sub-IFD
	CanonAFInfo         exif.FieldName = "Canon.AFInfo"
	CanonTimeInfo       exif.FieldName = "Canon.TimeInfo"
	Canon0x0000         exif.FieldName = "Canon.0x0000"
	Canon0x0003         exif.FieldName = "Canon.0x0003"
	Canon0x00b5         exif.FieldName = "Canon.0x00b5"
	Canon0x00c0         exif.FieldName = "Canon.0x00c0"
	Canon0x00c1         exif.FieldName = "Canon.0x00c1"
)

var makerNoteCanonFields = map[uint16]exif.FieldName{
	0x0000: Canon0x0000,
	0x0001: CanonCameraSettings,
	0x0002: exif.FocalLength,
	0x0003: Canon0x0003,
	0x0004: CanonShotInfo,
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
	0x0026: CanonAFInfo,
	0x0035: CanonTimeInfo,
	0x0083: OriginalDecisionDataOffset,
	0x00a4: WhiteBalanceTable,
	0x0095: LensModel,
	0x0096: InternalSerialNumber,
	0x0097: DustRemovalData,
	0x0099: CustomFunctions,
	0x00a0: ProcessingInfo,
	0x00aa: MeasuredColor,
	0x00b4: exif.ColorSpace,
	0x00b5: Canon0x00b5,
	0x00c0: Canon0x00c0,
	0x00c1: Canon0x00c1,
	0x00d0: VRDOffset,
	0x00e0: SensorInfo,
	0x4001: ColorData,
}

// CanonRaw - Raw Image from a Canon Camera
type CanonRaw struct{}

// RawCameraSettings - Get Canon camera Settings
func (cr *CanonRaw) RawCameraSettings(x *exif.Exif) (CameraSettings, error) {
	c := CameraSettings{}
	tag, err := x.Get(CanonCameraSettings)
	if err != nil {
		return c, err
	}

	c.ContinuousDrive = processCameraSettingsFields(tag, CanonContinuousDrive)
	c.RecordMode = processCameraSettingsFields(tag, CanonRecordMode)
	c.FocusMode = processCameraSettingsFields(tag, CanonFocusMode)
	c.ExposureMode = processCameraSettingsFields(tag, CanonExposureMode)
	c.MeteringMode = processCameraSettingsFields(tag, CanonMeteringMode)

	return c, nil
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
	CanonContinuousDrive: models.CanonContinuousDriveValues,
	CanonFocusMode:       models.CanonFocusModeValues,
	CanonRecordMode:      models.CanonRecordModeValues,
	CanonMeteringMode:    models.CanonMeteringModeValues,
	CanonExposureMode:    models.CanonExposureModeValues,
	CanonAESetting:       models.CanonAESettingValues,
}

var CanonMakerNoteTimeZones = map[int]string{
	0:     "+00:00",
	1:     "+12:45",
	2:     "+12:00",
	3:     "+11:00",
	4:     "+10:00",
	5:     "+9:30",
	6:     "+9:00",
	7:     "+8:00",
	8:     "+7:00",
	9:     "+6:30",
	10:    "+6:00",
	11:    "+5:45",
	12:    "+5:30",
	13:    "+5:00",
	14:    "+4:30",
	15:    "+4:00",
	16:    "+3:30",
	17:    "+3:00",
	18:    "+2:00",
	19:    "+1:00",
	20:    "+0:00",
	21:    "-1:00",
	22:    "-2:00",
	23:    "-3:00",
	24:    "-3:30",
	25:    "-4:00",
	26:    "-4:00",
	27:    "-5:00",
	28:    "-6:00",
	29:    "-7:00",
	30:    "-8:00",
	31:    "-9:00",
	32:    "-10:00",
	33:    "-11:00",
	32766: "+00:00",
}

// CanonMakerNoteTimeZoneValues
var CanonMakerNoteTimeZoneValues = map[int]string{
	0:     "n/a",
	1:     "Chatham Islands",
	2:     "Wellington",
	3:     "Solomon Islands",
	4:     "Sydney",
	5:     "Adelaide",
	6:     "Tokyo",
	7:     "Hong Kong",
	8:     "Bangkok",
	9:     "Yangon",
	10:    "Dhaka",
	11:    "Kathmandu",
	12:    "Delhi",
	13:    "Karachi",
	14:    "Kabul",
	15:    "Dubai",
	16:    "Tehran",
	17:    "Moscow",
	18:    "Cairo",
	19:    "Paris",
	20:    "London",
	21:    "Azores",
	22:    "Fernando de Noronha",
	23:    "Sao Paulo",
	24:    "Newfoundland",
	25:    "Santiago",
	26:    "Caracas",
	27:    "New York",
	28:    "Chicago",
	29:    "Denver",
	30:    "Los Angeles",
	31:    "Anchorage",
	32:    "Honolulu",
	33:    "Samoa",
	32766: "(not set)",
}

// CanonTimeZone -
func CanonTimeZone(x *exif.Exif) string {
	ti, err := x.Get(CanonTimeInfo)
	if err != nil {
		return CanonMakerNoteTimeZones[0]
	}

	timeZone, err := ti.Int(2)
	if err != nil {
		return CanonMakerNoteTimeZones[0]
	}

	return CanonMakerNoteTimeZones[timeZone]
}
