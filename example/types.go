package main

import (
	"github.com/evanoberholster/exif/exif"
	"github.com/evanoberholster/exif/tiff"
	"github.com/evanoberholster/exif/mknote"
)

type CameraSettingsField map[int]string

var Canon_ContinuousDriveValues = map[int]string{
	0: "Single",
	1: "Continuous", 
	2: "Movie",
	3: "Continuous, Speed Priority", 
	4: "Continuous, Low", 
	5: "Continuous, High", 
	6: "Silent Single", 
	9: "Single, Silent", 
	10: "Continuous, Silent",
}
	
var Canon_FocusModeValues = map[int]string{
	0:"One-shot AF", 
	1:"AI Servo AF", 
	2:"AI Focus AF", 
	3:"Manual Focus (3)", 
	4:"Single", 
	5:"Continuous",	  	
	6:"Manual Focus (6)",
	16:"Pan Focus", 
	256:"AF + MF", 
	512:"Movie Snap Focus", 
	519:"Movie Servo AF",
}

var Canon_ExposureModeValues = map[int]string{
	0:"Easy", 
	1:"Program AE",
	2:"Shutter speed priority AE",
	3:"Aperture-priority AE",
	4:"Manual",
	5:"Depth-of-field AE",
	6:"M-Dep",
	7:"Bulb",
	8:"Flexible-priority AE",
}

var Canon_RecordModeValues = map[int]string{
	1:"JPEG", 
	2:"CRW+THM", 
	3:"AVI+THM",
	4:"TIF",	  	
	5:"TIF+JPEG", 
	6:"CR2", 
	7:"CR2+JPEG", 
	9:"MOV",	  	
	10:"MP4", 
	11:"CRM", 
	12:"CR3", 
	13:"CR3+JPEG",
}

var Canon_MeteringModeValues = map[int]string{
	0:"Default",
	1:"Spot", 
	2:"Average", 
	3:"Evaluative",
	4:"Partial", 
	5:"Center-weighted average",
}

var Canon_AESettingValues = map[int]string{
	0:"Normal AE", 
	1:"Exposure Compensation", 
	2:"AE Lock", 
	3:"AE Lock + Exposure Compensation", 
	4:"No AE",
}

const (
	Canon_ContinuousDrive	 int = 5
	Canon_FocusMode	 		 int = 7
	Canon_RecordMode		 int = 9
	Canon_MeteringMode		 int = 17
	Canon_ExposureMode		 int = 20
	Canon_AESetting 		 int = 33
)

var Canon_CameraSettingsFields = map[int]CameraSettingsField{
	Canon_ContinuousDrive: Canon_ContinuousDriveValues,
	Canon_FocusMode: Canon_FocusModeValues,
	Canon_RecordMode: Canon_RecordModeValues,
	Canon_MeteringMode: Canon_MeteringModeValues,
	Canon_ExposureMode: Canon_ExposureModeValues,
	Canon_AESetting: Canon_AESettingValues,
}

type CameraSettings struct {
	RecordMode			string 		`json:"RecordMode"` 
	FocusMode			string 		`json:"FocusMode"`
	ExposureMode 		string 		`json:"ExposureMode"`
	ContinuousDrive		string 		`json:"ContinuousDrive"`
	MeteringMode		string 		`json:"MeteringMode"`
	//AESetting			string 		`json:"AESetting"`
}

func ProcessCameraSettingsFields(tag *tiff.Tag, i int) (string) {
	a, err := tag.Int(i)
	if err != nil {
		return ""
	}
	return Canon_CameraSettingsFields[i][a]
}

func Canon_CameraSettings(x *exif.Exif) (CameraSettings, error) {
	var c CameraSettings
	tag, err := x.Get(mknote.Canon_CameraSettings) 
	if err != nil {
		return c, err
	}
	
	c.ContinuousDrive = ProcessCameraSettingsFields(tag, Canon_ContinuousDrive)
	c.RecordMode = ProcessCameraSettingsFields(tag, Canon_RecordMode)
	c.FocusMode = ProcessCameraSettingsFields(tag, Canon_FocusMode)
	c.ExposureMode = ProcessCameraSettingsFields(tag, Canon_ExposureMode)
	c.MeteringMode = ProcessCameraSettingsFields(tag, Canon_MeteringMode)
	//c.AESetting = ProcessCameraSettingsFields(tag, Canon_AESetting)

	return c, nil
}

//

func ISOSpeed(x *exif.Exif) (int, error) {
	a, err := x.Get(exif.ISOSpeedRatings)
	if err != nil { return 0, err }
	return a.Int(0)
}

func Aperture(x *exif.Exif) (float32, error) {
	a, err := x.Get(exif.FNumber)
	if err != nil {
		return 0.0, err
	} 
	num, denom, err := a.Rat2(0)
	return float32(num)/float32(denom), err
}

func ShutterSpeed(x *exif.Exif) (Rational, error) {
	r := Rational{}
	a, err := x.Get(exif.ExposureTime)
	if err != nil { 
		return r.Set(0,0), err
	}
	num, denom, err := a.Rat2(0)
	return r.Set(num, denom), err
}

func ExposureBias(x *exif.Exif) (Rational, error) {
	r := Rational{}
	e, err := x.Get(exif.ExposureBiasValue)
	if err != nil {
		return r.Set(0,0), err
	}
	num, denom, err := e.Rat2(0)
	return r.Set(num,denom), err
}

type Rational struct {
	num, denom		int64 // Numerator and Denominator
}

func (r Rational) Set(n, d int64) Rational{
	r.num = n
	r.denom = d
	return r
}

func FetchString(x *exif.Exif, tagLabel exif.FieldName) (string, error) {
	a, err := x.Get(tagLabel)
	if err != nil {
		return "", err
	}
	return a.StringVal()
}

/////



