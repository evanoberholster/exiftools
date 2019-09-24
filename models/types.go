package models

import (
	"fmt"
	"strconv"
)

// Unknown variables
var (
	NoFlashFired        = FlashMode{false, 0}
	UnkownExposureMode  = ExposureMode(0)
	UnknownMeteringMode = MeteringMode(0)
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

// ExposureBias - [0] Denominator [1] Numerator
type ExposureBias [2]int64

// NewExposureBias - Set ExposureBias from Numerator and Denominator
func NewExposureBias(num, denom int64) ExposureBias {
	return ExposureBias{denom, num}
}

// String - String value of Exposure Bias
func (eb ExposureBias) String() string {
	return fmt.Sprintf("%d/%d", eb[1], eb[0])
}

// ShutterSpeed - [0] Denominator [1] Numerator
type ShutterSpeed [2]int64

// NewShutterSpeed - Set ShutterSpeed from Numerator and Denominator
func NewShutterSpeed(num, denom int64) ShutterSpeed {
	return ShutterSpeed{denom, num}
}

// String - return a ShutterSpeed as a string
func (ss ShutterSpeed) String() string {
	if ss[0] == 0 {
		return strconv.Itoa(int(ss[1]))
	}
	if ss[1] == 0 {
		return "Unknown"
	}
	return fmt.Sprintf("%d/%d", ss[1], ss[0])
}

// Orientation - Orientation of Image from exif
type Orientation uint8

// NewOrientation - Create new Image Orientation from Exif
func NewOrientation(i int) Orientation {
	return Orientation(i)
}

// String - Return an Orientation as a string
func (o Orientation) String() string {
	return OrientationValues[o]
}

// OrientationValues - Image orientation values
var OrientationValues = map[Orientation]string{
	1: "Horizontal (normal)",
	2: "Mirror horizontal",
	3: "Rotate 180",
	4: "Mirror vertical",
	5: "Mirror horizontal and rotate 270 CW",
	6: "Rotate 90 CW",
	7: "Mirror horizontal and rotate 90 CW",
	8: "Rotate 270 CW",
}
