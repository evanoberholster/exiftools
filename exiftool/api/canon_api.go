package api

import (
	"fmt"

	"github.com/evanoberholster/exiftools/exiftool/tags/mknote"
)

// CanonMakerNote API constants
const (
	CanonMakernotesIfdString = "IFD/Exif/Makernotes.Canon"
)

// CanonCameraSettings - Canon Makernote Camera Settings
// Incomplete
type CanonCameraSettings struct {
	Macromode         bool                        // [1]
	SelfTimer         bool                        // [2]
	ContinuousDrive   mknote.CanonContinuousDrive // [5]
	FocusMode         mknote.CanonFocusMode       // [7]
	MeteringMode      mknote.CanonMeteringMode    // [17]
	FocusRange        mknote.CanonFocusRange      // [18]
	CanonExposureMode mknote.CanonExposureMode    // [20]
	MaxFocalLength    int16                       // [23]
	MinFocalLength    int16                       // [24]
	//FocalUnits        int16                       // [25]
	//FocusContinuous   mknote.CanonFocusContinous  // [32]
	//SpotMeteringMode  bool                        // [39]
	//AESetting         mknote.CanonAESetting       // [33]
}

// CanonShotInfo - Canon Makernote Shot Information
// Incomplete
type CanonShotInfo struct {
	CameraTemperature      int16 // [12] 	CameraTemperature 	int16s 	(newer EOS models only)
	FlashExposureComp      int16 // [15] 	FlashExposureComp 	int16s
	AutoExposureBracketing int16 // [16] 	AutoExposureBracketing 	int16s
	AEBBracketValue        int16 // [17] 	AEBBracketValue 	int16s
	SelfTimer              int16 // 29 	SelfTimer2 	int16s
	//FocusDistance          mknote.FocusDistance // 19 	FocusDistanceUpper 	int16u // 20 	FocusDistanceLower 	int16u
}

// CanonFileInfo - Canon Makernote File Information
type CanonFileInfo struct {
	FocusDistance     mknote.FocusDistance    // 20 	FocusDistanceUpper 	int16u // 21 	FocusDistanceLower 	int16u
	BracketMode       mknote.CanonBracketMode // 3 	BracketMode 	int16s
	BracketValue      int16                   // 4 	BracketValue 	int16s
	BracketShotNumber int16                   // 5 	BracketShotNumber 	int16s
	LiveViewShooting  bool                    // 19 	LiveViewShooting 	int16s (bool)
}

// CanonAFInfo - Canon Makernote Autofocus Information
type CanonAFInfo struct {
	AFAreaMode    mknote.CanonAFAreaMode
	NumAFPoints   uint16
	ValidAFPoints uint16
	AFPoints      []mknote.AFPoint
	InFocus       []int
	Selected      []int
}

// CanonCameraSettings convenience func. "IFD/Exif/Makernotes.Canon" CanonCameraSettings
// Canon Camera Settings from the Makernote
func (res Results) CanonCameraSettings() (CanonCameraSettings, error) {
	ii, err := res.GetTag(CanonMakernotesIfdString, mknote.CanonCameraSettings).Short()
	if len(ii) < 24 || err != nil {
		return CanonCameraSettings{}, err
	}
	return CanonCameraSettings{
		Macromode:         intToBool(ii[1]),
		SelfTimer:         intToBool(ii[2]),
		ContinuousDrive:   mknote.CanonContinuousDrive(ii[5]),
		FocusMode:         mknote.CanonFocusMode(ii[7]),
		MeteringMode:      mknote.CanonMeteringMode(ii[17]),
		FocusRange:        mknote.CanonFocusRange(ii[18]),
		CanonExposureMode: mknote.CanonExposureMode(ii[20]),
	}, nil
}

// CanonShotInfo convenience func. "IFD/Exif/Makernotes.Canon" CanonShotInfo
// Canon Camera Shot Info from the Makernote
func (res Results) CanonShotInfo() (CanonShotInfo, error) {
	si, err := res.GetTag(CanonMakernotesIfdString, mknote.CanonShotInfo).Short()
	if len(si) < 24 || err != nil {
		return CanonShotInfo{}, err
	}

	return CanonShotInfo{
		SelfTimer:              int16(si[29]) / 10,
		CameraTemperature:      canonTempConv(si[12]),
		FlashExposureComp:      int16(si[15]),
		AutoExposureBracketing: int16(si[16]),
		AEBBracketValue:        canonEv(int16(si[17])),
		//FocusDistance:          mknote.NewFocusDistance(si[19], si[20]),
	}, nil
}

// CanonFileInfo convenience func. "IFD/Exif/Makernotes.Canon" CanonFileInfo
// Canon Camera File Info from the Makernote
func (res Results) CanonFileInfo() (CanonFileInfo, error) {
	fi, err := res.GetTag(CanonMakernotesIfdString, mknote.CanonFileInfo).Short()
	if len(fi) < 21 || err != nil {
		return CanonFileInfo{}, err
	}
	return CanonFileInfo{
		FocusDistance:     mknote.NewFocusDistance(fi[20], fi[21]),
		BracketMode:       mknote.CanonBracketMode(fi[3]),
		BracketValue:      canonEv(int16(fi[4])),
		BracketShotNumber: int16(fi[5]),
		LiveViewShooting:  intToBool(fi[19]),
	}, nil
}

