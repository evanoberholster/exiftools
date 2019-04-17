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
	fname := "../../test/img/2.CR2"

	f, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}

	metadata(f)
}

type Metadata struct {
	FileSize					int64 		`json:"FileSize"` 		  // OK
	MIMEType					string 		`json:"MIMEType"`
	ImageWidth              	int     	`json:"ImageWidth"`
	ImageHeight             	int     	`json:"ImageHeight"`
	CameraMake                  string  	`json:"CameraMake"`		  // OK
	CameraModel                 string  	`json:"CameraModel"`	  // OK
	CameraSerial				string 		`json:"CameraSerial"`	  // OK
	LensModel					string 		`json:"LensModel"`		  // OK
	LensSerial					string 		`json:"LensSerial"` 	  // OK
	Artist                    	string  	`json:"Artist"`
	Copyright                 	string  	`json:"Copyright"`
	Aperture 					float32  	`json:"Aperture"` 		  // OK
	ShutterSpeed 				Rational	`json:"ShutterSpeed"` 	  // OK
	ISOSpeed                   	int     	`json:"ISO"`  			  // OK
	ExposureBias 				Rational	`json:"ExposureBias"` 	  // OK
	ExposureProgram           	string 	  	`json:"ExposureProgram"`  // OK
	MeteringMode				int 		`json:"MeteringMode"`	
	Orientation             	int 	  	`json:"Orientation"`
	Flash						int 		`json:"Flash"`
	FocalLength 				float32		`json:"FocalLength"` 	  // OK
	FocalLengthEqv 				float32		`json:"FocalLengthEqv"` // mm
	GPSLatitude 				float64 	`json:"GPSLatitude"`	  // OK
	GPSLongitude 				float64 	`json:"GPSLongitude"`	  // OK
	GPSAltitude 				float32		`json:"GPSAltitude"`	  // OK
	GPSTimeStamp 				time.Time   `json:"GPSTimeStamp"`	  // OK
	DateTimeOriginal			time.Time 	`json:"DateTimeOriginal"` // OK
	ModifyTimeStamp 			time.Time 	`json:"ModifyTimeStamp"`
	XMPKeywords					[]string 	`json:"XMPKeywords"`
	XMPRating 					int 		`json:"XMPRating"`
	XMPColorLabel 				string 		`json:"XMPColorLabel"`
	XMPRights					string 		`json:"XMPRights"` 
	Software					string 		`json:"Software"`		  // OK
	CameraSettings 				CameraSettings 	`json:"CameraSettings"`
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
		log.Fatal(err)
	}

	// Create Metadata
	m := NewMetadata(f)

	m.GPSAltitude, err = x.GPSAltitude()
	m.GPSTimeStamp, err = x.GPSTimeStamp()
	if err != nil {
		log.Println(err)
	}

	m.DateTimeOriginal, _ = x.DateTime()

	m.FocalLength, err = x.FocalLength()

	m.FocalLengthEqv, _ = FocalLengthIn35mmFilm(x)
	m.ExposureProgram, err = ExposureProgram(x)

	// Image Size
	m.ImageWidth, m.ImageHeight = GetImageSize(x)

	m.CameraMake, _ = FetchString(x, exif.Make)
	m.CameraModel, _ = FetchString(x, exif.Model)
	m.CameraSerial, _ = FetchString(x, exif.SerialNumber)
	
	m.LensModel, _ = FetchString(x, exif.LensModel)
	m.LensSerial, _ = FetchString(x, exif.LensSerialNumber)

	m.GPSLatitude, m.GPSLongitude, _ = x.LatLong()
	
	m.ISOSpeed, _ = ISOSpeed(x)
	m.Aperture, _ = Aperture(x)
	m.ShutterSpeed, _ = ShutterSpeed(x)
	m.ExposureBias, _ = ExposureBias(x)
	
	m.Software, _ = x.Software()
	c, err := Canon_CameraSettings(x)
	if err == nil {
		m.CameraSettings = c
	}

	//log.Println(x)
	log.Println("\n",m)

	// Two convenience functions exist for date/time taken and GPS coords:
	tm, _ := x.DateTime()
	fmt.Println("Taken: ", tm)

}

func ExposureProgram(x *exif.Exif) (string, error) {
	tag, err := x.Get(exif.ExposureProgram)
	if err != nil {
		return "", err
	}
	_, err = tag.Int(0)
	if err != nil {
		return "", err
	}
	//log.Println(tag, tag.Count, tag.Type)
	return "", nil
}

func FocalLengthIn35mmFilm(x *exif.Exif) (float32, error) {
	_, err := x.Get(exif.FocalLengthIn35mmFilm)
	if err != nil {
		return 0.0, err
	}

	//log.Println(tag.Type, tag.Count, tag)
	return 0.0, nil
}

func GetImageSize(x *exif.Exif) (int, int) {
	w, err := x.Get(exif.ImageWidth)
	l, err := x.Get(exif.ImageLength)
	if err != nil {
		log.Println(err)
		return 0,0
	}
	//log.Println(w.Type, l.Type, w.Count)
	width, _ := w.Int(0)
	length, _ := l.Int(0)
	return width, length
}