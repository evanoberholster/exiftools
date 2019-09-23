package main

import (
	"fmt"
	"strconv"
	"time"

	_ "trimmer.io/go-xmp/models"
	"trimmer.io/go-xmp/models/dc"
	xmpbase "trimmer.io/go-xmp/models/xmp_base"
	"trimmer.io/go-xmp/xmp"

	"github.com/evanoberholster/exif/exif"
	"github.com/evanoberholster/exif/tiff"
)

// ExposureMode - Mode in which the Exposure was taken.
type ExposureMode uint8

// NewExposureMode -
func NewExposureMode(m int) ExposureMode {
	return ExposureMode(m)
}

// String - Return Exposure Mode as a string
func (em ExposureMode) String() string {
	return exposureModeValues[em]
}

// UnkownExposureMode - Unknown Exposure Mode
var UnkownExposureMode = ExposureMode(0)

// ExposureModeValues -
var exposureModeValues = map[ExposureMode]string{
	0: "Not Defined",
	1: "Manual",
	2: "Program AE",
	3: "Aperture-priority AE",
	4: "Shutter speed priority AE",
	5: "Creative (Slow speed)",
	6: "Action (High speed)",
	7: "Portrait",
	8: "Landscape",
	9: "Bulb",
}

// MeteringMode - Mode in which the Photo was metered.
type MeteringMode uint8

// NewMeteringMode - Create new Metering Mode
func NewMeteringMode(m int) MeteringMode {
	return MeteringMode(m)
}

// String - Return Metering Mode as a string
func (mm MeteringMode) String() string {
	return meteringModeValues[mm]
}

// UnknownMeteringMode - Unknown Metering Mode
var UnknownMeteringMode = MeteringMode(0)

// MeteringModeValues -
// Derived from https://sno.phy.queensu.ca/~phil/exiftool/TagNames/EXIF.html (23/09/2019)
var meteringModeValues = map[MeteringMode]string{
	0:   "Unknown",
	1:   "Average",
	2:   "Center-weighted average",
	3:   "Spot",
	4:   "Multi-spot",
	5:   "Multi-segment",
	6:   "Partial",
	255: "Other",
}

// OrientationValues -
var OrientationValues = map[int]string{
	1: "Horizontal (normal)",
	2: "Mirror horizontal",
	3: "Rotate 180",
	4: "Mirror vertical",
	5: "Mirror horizontal and rotate 270 CW",
	6: "Rotate 90 CW",
	7: "Mirror horizontal and rotate 90 CW",
	8: "Rotate 270 CW",
}

// CameraSettingsField -
type CameraSettingsField map[int]string

// CameraSettings -
type CameraSettings struct {
	RecordMode      string `json:"RecordMode"`
	FocusMode       string `json:"FocusMode"`
	ExposureMode    string `json:"ExposureMode"`
	ContinuousDrive string `json:"ContinuousDrive"`
	MeteringMode    string `json:"MeteringMode"`
}

// ProcessCameraSettingsFields -
func ProcessCameraSettingsFields(tag *tiff.Tag, i int) string {
	a, err := tag.Int(i)
	if err != nil {
		return ""
	}
	return CanonCameraSettingsFields[i][a]
}

// ISOSpeed -
func ISOSpeed(x *exif.Exif) (int, error) {
	a, err := x.Get(exif.ISOSpeedRatings)
	if err != nil {
		return 0, err
	}
	return a.Int(0)
}

// Aperture -
func Aperture(x *exif.Exif) (float32, error) {
	a, err := x.Get(exif.FNumber)
	if err != nil {
		return 0.0, err
	}
	num, denom, err := a.Rat2(0)
	return float32(num) / float32(denom), err
}

// GetShutterSpeed -
func GetShutterSpeed(x *exif.Exif) (ShutterSpeed, error) {
	a, err := x.Get(exif.ExposureTime)
	if err != nil {
		return NewShutterSpeed(0, 0), err
	}
	nom, denom, err := a.Rat2(0)
	return NewShutterSpeed(nom, denom), err
}

// ExposureBias -
func ExposureBias(x *exif.Exif) (Rational, error) {
	r := Rational{}
	e, err := x.Get(exif.ExposureBiasValue)
	if err != nil {
		return r.Set(0, 0), err
	}
	num, denom, err := e.Rat2(0)
	return r.Set(num, denom), err
}

// GetExposureMode -
func GetExposureMode(x *exif.Exif) (ExposureMode, error) {
	tag, err := x.Get(exif.ExposureProgram)
	if err != nil {
		return UnkownExposureMode, err
	}
	v, err := tag.Int(0)
	if err != nil {
		return UnkownExposureMode, err
	}
	return NewExposureMode(v), nil
}

// GetMeteringMode -
func GetMeteringMode(x *exif.Exif) (MeteringMode, error) {
	tag, err := x.Get(exif.MeteringMode)
	if err != nil {
		return UnknownMeteringMode, err
	}
	v, err := tag.Int(0)
	if err != nil {
		return UnknownMeteringMode, err
	}
	return NewMeteringMode(v), nil
}

// Orientation -
func Orientation(x *exif.Exif) (int, error) {
	i, err := x.Get(exif.Orientation)
	if err != nil {
		return 1, err
	}
	return i.Int(0)
}

// Flash -
func Flash(x *exif.Exif) (FlashMode, error) {
	tag, err := x.Get(exif.Flash)
	if err != nil {
		return NoFlashFired, err
	}
	v, err := tag.Int(0)
	if err != nil {
		return NoFlashFired, err
	}
	return NewFlashMode(v), nil
}

