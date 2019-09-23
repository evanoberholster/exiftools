package main

// FlashMode - Mode in which the camera Flash was used.
// (bool) - true if flash was fired
// (uint8) - value of FlashMode
type FlashMode struct {
	bool
	uint8
}

// NewFlashMode - Create new Flash Mode
func NewFlashMode(m int) FlashMode {
	mode := uint8(m)
	return FlashMode{flashBoolValues[mode], mode}
}

// NoFlashFired - Flash was not fired, false and 0
var NoFlashFired = FlashMode{false, 0}

// FlashValues -
// Derived from https://sno.phy.queensu.ca/~phil/exiftool/TagNames/EXIF.html#Flash (23/09/2019)
var FlashValues = map[uint8]string{
	0:  "No Flash",
	1:  "Fired",
	5:  "Fired, Return not detected",
	7:  "Fired, Return detected",
	8:  "On, Did not fire",
	9:  "On, Fired",
	13: "On, Return not detected",
	15: "On, Return detected",
	16: "Off, Did not fire",
	20: "Off, Did not fire, Return not detected",
	24: "Auto, Did not fire",
	25: "Auto, Fired",
	29: "Auto, Fired, Return not detected",
	31: "Auto, Fired, Return detected",
	32: "No flash function",
	48: "Off, No flash function",
	65: "Fired, Red-eye reduction",
	69: "Fired, Red-eye reduction, Return not detected",
	71: "Fired, Red-eye reduction, Return detected",
	73: "On, Red-eye reduction",
	77: "On, Red-eye reduction, Return not detected",
	79: "On, Red-eye reduction, Return detected",
	80: "Off, Red-eye reduction",
	88: "Auto, Did not fire, Red-eye reduction",
	89: "Auto, Fired, Red-eye reduction",
	93: "Auto, Fired, Red-eye reduction, Return not detected",
	95: "Auto, Fired, Red-eye reduction, Return detected",
}

// flashBoolValues -
// (bool) - true if the flash was fired
var flashBoolValues = map[uint8]bool{
	0:  false,
	1:  true,
	5:  true,
	7:  true,
	8:  false,
	9:  true,
	13: true,
	15: true,
	16: false,
	20: false,
	24: false,
	25: true,
	29: true,
	31: true,
	32: false,
	48: false,
	65: true,
	69: true,
	71: true,
	73: true,
	77: true,
	79: true,
	80: false,
	88: false,
	89: true,
	93: true,
	95: true,
}
