package main

import (
	"fmt"
	"log"
	"os"

	"github.com/evanoberholster/exif/exif"
	"github.com/evanoberholster/exif/mknote"
)

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
	ExposureProgram           	int 	  	`json:"ExposureProgram"`
	MeteringMode				int 		`json:"MeteringMode"`	
	Orientation             	int 	  	`json:"Orientation"`
	Flash						int 		`json:"Flash"`
	Software					string 		`json:"Software"`
	FocalLength 				float64		`json:"FocalLength"` // mm
}

func NewMetadata() *Metadata {
	return &Metadata{}
}

// MIME

// EXIF
func metadata(x *exif.Exif) {
	m := NewMetadata()

	m.ISOSpeed = 0


	m.Aperture = GetAperture(x)
	m.ShutterSpeed = GetShutterSpeed(x)

	log.Println(m)
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

func GetAperture(x *exif.Exif) float32 {
	a, err := x.Get(exif.FNumber)
	if err != nil {
		log.Println(err)
		return 0.0 // default
	} 
	num, denom, _ := a.Rat2(0)
	return float32(num)/float32(denom)
}

func GetLensInfo(x *exif.Exif) (string, string) {
	// LensModel 
	// LensSerial
}

func GetCameraInfo(x *exif.Exif) (string, string, string) {
	// CameraMake
	// CameraModel
	// CameraSerial
}

// JPEG Size

func main() {
	fname := "../../test/img/4.CR2"

	f, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}

	// Optionally register camera makenote data parsing - currently Nikon and
	// Canon are supported.
	exif.RegisterParsers(mknote.All...)

	x, err := exif.Decode(f)
	if err != nil {
		log.Println(err)
	}

	metadata(x)

	camModel, _ := x.Get(exif.Model) // normally, don't ignore errors!
	fmt.Println(camModel)

	focal, _ := x.Get(exif.FocalLength)
	numer, denom, _ := focal.Rat2(0) // retrieve first (only) rat. value
	fmt.Printf("%v/%v\n", numer, denom)

	lensModel, err := x.Get(mknote.SerialNumber)
	
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(lensModel)
	}

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

	lat, long, _ := x.LatLong()
	fmt.Println("lat, long: ", lat, ", ", long)
}