// CanonAFInfo -
// Canon Camera AutoFocus Information from the Makernote
func (res Results) CanonAFInfo() (afInfo CanonAFInfo, err error) {
	defer func() {
		if state := recover(); state != nil {
			err = state.(error)
		}
	}()
	af, err := res.GetTag(CanonMakernotesIfdString, mknote.CanonAFInfo2).Short()
	if len(af) < 8 || err != nil {
		panic(ErrParseTag)
	}

	afInfo = CanonAFInfo{
		AFAreaMode:    mknote.CanonAFAreaMode(af[1]),
		NumAFPoints:   af[2],
		ValidAFPoints: af[3],
	}

	if infocus, selected, err := pointsInFocus(af, int(afInfo.ValidAFPoints)); err == nil {
		afInfo.InFocus = infocus
		afInfo.Selected = selected
	} else {
		panic(err)
	}

	validPoints := int(af[3])
	// AFPoints
	afInfo.AFPoints = make([]mknote.AFPoint, validPoints)
	xAdjust := int16(af[4] / 2) // Adjust x-axis
	yAdjust := int16(af[5] / 2) // Adjust y-axis

	for i := 0; i < validPoints; i++ { // Start at an offset of 8
		offset := 8 + i
		w := int16(af[offset])
		h := int16(af[offset+validPoints])
		x := int16(af[offset+(2*validPoints)]) + xAdjust - (w / 2)
		y := int16(af[offset+(3*validPoints)]) + yAdjust - (h / 2)
		afInfo.AFPoints[i] = mknote.NewAFPoint(w, h, x, y)
	}

	return afInfo, nil
}

func pointsInFocus(af []uint16, validPoints int) (inFocus []int, selected []int, err error) {
	var count int
	// NumAFPoints may be 7, 9, 11, 19, 31, 45 or 61, depending on the camera model.
	switch validPoints {
	case 7:
		count = 1 // 1
	case 9, 11:
		count = 1 // 1
	case 19, 31:
		count = 2 // 2
	case 45:
		count = 3 // 3
	case 61:
		count = 4 // 4
	case 65:
		count = 5 // 5
	default:
		panic(fmt.Errorf("Error parsing AFPoints from Canon Makernote. Expected 7, 9, 11, 19, 31, 45 or 61 got %d", validPoints))
	}
	off := 8 + (validPoints * 4)
	inFocus = decodeBits(af[off:off+count], 16)
	selected = decodeBits(af[off+count:off+count+count], 16)
	return
}

// decodeBits - ported from Phil Harvey's exiftool
// Updated May-10-2020
// https://github.com/exiftool/exiftool/lib/Image/ExifTool.pm
func decodeBits(vals []uint16, bits int) (list []int) {
	var num int
	var n int
	for _, a := range vals {
		for i := 0; i < bits; i++ {
			n = i + num
			if a&(1<<i) > 0 {
				list = append(list, n)
			}
		}
		num += bits
	}
	return
}

// canonTempConv - ported from Phil Harvey's exiftool
// Updated May-10-2020
// https://github.com/exiftool/exiftool/lib/Image/ExifTool/Canon.pm
func canonTempConv(val uint16) int16 {
	if val <= 0 {
		return 0
	}
	return int16(val) - 128
}

//# Print exposure compensation fraction
//sub PrintFraction($)
//{
//    my $val = shift;
//    my $str;
//    if (defined $val) {
//        $val *= 1.00001;    # avoid round-off errors
//        if (not $val) {
//            $str = '0';
//        } elsif (int($val)/$val > 0.999) {
//            $str = sprintf("%+d", int($val));
//        } elsif ((int($val*2))/($val*2) > 0.999) {
//            $str = sprintf("%+d/2", int($val * 2));
//        } elsif ((int($val*3))/($val*3) > 0.999) {
//            $str = sprintf("%+d/3", int($val * 3));
//        } else {
//            $str = sprintf("%+.3g", $val);
//        }
//    }
//    return $str;
//}

func intToBool(i uint16) bool {
	return i == 1
}

// canonEv - ported from Phil Harvey's exiftool
// Updated May-10-2020
// https://github.com/exiftool/exiftool/lib/Image/ExifTool/Canon.pm
func canonEv(val int16) int16 {
	var sign int16
	if val < 0 {
		val = -val
		sign = -1
	} else {
		sign = 1
	}
	frac := val & 0x1f
	val -= frac
	// Convert 1/3 and 2/3 codes
	if frac == 0x0c {
		frac = 0x20 / 3
	} else if frac == 0x14 {
		frac = 0x40 / 3
	}
	return sign * (val + frac) / 0x20
}