// Rational -
type Rational struct {
	Num, Denom int64 // Numerator and Denominator
	Value      string
}

// ShutterSpeed -
type ShutterSpeed struct {
	N int64
	D int64
}

// NewShutterSpeed - Set ShutterSpeed from Numerator and Demoninator
func NewShutterSpeed(nom, denom int64) ShutterSpeed {
	return ShutterSpeed{nom, denom}
}

// String -
func (ss ShutterSpeed) String() string {
	if ss.D == 0 {
		return strconv.Itoa(int(ss.D))
	}
	if ss.N == 0 {
		return "Unknown"
	}
	return fmt.Sprintf("%d/%d", ss.N, ss.D)
}

// Set -
func (r Rational) Set(n, d int64) Rational {
	r.Num = n
	r.Denom = d
	r.Value = fmt.Sprintf("%d/%d", n, d)
	return r
}

// FetchString -
func FetchString(x *exif.Exif, tagLabel exif.FieldName) (string, error) {
	a, err := x.Get(tagLabel)
	if err != nil {
		return "", err
	}
	return a.StringVal()
}

// Metadata -
type Metadata struct {
	FileSize   int64      `json:"FileSize"`
	MIMEType   string     `json:"MIMEType"`
	DublinCore DublinCore `json:"DublinCore"`
	XmpBase    XmpBase    `json:"XmpBase"`
	Exif       Exif       `json:"Exif"`
}

// Exif Metadata
type Exif struct {
	ImageWidth       int            `json:"ImageWidth"`
	ImageHeight      int            `json:"ImageHeight"`
	CameraMake       string         `json:"CameraMake"`   // OK
	CameraModel      string         `json:"CameraModel"`  // OK
	CameraSerial     string         `json:"CameraSerial"` // OK
	LensModel        string         `json:"LensModel"`    // OK
	LensSerial       string         `json:"LensSerial"`   // OK
	Artist           string         `json:"Artist"`       // OK
	Copyright        string         `json:"Copyright"`    // OK
	Aperture         float32        `json:"Aperture"`     // OK
	ShutterSpeed     ShutterSpeed   `json:"ShutterSpeed"` // OK
	ISOSpeed         int            `json:"ISO"`          // OK
	ExposureBias     Rational       `json:"ExposureBias"` // OK
	ExposureMode     ExposureMode   `json:"ExposureMode"` // OK
	MeteringMode     MeteringMode   `json:"MeteringMode"` // OK
	Orientation      int            `json:"Orientation"`
	Flash            FlashMode      `json:"Flash"`            // OK
	FocalLength      float32        `json:"FocalLength"`      // OK
	FocalLengthEqv   float32        `json:"FocalLengthEqv"`   // mm
	GPSLatitude      float64        `json:"GPSLatitude"`      // OK
	GPSLongitude     float64        `json:"GPSLongitude"`     // OK
	GPSAltitude      float32        `json:"GPSAltitude"`      // OK
	GPSTimeStamp     time.Time      `json:"GPSTimeStamp"`     // OK
	DateTimeOriginal time.Time      `json:"DateTimeOriginal"` // OK
	ModifyTimeStamp  time.Time      `json:"ModifyTimeStamp"`
	Software         string         `json:"Software"` // OK
	ImageDescription string         `json:"ImageDescription"`
	CameraSettings   CameraSettings `json:"CameraSettings"`
}

// XMP

// DublinCore -
///// Modified DublinCore 17/04/2019 https://godoc.org/trimmer.io/go-xmp/models/dc#DublinCore
type DublinCore struct {
	Creator     string   `xmp:"dc:creator"`
	Description string   `xmp:"dc:description"`
	Format      string   `xmp:"dc:format"`
	Rights      string   `xmp:"dc:rights"`
	Source      string   `xmp:"dc:source"`
	Subject     []string `xmp:"dc:subject"`
	Title       string   `xmp:"dc:title"`
}

func xmpDublinCore(m *xmp.Document) DublinCore {
	var d DublinCore
	c := dc.FindModel(m)
	if c == nil {
		return d
	}

	creator := []string(c.Creator)
	if len(creator) > 0 {
		d.Creator = creator[0]
	}
	s := []string(c.Subject)
	if s == nil {
		s = []string{}
	}
	d.Subject = s
	d.Description = c.Description.Default()
	d.Format = string(c.Format)
	d.Rights = c.Rights.Default()
	d.Title = c.Title.Default()

	return d
}

// XmpBase -
///// Modified XmpBase 17/04/2019 https://godoc.org/trimmer.io/go-xmp/models/xmp_base#XmpBase
type XmpBase struct {
	CreateDate   time.Time `xmp:"xmp:CreateDate"`
	CreatorTool  string    `xmp:"xmp:CreatorTool"`
	Identifier   string    `xmp:"xmp:Identifier"`
	Label        string    `xmp:"xmp:Label"`
	MetadataDate time.Time `xmp:"xmp:MetadataDate"`
	ModifyDate   time.Time `xmp:"xmp:ModifyDate"`
	Rating       int       `xmp:"xmp:Rating"`
}

func xmpBase(m *xmp.Document) XmpBase {
	var b XmpBase
	c := xmpbase.FindModel(m)
	if c == nil {
		return XmpBase{}
	}

	b.CreateDate = c.CreateDate.Value()
	b.MetadataDate = c.MetadataDate.Value()
	b.ModifyDate = c.ModifyDate.Value()
	b.Label = string(c.Label)
	b.Rating = int(c.Rating)
	b.CreatorTool = c.CreatorTool.String()

	return b
}
