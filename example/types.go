package main

import (
	"time"
	"fmt"
	"trimmer.io/go-xmp/xmp"
	_ "trimmer.io/go-xmp/models"
	"trimmer.io/go-xmp/models/dc"
	"trimmer.io/go-xmp/models/xmp_base"

	"github.com/evanoberholster/exif/exif"
	"github.com/evanoberholster/exif/tiff"
	"github.com/evanoberholster/exif/mknote"
)

var Exif_ExposureModeValues = map[int]string{
	0:"Not Defined", 
	1:"Manual", 
	2:"Program AE", 
	3:"Aperture-priority AE", 
	4:"Shutter speed priority AE", 
	5:"Creative (Slow speed)", 
	6:"Action (High speed)", 
	7:"Portrait", 
	8:"Landscape", 
	9:"Bulb",
}

var Exif_MeteringModeValues = map[int]string{
	0:"Unknown", 
	1:"Average", 
	2:"Center-weighted average", 
	3:"Spot", 
	4:"Multi-spot", 
	5:"Multi-segment", 
	6:"Partial", 
	255:"Other",
}

var Exif_OrientationValues = map[int]string {
	1:"Horizontal (normal)", 
	2:"Mirror horizontal", 
	3:"Rotate 180", 
	4:"Mirror vertical", 
	5:"Mirror horizontal and rotate 270 CW",
	6:"Rotate 90 CW", 
	7:"Mirror horizontal and rotate 90 CW", 
	8:"Rotate 270 CW",
}

var Exif_FlashBoolValues = map[int]bool{
	0:false,
	1:true,
	5:true,
	7:true,
	8:false,
	9:true,
	13:true,
	15:true,
	16:false,
	20:false,
	24:false,
	25:true,
	29:true,
	31:true,
	32:false,
	48:false,
	65:true,
	69:true,
	71:true,
	73:true,
	77:true,
	79:true,
	80:false,
	88:false,
	89:true,
	93:true,
	95:true,
}

var Exif_FlashValues = map[int]string{
	0:"No Flash",
	1:"Fired",
	5:"Fired, Return not detected",
	7:"Fired, Return detected",
	8:"On, Did not fire",
	9:"On, Fired",
	13:"On, Return not detected",
	15:"On, Return detected",
	16:"Off, Did not fire",
	20:"Off, Did not fire, Return not detected",
	24:"Auto, Did not fire",
	25:"Auto, Fired",
	29:"Auto, Fired, Return not detected",
	31:"Auto, Fired, Return detected",
	32:"No flash function",
	48:"Off, No flash function",
	65:"Fired, Red-eye reduction",
	69:"Fired, Red-eye reduction, Return not detected",
	71:"Fired, Red-eye reduction, Return detected",
	73:"On, Red-eye reduction",
	77:"On, Red-eye reduction, Return not detected",
	79:"On, Red-eye reduction, Return detected",
	80:"Off, Red-eye reduction",
	88:"Auto, Did not fire, Red-eye reduction",
	89:"Auto, Fired, Red-eye reduction",
	93:"Auto, Fired, Red-eye reduction, Return not detected",
	95:"Auto, Fired, Red-eye reduction, Return detected",
}

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
}

func ProcessCameraSettingsFields(tag *tiff.Tag, i int) (string) {
	a, err := tag.Int(i)
	if err != nil {
		return ""
	}
	return Canon_CameraSettingsFields[i][a]
}

