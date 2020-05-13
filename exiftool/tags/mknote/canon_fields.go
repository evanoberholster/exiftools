package mknote

// CanonContinuousDrive is part of the CanonCameraSettings field
type CanonContinuousDrive int16

func (ccd CanonContinuousDrive) Int16() int16 {
	return int16(ccd)
}

func (ccd CanonContinuousDrive) String() string {
	return mapCanonContinuousDrive[ccd]
}

var mapCanonContinuousDrive = map[CanonContinuousDrive]string{
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

// CanonFocusMode is part of the CanonCameraSettings field
type CanonFocusMode int16

func (cfm CanonFocusMode) String() string {
	return mapCanonFocusMode[cfm]
}

var mapCanonFocusMode = map[CanonFocusMode]string{
	0:   "One-shot AF",
	1:   "AI Servo AF",
	2:   "AI Focus AF",
	3:   "Manual Focus",
	4:   "Single",
	5:   "Continuous",
	6:   "Manual Focus",
	16:  "Pan Focus",
	256: "AF + MF",
	512: "Movie Snap Focus",
	519: "Movie Servo AF",
}

// CanonMeteringMode is part of the CanonCameraSettings field
type CanonMeteringMode int16

func (cmm CanonMeteringMode) String() string {
	return mapCanonMeteringMode[cmm]
}

var mapCanonMeteringMode = map[CanonMeteringMode]string{
	0: "Default",
	1: "Spot",
	2: "Average",
	3: "Evaluative",
	4: "Partial",
	5: "Center-weighted average",
}

// CanonFocusRange is part of the CanonCameraSettings field
type CanonFocusRange int16

func (cfr CanonFocusRange) String() string {
	return mapCanonFocusRange[cfr]
}

var mapCanonFocusRange = map[CanonFocusRange]string{
	0: "Manual",
	1: "Auto",
	2: "Not Known",
	3: "Macro",
	4: "Very Close",
	5: "Close	   	",
	6:  "Middle Range",
	7:  "Far Range",
	8:  "Pan Focus",
	9:  "Super Macro",
	10: "Infinity",
}

// CanonExposureMode is part of the CanonCameraSettings field
type CanonExposureMode int16

func (cem CanonExposureMode) String() string {
	return mapCanonExposureMode[cem]
}

var mapCanonExposureMode = map[CanonExposureMode]string{
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

// FocusDistance -
type FocusDistance [2]int16

// NewFocusDistance creates a new FocusDistance with the upper
// and lower limits
func NewFocusDistance(upper, lower uint16) FocusDistance {
	return FocusDistance{int16(upper), int16(lower)}
}

// CanonBracketMode - Canon Makernote Backet Mode
type CanonBracketMode int16

func (cbm CanonBracketMode) String() string {
	return mapCanonBracketMode[cbm]
}

// Active - returns true if BracketMode is On
func (cbm CanonBracketMode) Active() bool {
	return cbm != 0
}

var mapCanonBracketMode = map[CanonBracketMode]string{
	0: "Off",
	1: "AEB",
	2: "FEB",
	3: "ISO",
	4: "WB",
}

// CanonAESetting - Canon Makernote AutoExposure Setting
type CanonAESetting int16

var mapCanonAESetting = map[CanonAESetting]string{
	0: "Normal AE",
	1: "Exposure Compensation",
	2: "AE Lock",
	3: "AE Lock + Exposure Compensation",
	4: "No AE",
}

type CanonAFAreaMode int16

func (caf CanonAFAreaMode) String() string {
	return mapCanonAFAreaMode[caf]
}

var mapCanonAFAreaMode = map[CanonAFAreaMode]string{
	0:  "Off (Manual Focus)",
	1:  "AF Point Expansion (surround)",
	2:  "Single-point AF",
	4:  "Auto",
	5:  "Face Detect AF",
	6:  "Face + Tracking",
	7:  "Zone AF",
	8:  "AF Point Expansion (4 point)",
	9:  "Spot AF",
	10: "AF Point Expansion (8 point)",
	11: "Flexizone Multi (49 point)",
	12: "Flexizone Multi (9 point)",
	13: "Flexizone Single",
	14: "Large Zone AF",
}

// AFPoint - AutoFocusPoint
type AFPoint [4]int16

// NewAFPoint - creates a new AFPoint from
// width, height, x-axis coord and y-axis coord
func NewAFPoint(w, h, x, y int16) AFPoint {
	return AFPoint{w, h, x, y}
}
