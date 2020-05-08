package api

import (
	"github.com/evanoberholster/exiftools/exiftool/tags/ifd"
)

func (res Results) Artist() (artist string, err error) {
	return res.GetTag("IFD", ifd.Artist).String()
}

func (res Results) Copyright() (copyright string, err error) {
	return res.GetTag("IFD", ifd.Copyright).String()
}

func (res Results) Make() (make string, err error) {
	return res.GetTag("IFD", ifd.Make).String()
}

func (res Results) Model() (model string, err error) {
	return res.GetTag("IFD", ifd.Model).String()
}

func (res Results) Dimensions() (width, height int, err error) {
	width, err = res.GetTag("IFD", ifd.ImageWidth).Int()
	if err != nil {
		return 0, 0, err
	}
	height, err = res.GetTag("IFD", ifd.ImageLength).Int()
	if err != nil {
		return 0, 0, err
	}
	return
}

// ExposureMode
// Flash
// MeteringMode
// ExposureBiasValue
// ApertureValue
// ShutterSpeedValue
// FNumber
// ExposureTime
// Orientation
