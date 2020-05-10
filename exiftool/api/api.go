package api

import (
	"github.com/evanoberholster/exiftools/exiftool/tags/ifd"
	"github.com/evanoberholster/exiftools/exiftool/tags/ifdexif"
)

// Artist convenience func. "IFD" Artist
func (res Results) Artist() (artist string, err error) {
	return res.GetTag("IFD", ifd.Artist).String()
}

// Copyright convenience func. "IFD" Copyright
func (res Results) Copyright() (copyright string, err error) {
	return res.GetTag("IFD", ifd.Copyright).String()
}

// CameraMake convenience func. "IFD" Make
func (res Results) CameraMake() (make string, err error) {
	return res.GetTag("IFD", ifd.Make).String()
}

// LensMake convenience func. "IFD/Exif" LensMake
func (res Results) LensMake() (make string, err error) {
	return res.GetTag("IFD/Exif", ifdexif.LensMake).String()
}

// CameraModel convenience func. "IFD" Model
func (res Results) CameraModel() (model string, err error) {
	return res.GetTag("IFD", ifd.Model).String()
}

// LensModel convenience func. "IFD/Exif" LensModel
func (res Results) LensModel() (model string, err error) {
	return res.GetTag("IFD/Exif", ifdexif.LensModel).String()
}

// CameraSerial convenience func. "IFD/Exif" BodySerialNumber
func (res Results) CameraSerial() (serial string, err error) {
	// BodySerialNumber
	if serial, err = res.GetTag("IFD/Exif", ifdexif.BodySerialNumber).String(); err == nil && serial != "" {
		return
	}

	// CameraSerialNumber
	return res.GetTag("IFD", ifd.CameraSerialNumber).String()
}

// LensSerial convenience func. "IFD/Exif" LensSerialNumber
func (res Results) LensSerial() (serial string, err error) {
	return res.GetTag("IFD/Exif", ifdexif.LensSerialNumber).String()
}

// Dimensions convenience func. "IFD" Dimensions
func (res Results) Dimensions() (width, height int, err error) {
	width, err = res.GetTag("IFD/Exif", ifdexif.PixelXDimension).Int()
	if err == nil {
		height, err = res.GetTag("IFD/Exif", ifdexif.PixelYDimension).Int()
		if err == nil {
			return width, height, err
		}
	}

	width, err = res.GetTag("IFD", ifd.ImageWidth).Int()
	if err == nil {
		height, err = res.GetTag("IFD", ifd.ImageLength).Int()
		if err == nil {
			return width, height, err
		}
	}

	return 0, 0, ErrEmptyTag
}

// ExposureProgram convenience func. "IFD/Exif" ExposureProgram
func (res Results) ExposureProgram() (ExposureMode, error) {
	ep, err := res.GetTag("IFD/Exif", ifdexif.ExposureProgram).Int()
	return ExposureMode(ep), err
}

// MeteringMode convenience func. "IFD/Exif" MeteringMode
func (res Results) MeteringMode() (MeteringMode, error) {
	mm, err := res.GetTag("IFD/Exif", ifdexif.MeteringMode).Int()
	return MeteringMode(mm), err
}

// ExposureBiasValue convenience func. "IFD/Exif" ExposureBiasValue
// WIP
func (res Results) ExposureBiasValue() (string, error) {
	return "", nil
}

// ShutterSpeed convenience func. "IFD/Exif" ExposureTime
func (res Results) ShutterSpeed() (ShutterSpeed, error) {
	// ShutterSpeedValue
	// ExposureTime
	et, err := res.GetTag("IFD/Exif", ifdexif.ExposureTime).Rational()
	return ShutterSpeed{int(et[0].Numerator), int(et[0].Denominator)}, err
}

// Aperture convenience func. "IFD/Exif" FNumber
func (res Results) Aperture() (float32, error) {
	// ApertureValue
	// FNumber
	fn, err := res.GetTag("IFD/Exif", ifdexif.FNumber).Rational()
	return float32(fn[0].Numerator) / float32(fn[0].Denominator), err
}

// FocalLength convenience func. "IFD/Exif" FocalLength
// Lens Focal Length in mm
func (res Results) FocalLength() (FocalLength, error) {
	// FocalLength
	fn, err := res.GetTag("IFD/Exif", ifdexif.FocalLength).Rational()
	return FocalLength(float32(fn[0].Numerator) / float32(fn[0].Denominator)), err
}

// FocalLengthIn35mmFilm convenience func. "IFD/Exif" FocalLengthIn35mmFilm
// Lens Focal Length Equivalent for 35mm sensor in mm
func (res Results) FocalLengthIn35mmFilm() (FocalLength, error) {
	// FocalLengthIn35mmFilm
	fn, err := res.GetTag("IFD/Exif", ifdexif.FocalLengthIn35mmFilm).Int()
	return FocalLength(fn), err
}

// ISOSpeed convenience func. "IFD/Exif" ISOSpeed
func (res Results) ISOSpeed() (int, error) {
	iso, err := res.GetTag("IFD/Exif", ifdexif.ISOSpeedRatings).Int()
	return iso, err
}

// Flash convenience func. "IFD/Exif" Flash
func (res Results) Flash() (FlashMode, error) {
	f, err := res.GetTag("IFD/Exif", ifdexif.Flash).Int()
	return FlashMode(f), err
}

// Orientation convenience func. "IFD" Orientation
// WIP
func (res Results) Orientation() (string, error) {
	// Orientation
	return "", nil
}
