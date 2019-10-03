package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"encoding/json"

	"github.com/TylerBrock/colorjson"

	"github.com/evanoberholster/exif/exif"
	"github.com/evanoberholster/exif/mknote"
	"github.com/evanoberholster/exif/xmp"
	"github.com/evanoberholster/filetype"
)

func main() {
	fname := "../../test/img/2.CR2" //.jpg"

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
	doc, err := xmp.ReadXMPDocument(f)
	if err != nil {
		return err
	}

	m.DublinCore = xmp.DublinCore(doc)
	m.XmpBase = xmp.Base(doc)

	return nil
}

// EXIF
func (m *Metadata) exifMetadata(f *os.File) error {
	m.Exif = Exif{}
	f.Seek(0, 0)

	// Optionally register camera makenote data parsing - currently Nikon and
	// Canon are supported.
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

	// ModifyTimeStamp

	fmt.Println(x.PreviewImage())
	fmt.Println(x.JpegThumbnail())

	cr := new(mknote.CanonRaw)
	m.Exif.CameraSettings, _ = cr.RawCameraSettings(x)

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
