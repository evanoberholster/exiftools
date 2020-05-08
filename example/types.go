package main

import (
	"time"

	"github.com/evanoberholster/exiftools/models"
	"github.com/evanoberholster/exiftools/xmp"
)

// Metadata -
type Metadata struct {
	FileSize   int64          `json:"FileSize"`
	MIMEType   string         `json:"MIMEType"`
	DublinCore xmp.DublinCore `json:"DublinCore"`
	XmpBase    xmp.XmpBase    `json:"XmpBase"`
	Exif       Exif           `json:"Exif"`
	//CanonRaw   mknote.CanonRaw   `json:"CanonRaw"`
}

// Exif Metadata
type Exif struct {
	ImageWidth       uint                `json:"ImageWidth"`
	ImageHeight      uint                `json:"ImageHeight"`
	CameraMake       string              `json:"CameraMake"`       // OK
	CameraModel      string              `json:"CameraModel"`      // OK
	CameraSerial     string              `json:"CameraSerial"`     // OK
	LensModel        string              `json:"LensModel"`        // OK
	LensSerial       string              `json:"LensSerial"`       // OK
	Artist           string              `json:"Artist"`           // OK
	Copyright        string              `json:"Copyright"`        // OK
	Aperture         float32             `json:"Aperture"`         // OK
	ShutterSpeed     models.ShutterSpeed `json:"ShutterSpeed"`     // OK
	ISOSpeed         int                 `json:"ISO"`              // OK
	ExposureBias     models.ExposureBias `json:"ExposureBias"`     // OK
	ExposureMode     models.ExposureMode `json:"ExposureMode"`     // OK
	MeteringMode     models.MeteringMode `json:"MeteringMode"`     // OK
	Orientation      models.Orientation  `json:"Orientation"`      // Ok
	Flash            models.FlashMode    `json:"Flash"`            // OK
	FocalLength      float32             `json:"FocalLength"`      // OK
	FocalLengthEqv   float32             `json:"FocalLengthEqv"`   // mm
	GPSLatitude      float64             `json:"GPSLatitude"`      // OK
	GPSLongitude     float64             `json:"GPSLongitude"`     // OK
	GPSAltitude      float32             `json:"GPSAltitude"`      // OK
	GPSTimeStamp     time.Time           `json:"GPSTimeStamp"`     // OK
	DateTimeOriginal time.Time           `json:"DateTimeOriginal"` // OK
	ModifyTimeStamp  time.Time           `json:"ModifyTimeStamp"`  // Ok
	Software         string              `json:"Software"`         // OK
	ImageDescription string              `json:"ImageDescription"`
}
