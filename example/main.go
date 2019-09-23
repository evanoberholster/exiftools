package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"encoding/json"

	"github.com/TylerBrock/colorjson"

	"github.com/evanoberholster/exif/exif"
	"github.com/evanoberholster/exif/mknote"
	"trimmer.io/go-xmp/xmp"
)

func main() {
	fname := "../../test/img/4.CR2" //.jpg"

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

// NewMetadata -
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
func (m *Metadata) xmpMetadata(f *os.File) error {
	f.Seek(0, 0)

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
func (m *Metadata) exifMetadata(f *os.File) error {
	m.Exif = Exif{}
	f.Seek(0, 0)

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
	m.Exif.FocalLength, _ = x.FocalLength()

	m.Exif.FocalLengthEqv, _ = FocalLengthIn35mmFilm(x)
	m.Exif.ExposureMode, _ = GetExposureMode(x)
	m.Exif.MeteringMode, _ = GetMeteringMode(x)

	m.Exif.Flash, _ = Flash(x)

	// Image Size
	m.Exif.ImageWidth, m.Exif.ImageHeight = GetImageSize(x)

	m.Exif.CameraMake, _ = FetchString(x, exif.Make)
	m.Exif.CameraModel, _ = FetchString(x, exif.Model)
	m.Exif.CameraSerial, _ = FetchString(x, exif.SerialNumber)
	m.Exif.Artist, _ = FetchString(x, exif.Artist)
	m.Exif.Copyright, _ = FetchString(x, exif.Copyright)
	m.Exif.Software, _ = FetchString(x, exif.Software)
	m.Exif.ImageDescription, _ = FetchString(x, exif.ImageDescription)
	m.Exif.Orientation, _ = Orientation(x)

	m.Exif.LensModel, _ = FetchString(x, exif.LensModel)
	m.Exif.LensSerial, _ = FetchString(x, exif.LensSerialNumber)

	m.Exif.ISOSpeed, _ = ISOSpeed(x)
	m.Exif.Aperture, _ = Aperture(x)
	m.Exif.ShutterSpeed, _ = GetShutterSpeed(x)
	m.Exif.ExposureBias, _ = ExposureBias(x)

	//m.Exif.CameraSettings, err = CanonCameraSettings(x)

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

// FocalLengthIn35mmFilm -
func FocalLengthIn35mmFilm(x *exif.Exif) (float32, error) {
	_, err := x.Get(exif.FocalLengthIn35mmFilm)
	if err != nil {
		return 0.0, err
	}

	//log.Println(tag.Type, tag.Count, tag)
	return 0.0, nil
}

// GetImageSize -
func GetImageSize(x *exif.Exif) (int, int) {
	w, err := x.Get(exif.ImageWidth)
	l, err := x.Get(exif.ImageLength)
	if err != nil {
		log.Println(err)
		return 0, 0
	}
	//log.Println(w.Type, l.Type, w.Count)
	width, _ := w.Int(0)
	length, _ := l.Int(0)
	return width, length
}
