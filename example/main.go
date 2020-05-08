package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"encoding/json"

	"github.com/TylerBrock/colorjson"

	"github.com/evanoberholster/exiftools/exif"
	"github.com/evanoberholster/exiftools/mknote"
	"github.com/evanoberholster/exiftools/xmp"
	"github.com/evanoberholster/filetype"
)

func main() {
	fname := "../../test/img/a.jpg" //.jpg"

	f, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	start := time.Now()
	metadata(f)
	log.Println(time.Since(start))
}

// NewMetadata -
func NewMetadata(f *os.File) *Metadata {
	f.Seek(0, 0)
	// Check fileSize
	fi, err := f.Stat()
	if err != nil {
		log.Println(err)
	}

	kind, _ := filetype.MatchReader(f)

	// Check file MimeType
	return &Metadata{
		FileSize: fi.Size(),
		MIMEType: kind.MIME.Value,
	}

}

// XMP
func (m *Metadata) xmpMetadata(f *os.File) error {
	start := time.Now()
	doc, err := xmp.ReadXMPDocument(f)
	if err != nil {
		return err
	}

	m.DublinCore = xmp.GetDublinCore(doc)
	m.XmpBase = xmp.GetBase(doc)

	fmt.Println(time.Since(start), m.DublinCore, m.XmpBase)

	return nil
}

// EXIF
func (m *Metadata) exifMetadata(f *os.File) error {
	m.Exif = Exif{}
	f.Seek(0, 0)

	// Optionally register camera makernote data parsing - Canon, Nikon and AdobeDNG are supported
	exif.RegisterParsers(mknote.All...)

	x, err := exif.DecodeWithParseHeader(f)
	if err != nil {
		return err
	}
	//log.Println(x)

	m.Exif.GPSAltitude, _ = x.GPSAltitude()
	m.Exif.GPSTimeStamp, _ = x.GPSTimeStamp()

	m.Exif.GPSLatitude, m.Exif.GPSLongitude, _ = x.LatLong()

	m.Exif.DateTimeOriginal, _ = x.DateTime()

	m.Exif.FocalLength, _ = x.FocalLength(exif.FocalLength)

	m.Exif.FocalLengthEqv, _ = x.FocalLength(exif.FocalLengthIn35mmFilm)
	m.Exif.ExposureMode, _ = x.GetExposureMode()
	m.Exif.MeteringMode, _ = x.GetMeteringMode()
	m.Exif.Flash, _ = x.GetFlashMode()

	// Image Size
	m.Exif.ImageHeight, _ = x.GetUints(exif.ImageLength, exif.PixelYDimension)
	//if err != nil {
	//	fmt.Println("Error", err)
	//}
	m.Exif.ImageWidth, _ = x.GetUints(exif.ImageWidth, exif.PixelXDimension)
	//if err != nil {
	//	fmt.Println("Error", err)
	//}
	//m.Exif.ImageWidth, m.Exif.ImageHeight = x.GetImageSize()
	m.Exif.CameraMake, _ = x.GetString(exif.Make)
	m.Exif.CameraModel, _ = x.GetString(exif.Model)
	m.Exif.CameraSerial, _ = x.GetString(exif.SerialNumber)
	m.Exif.Artist, _ = x.GetString(exif.Artist)
	m.Exif.Copyright, _ = x.GetString(exif.Copyright)
	m.Exif.Software, _ = x.GetString(exif.Software)
	m.Exif.ImageDescription, _ = x.GetString(exif.ImageDescription)
	m.Exif.Orientation, _ = x.GetOrientation()
	m.Exif.LensModel, _ = x.GetString(exif.LensModel)
	m.Exif.LensSerial, _ = x.GetString(exif.LensSerialNumber)

	m.Exif.ISOSpeed, _ = x.GetISOSpeed()
	m.Exif.Aperture, _ = x.GetAperture()
	m.Exif.ShutterSpeed, _ = x.GetShutterSpeed()
	m.Exif.ExposureBias, _ = x.GetExposureBias()

	//colorJSON(a)
	//fmt.Println(x.DateTime())
	cr := new(mknote.CanonRaw)
	cr.Get(x)
	fmt.Println(cr)
	//m.CanonRaw = *cr
	//fmt.Println(x.Get(mknote.CanonFileInfo))

	a, _ := json.Marshal(m)
	ioutil.WriteFile("m.json", []byte(a), 0644)

	return nil
}

func metadata(f *os.File) {
	var err error

	// Create New Metadata
	m := NewMetadata(f)

	if err = m.xmpMetadata(f); err != nil {
		log.Println(err)
	}

	if err = m.exifMetadata(f); err != nil {
		log.Println(err)
	}

	//fmt.Println(m)
}

func colorJSON(b []byte) {
	var obj map[string]interface{}
	json.Unmarshal([]byte(b), &obj)
	// Make a custom formatter with indent set
	f := colorjson.NewFormatter()
	f.Indent = 4

	// Marshall the Colorized JSON
	s, _ := f.Marshal(obj)
	fmt.Println(string(s))
}
