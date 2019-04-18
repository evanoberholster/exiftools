package main

import (
	"fmt"
	"log"
	"os"
	"io"
	"time"

	"encoding/json"
	"github.com/TylerBrock/colorjson"

	"github.com/evanoberholster/exif/exif"
	"github.com/evanoberholster/exif/mknote"
	"trimmer.io/go-xmp/xmp"

)

func main() {
	fname := "../../test/img/4.CR2"

	f, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	start := time.Now()
	metadata(f)

	elapsed := time.Since(start)
	log.Println(elapsed)
}

type Exif struct {
	ImageWidth              	int     	`json:"ImageWidth"`
	ImageHeight             	int     	`json:"ImageHeight"`
	CameraMake                  string  	`json:"CameraMake"`		  // OK
	CameraModel                 string  	`json:"CameraModel"`	  // OK
	CameraSerial				string 		`json:"CameraSerial"`	  // OK
	LensModel					string 		`json:"LensModel"`		  // OK
	LensSerial					string 		`json:"LensSerial"` 	  // OK
	Artist                    	string  	`json:"Artist"` 		  // OK
	Copyright                 	string  	`json:"Copyright"` 		  // OK
	Aperture 					float32  	`json:"Aperture"` 		  // OK
	ShutterSpeed 				Rational	`json:"ShutterSpeed"` 	  // OK
	ISOSpeed                   	int     	`json:"ISO"`  			  // OK
	ExposureBias 				Rational	`json:"ExposureBias"` 	  // OK
	ExposureProgram           	string 	  	`json:"ExposureProgram"`  // OK
	MeteringMode				string 		`json:"MeteringMode"`	  // OK
	Orientation             	int 	  	`json:"Orientation"`
	Flash						ExifFlash	`json:"Flash"` 			  // OK
	FocalLength 				float32		`json:"FocalLength"` 	  // OK
	FocalLengthEqv 				float32		`json:"FocalLengthEqv"`   // mm
	GPSLatitude 				float64 	`json:"GPSLatitude"`	  // OK
	GPSLongitude 				float64 	`json:"GPSLongitude"`	  // OK
	GPSAltitude 				float32		`json:"GPSAltitude"`	  // OK
	GPSTimeStamp 				time.Time   `json:"GPSTimeStamp"`	  // OK
	DateTimeOriginal			time.Time 	`json:"DateTimeOriginal"` // OK
	ModifyTimeStamp 			time.Time 	`json:"ModifyTimeStamp"`
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

// XMP
func (m *Metadata) xmpMetadata(f *os.File) (error) {
	f.Seek(0,0)

	bb, err := xmp.ScanPackets(f)
	if (err != nil && err != io.EOF) || len(bb) == 0 {
		return err
	}

	doc := &xmp.Document{}
	if err := xmp.Unmarshal(bb[0], doc); err != nil {
		return err
	}

	m.DublinCore = xmpDublinCore(doc)

	m.XmpBase = xmpBase(doc)

	return nil
}

// EXIF
func (m *Metadata) exifMetadata(f *os.File) (error) {
	m.Exif = Exif{}
	f.Seek(0,0)

	// Optionally register camera makenote data parsing - currently Nikon and
	// Canon are supported.
	exif.RegisterParsers(mknote.All...)

	x, err := exif.Decode(f)
	if err != nil {
		return err
	}

	log.Println(x)

	m.Exif.GPSAltitude, err = x.GPSAltitude()
	m.Exif.GPSTimeStamp, err = x.GPSTimeStamp()
	if err != nil {
		log.Println(err)
	}
	m.Exif.GPSLatitude, m.Exif.GPSLongitude, _ = x.LatLong()

	m.Exif.DateTimeOriginal, _ = x.DateTime()
	m.Exif.FocalLength, err = x.FocalLength()

	m.Exif.FocalLengthEqv, _ = FocalLengthIn35mmFilm(x)
	m.Exif.ExposureProgram, _ = ExposureProgram(x)
	m.Exif.MeteringMode, _ = MeteringMode(x)

	m.Exif.Flash, _ = Flash(x)
	log.Println(m.Exif.Flash)

	// Image Size
	m.Exif.ImageWidth, m.Exif.ImageHeight = GetImageSize(x)

	m.Exif.CameraMake, _ = FetchString(x, exif.Make)
	m.Exif.CameraModel, _ = FetchString(x, exif.Model)
	m.Exif.CameraSerial, _ = FetchString(x, exif.SerialNumber)
	m.Exif.Artist, _ = FetchString(x, exif.Artist)
	m.Exif.Copyright, _ = FetchString(x, exif.Copyright)

	
	m.Exif.LensModel, _ = FetchString(x, exif.LensModel)
	m.Exif.LensSerial, _ = FetchString(x, exif.LensSerialNumber)

	m.Exif.ISOSpeed, _ = ISOSpeed(x)
	m.Exif.Aperture, _ = Aperture(x)
	m.Exif.ShutterSpeed, _ = ShutterSpeed(x)
	m.Exif.ExposureBias, _ = ExposureBias(x)
	
	m.Exif.Software, _ = x.Software()
	c, err := Canon_CameraSettings(x)
	if err == nil {
		m.Exif.CameraSettings = c
	}

	// Two convenience functions exist for date/time taken and GPS coords:
	tm, _ := x.DateTime()
	fmt.Println("Taken: ", tm)

	return nil
}


func metadata(f *os.File) {
	var err error

	// Create Metadata
	m := NewMetadata(f)

	err = m.xmpMetadata(f)
	if err != nil {
		log.Println(err)
	}
	err = m.exifMetadata(f)
	if err != nil {
		log.Println(err)
	}

	a, _ := json.Marshal(m)
	colorJSON(a)

	//log.Println(m)
	//log.Println(tag, tag.Count, tag.Type)
}

func colorJSON(b []byte) (){

	var obj map[string]interface{}
    json.Unmarshal([]byte(b), &obj)
	// Make a custom formatter with indent set
    f := colorjson.NewFormatter()
    f.Indent = 4

    // Marshall the Colorized JSON
    s, _ := f.Marshal(obj)
    fmt.Println(string(s))
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