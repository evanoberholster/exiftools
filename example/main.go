package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/evanoberholster/exif/exif"
	"github.com/evanoberholster/exif/mknote"
)

func main() {
	fname := "../../test/img/15.jpg"
	//fname := "../exif/samples/"

	f, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}

	metadata(f)
}

type Metadata struct {
	FileSize					int64 		`json:"FileSize"`
	MIMEType					string 		`json:"MIMEType"`
	ImageWidth              	int     	`json:"ImageWidth"`
	ImageHeight             	int     	`json:"ImageHeight"`
	CameraMake                  string  	`json:"CameraMake"`
	CameraModel                 string  	`json:"CameraModel"`
	CameraSerial				string 		`json:"CameraSerial"`
	LensModel					string 		`json:"LensModel"`
	LensSerial					string 		`json:"LensSerial"`
	Artist                    	string  	`json:"Artist"`
	Copyright                 	string  	`json:"Copyright"`
	Aperture 					float32  	`json:"Aperture"`
	ShutterSpeed 				Rational	`json:"ShutterSpeed"`
	ISOSpeed                   	int     	`json:"ISO"`
	ExposureBias 				Rational	`json:"ExposureBias"`
	ExposureProgram           	int 	  	`json:"ExposureProgram"`
	MeteringMode				int 		`json:"MeteringMode"`	
	Orientation             	int 	  	`json:"Orientation"`
	Flash						int 		`json:"Flash"`
	Software					string 		`json:"Software"`
	FocalLength 				float32		`json:"FocalLength"` // mm
	GPSLatitude 				float64 	`json:"GPSLatitude"`
	GPSLongitude 				float64 	`json:"GPSLongitude"`
	GPSAltitude 				float32		`json:"GPSAltitude"`
	GPSTimeStamp 				time.Time   `json:"GPSTimeStamp"`
	DateTimeOriginal			time.Time 	`json:"DateTimeOriginal"`
}

func NewMetadata(file *os.File) *Metadata {
	// Check fileSize
	fi, err := file.Stat()
	if err != nil {
	    log.Println(err)
	}

	// Check file MimeType

	return &Metadata{
		FileSize: fi.Size(),
		MIMEType: "image/jpeg",
	}
}

// EXIF
func metadata(f *os.File) {
	// Optionally register camera makenote data parsing - currently Nikon and
	// Canon are supported.
	exif.RegisterParsers(mknote.All...)

	x, err := exif.Decode(f)
	if err != nil {
		log.Println(err)
	}

	// Create Metadata
	m := NewMetadata(f)

	m.GPSTimeStamp, _ = x.GPSTimeStamp()
	m.DateTimeOriginal, _ = x.DateTime()

	// Image Size
	m.ImageWidth, m.ImageHeight = GetImageSize(x)

	m.CameraMake, m.CameraModel, m.CameraSerial = GetCameraInfo(x)
	m.LensModel, m.LensSerial = GetLensInfo(x)
	m.GPSLatitude, m.GPSLongitude, _ = x.LatLong()
	m.GPSAltitude, err = x.GPSAltitude()
	log.Println(err)

	m.Aperture = GetAperture(x)
	m.ShutterSpeed = GetShutterSpeed(x)
	m.ExposureBias = GetExposureBias(x)
	m.ISOSpeed = GetISOSpeed(x)
	m.FocalLength, _ = FocalLength(x)

	log.Println(m)

	focal, _ := x.Get(exif.FocalLength)
	numer, denom, _ := focal.Rat2(0) // retrieve first (only) rat. value
	fmt.Printf("%v/%v\n", numer, denom)

	imageWidth, err := x.Get(exif.MeteringMode)
	imageHeight, err := x.Get(exif.Software)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(imageWidth, imageHeight)
	}

	// Two convenience functions exist for date/time taken and GPS coords:
	tm, _ := x.DateTime()
	fmt.Println("Taken: ", tm)

}

func FocalLength(x *exif.Exif) (float32, error) {
	tag, err := x.Get(exif.FocalLength)
	if err != nil {
		return 0, fmt.Errorf("Cannot parse Focal Length: %v", err)
	}
	num, denom, _ := tag.Rat2(0)
	log.Println(tag)
	return float32(num)/float32(denom), nil
}

func GetShutterSpeed(x *exif.Exif) Rational {
	r := Rational{}
	a, err := x.Get(exif.ExposureTime)
	if err != nil { 
		return r.Set(0,0) // default
	}
	num, denom, _ := a.Rat2(0)
	return r.Set(num, denom)
}

func GetExposureBias(x *exif.Exif) Rational {
	r := Rational{}
	e, err := x.Get(exif.ExposureBiasValue)
	if err != nil {
		return r.Set(0,0)
	}
	num, denom, _ := e.Rat2(0)
	return r.Set(num,denom)
}

func GetAperture(x *exif.Exif) float32 {
	a, err := x.Get(exif.FNumber)
	if err != nil {
		log.Println(err)
		return 0.0 // default
	} 
	num, denom, _ := a.Rat2(0)
	return float32(num)/float32(denom)
}

func GetISOSpeed(x *exif.Exif) int {
	a, err := x.Get(exif.ISOSpeedRatings)
	if err != nil {
		log.Println(err)
		return 0.0 // default
	} 
	iso, _ := a.Int(0)
	return iso
}

func GetImageSize(x *exif.Exif) (int, int) {
	w, err := x.Get(exif.ImageWidth)
	l, err := x.Get(exif.ImageLength)
	if err != nil {
		log.Println(err)
		return 0,0
	}
	width, _ := w.Int(0)
	length, _ := l.Int(0)
	return width, length
}

func GetLensInfo(x *exif.Exif) (string, string) {
	// LensModel 
	lensModel := ""
	lm, err := x.Get(exif.LensModel)
	if err != nil {
		log.Println(err)
	} else {
		lensModel, _ = lm.StringVal()
	}

	// LensSerial 
	lensSerialNum := ""
	lsn, err := x.Get(exif.LensSerialNumber)
	if err != nil {
		log.Println(err)
	} else {
		lensSerialNum, _ = lsn.StringVal()
	}
	return lensModel, lensSerialNum
}

func GetCameraInfo(x *exif.Exif) (string, string, string) {
	// CameraMake
	cameraMake := ""
	cma, err := x.Get(exif.Make)
	if err != nil {
		log.Println(err)
	} else {
		cameraMake, _ = cma.StringVal()
	}
	// CameraModel
	cameraModel := ""
	cm, err := x.Get(exif.Model)
	if err != nil {
		log.Println(err)
	} else {
		cameraModel, _ = cm.StringVal()
	}

	// CameraSerial 
	cameraSerial := ""
	cs, err := x.Get(exif.SerialNumber)
	if err != nil {
		log.Println(err)
	} else {
		cameraSerial, _ = cs.StringVal()
	}

	return cameraMake,cameraModel,cameraSerial
}