func Canon_CameraSettings(x *exif.Exif) (CameraSettings, error) {
	c := CameraSettings{}
	tag, err := x.Get(mknote.Canon_CameraSettings) 
	if err != nil {
		return c, err
	}
	
	c.ContinuousDrive = ProcessCameraSettingsFields(tag, Canon_ContinuousDrive)
	c.RecordMode = ProcessCameraSettingsFields(tag, Canon_RecordMode)
	c.FocusMode = ProcessCameraSettingsFields(tag, Canon_FocusMode)
	c.ExposureMode = ProcessCameraSettingsFields(tag, Canon_ExposureMode)
	c.MeteringMode = ProcessCameraSettingsFields(tag, Canon_MeteringMode)

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

func ExposureMode(x *exif.Exif) (string, error) {
	tag, err := x.Get(exif.ExposureProgram)
	if err != nil {
		return "", err
	}
	v, err := tag.Int(0)
	if err != nil {
		return "", err
	}
	return Exif_ExposureModeValues[v], nil
}

func MeteringMode(x *exif.Exif) (string, error) {
	tag, err := x.Get(exif.MeteringMode)
	if err != nil {
		return "", err
	}
	v, err := tag.Int(0)
	if err != nil {
		return "", err
	}
	return Exif_MeteringModeValues[v], nil
}

func Orientation(x *exif.Exif) (int, error) {
	i, err := x.Get(exif.Orientation)
	if err != nil {
		return 1, err
	}
	return i.Int(0)
}

func Flash(x *exif.Exif) (ExifFlash, error) {
	tag, err := x.Get(exif.Flash)
	if err != nil {
		return ExifFlash{}, err
	}
	v, err := tag.Int(0)
	if err != nil {
		return ExifFlash{}, err
	}

	return ExifFlash{
			Fired:Exif_FlashBoolValues[v],
			Description:Exif_FlashValues[v],
			Value:int(v),
		}, nil
}

type ExifFlash struct {
	Fired 			bool
	Description 	string
	Value			int
}

type Rational struct {
	Num, Denom		int64 // Numerator and Denominator
	Value 			string		
}

func (r Rational) Set(n, d int64) Rational{
	r.Num = n
	r.Denom = d
	r.Value = fmt.Sprintf("%d/%d",n,d)
	return r
}

func FetchString(x *exif.Exif, tagLabel exif.FieldName) (string, error) {
	a, err := x.Get(tagLabel)
	if err != nil {
		return "", err
	}
	return a.StringVal()
}

// Metadata
type Metadata struct {
	FileSize		int64 		`json:"FileSize"` 		  
	MIMEType		string 		`json:"MIMEType"`
	DublinCore		DublinCore 	`json:"DublinCore"`
	XmpBase 		XmpBase 	`json:"XmpBase"`
	Exif 			Exif 		`json:"Exif"`
}

// Exig

type Exif struct {
	ImageWidth              	int     		`json:"ImageWidth"`
	ImageHeight             	int     		`json:"ImageHeight"`
	CameraMake                  string  		`json:"CameraMake"`		  // OK
	CameraModel                 string  		`json:"CameraModel"`	  // OK
	CameraSerial				string 			`json:"CameraSerial"`	  // OK
	LensModel					string 			`json:"LensModel"`		  // OK
	LensSerial					string 			`json:"LensSerial"` 	  // OK
	Artist                    	string  		`json:"Artist"` 		  // OK
	Copyright                 	string  		`json:"Copyright"` 		  // OK
	Aperture 					float32  		`json:"Aperture"` 		  // OK
	ShutterSpeed 				Rational		`json:"ShutterSpeed"` 	  // OK
	ISOSpeed                   	int     		`json:"ISO"`  			  // OK
	ExposureBias 				Rational		`json:"ExposureBias"` 	  // OK
	ExposureMode	          	string 	  		`json:"ExposureMode"`  	  // OK
	MeteringMode				string 			`json:"MeteringMode"`	  // OK
	Orientation             	int 	  		`json:"Orientation"`
	Flash						ExifFlash		`json:"Flash"` 			  // OK
	FocalLength 				float32			`json:"FocalLength"` 	  // OK
	FocalLengthEqv 				float32			`json:"FocalLengthEqv"`   // mm
	GPSLatitude 				float64 		`json:"GPSLatitude"`	  // OK
	GPSLongitude 				float64 		`json:"GPSLongitude"`	  // OK
	GPSAltitude 				float32			`json:"GPSAltitude"`	  // OK
	GPSTimeStamp 				time.Time   	`json:"GPSTimeStamp"`	  // OK
	DateTimeOriginal			time.Time 		`json:"DateTimeOriginal"` // OK
	ModifyTimeStamp 			time.Time 		`json:"ModifyTimeStamp"`
	Software					string 			`json:"Software"`		  // OK
	ImageDescription			string 			`json:"ImageDescription"`
	CameraSettings 				CameraSettings 	`json:"CameraSettings"`
}

// XMP
///// Modified DublinCore 17/04/2019 https://godoc.org/trimmer.io/go-xmp/models/dc#DublinCore

type DublinCore struct {
    Creator     string  	`xmp:"dc:creator"`
    Description string   	`xmp:"dc:description"`
    Format      string      `xmp:"dc:format"`
    Rights      string   	`xmp:"dc:rights"`
    Source      string      `xmp:"dc:source"`
    Subject     []string 	`xmp:"dc:subject"`
    Title       string   	`xmp:"dc:title"`
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

///// Modified XmpBase 17/04/2019 https://godoc.org/trimmer.io/go-xmp/models/xmp_base#XmpBase

type XmpBase struct {
	CreateDate   time.Time               `xmp:"xmp:CreateDate"`
	CreatorTool  string          		 `xmp:"xmp:CreatorTool"`
	Identifier   string 		         `xmp:"xmp:Identifier"`
	Label        string                  `xmp:"xmp:Label"`
	MetadataDate time.Time               `xmp:"xmp:MetadataDate"`
	ModifyDate   time.Time               `xmp:"xmp:ModifyDate"`
	Rating       int 	                 `xmp:"xmp:Rating"`
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



