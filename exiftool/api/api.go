package api

import (
	"errors"
	"fmt"
	"strings"

	"github.com/evanoberholster/exiftools/exiftool/tags/ifd"
	"github.com/evanoberholster/exiftools/exiftool/tags/ifdexif"
)

// API Errors
var (
	ErrEmptyTag = errors.New("Error empty tag")
	ErrParseTag = fmt.Errorf("Error parsing tag")
	ErrTagType  = errors.New("Error wrong tag Type")

	// ErrGpsCoordsNotValid means that some part of the geographic data were unparseable.
	ErrGpsCoordsNotValid = errors.New("GPS coordinates not valid")
)

// Artist convenience func. "IFD" Artist
func (res ExifResults) Artist() (artist string, err error) {
	return res.GetTag(ifd.IfdRoot, 0, ifd.Artist).GetString(res.exifReader)
}

// Copyright convenience func. "IFD" Copyright
func (res ExifResults) Copyright() (copyright string, err error) {
	return res.GetTag(ifd.IfdRoot, 0, ifd.Copyright).GetString(res.exifReader)
}

// CameraMake convenience func. "IFD" Make
func (res ExifResults) CameraMake() (make string, err error) {
	return res.GetTag(ifd.IfdRoot, 0, ifd.Make).GetString(res.exifReader)
}

// LensMake convenience func. "IFD/Exif" LensMake
func (res ExifResults) LensMake() (make string, err error) {
	return res.GetTag(ifdexif.FqIfdExif, 0, ifdexif.LensMake).GetString(res.exifReader)
}

// CameraModel convenience func. "IFD" Model
func (res ExifResults) CameraModel() (model string, err error) {
	return res.GetTag(ifd.IfdRoot, 0, ifd.Model).GetString(res.exifReader)
}

// LensModel convenience func. "IFD/Exif" LensModel
func (res ExifResults) LensModel() (model string, err error) {
	return res.GetTag(ifdexif.FqIfdExif, 0, ifdexif.LensModel).GetString(res.exifReader)
}

// CameraSerial convenience func. "IFD/Exif" BodySerialNumber
func (res ExifResults) CameraSerial() (serial string, err error) {
	// BodySerialNumber
	if serial, err = res.GetTag(ifdexif.FqIfdExif, 0, ifdexif.BodySerialNumber).GetString(res.exifReader); err == nil && serial != "" {
		return
	}

	// CameraSerialNumber
	return res.GetTag(ifd.IfdRoot, 0, ifd.CameraSerialNumber).GetString(res.exifReader)
}

// LensSerial convenience func. "IFD/Exif" LensSerialNumber
func (res ExifResults) LensSerial() (serial string, err error) {
	return res.GetTag(ifdexif.FqIfdExif, 0, ifdexif.LensSerialNumber).GetString(res.exifReader)
}

// Dimensions convenience func. "IFD" Dimensions
func (res ExifResults) Dimensions() (width, height int, err error) {
	width, err = res.GetTag(ifdexif.FqIfdExif, 0, ifdexif.PixelXDimension).GetInt(res.exifReader)
	if err == nil {
		height, err = res.GetTag(ifdexif.FqIfdExif, 0, ifdexif.PixelYDimension).GetInt(res.exifReader)
		if err == nil {
			return width, height, err
		}
	}

	width, err = res.GetTag(ifd.IfdRoot, 0, ifd.ImageWidth).GetInt(res.exifReader)
	if err == nil {
		height, err = res.GetTag(ifd.IfdRoot, 0, ifd.ImageLength).GetInt(res.exifReader)
		if err == nil {
			return width, height, err
		}
	}

	return 0, 0, ErrEmptyTag
}

// ExposureProgram convenience func. "IFD/Exif" ExposureProgram
func (res ExifResults) ExposureProgram() (ExposureMode, error) {
	ep, err := res.GetTag(ifdexif.FqIfdExif, 0, ifdexif.ExposureProgram).GetInt(res.exifReader)
	return ExposureMode(ep), err
}

// MeteringMode convenience func. "IFD/Exif" MeteringMode
func (res ExifResults) MeteringMode() (MeteringMode, error) {
	mm, err := res.GetTag(ifdexif.FqIfdExif, 0, ifdexif.MeteringMode).GetInt(res.exifReader)
	return MeteringMode(mm), err
}

// ShutterSpeed convenience func. "IFD/Exif" ExposureTime
func (res ExifResults) ShutterSpeed() (ShutterSpeed, error) {
	// ShutterSpeedValue
	// ExposureTime
	num, denom, err := res.GetTag(ifdexif.FqIfdExif, 0, ifdexif.ExposureTime).GetRational(res.exifReader)
	return ShutterSpeed{num, denom}, err
}

// Aperture convenience func. "IFD/Exif" FNumber
func (res ExifResults) Aperture() (float32, error) {
	// ApertureValue
	// FNumber
	if num, denom, err := res.GetTag(ifdexif.FqIfdExif, 0, ifdexif.FNumber).GetRational(res.exifReader); err == nil {
		if f := float32(num) / float32(denom); f > 0.0 {
			return f, nil
		}
	}
	return 0.0, ErrParseTag
}

// FocalLength convenience func. "IFD/Exif" FocalLength
// Lens Focal Length in mm
func (res ExifResults) FocalLength() (FocalLength, error) {
	// FocalLength
	if num, denom, err := res.GetTag(ifdexif.FqIfdExif, 0, ifdexif.FocalLength).GetRational(res.exifReader); err == nil {
		if f := float32(num) / float32(denom); f > 0.0 {
			return FocalLength(float32(num) / float32(denom)), nil
		}
	}
	return 0.0, ErrEmptyTag
}

// FocalLengthIn35mmFilm convenience func. "IFD/Exif" FocalLengthIn35mmFilm
// Lens Focal Length Equivalent for 35mm sensor in mm
func (res ExifResults) FocalLengthIn35mmFilm() (FocalLength, error) {
	// FocalLengthIn35mmFilm
	if fn, err := res.GetTag(ifdexif.FqIfdExif, 0, ifdexif.FocalLengthIn35mmFilm).GetInt(res.exifReader); err == nil {
		if fn > 0.0 {
			return FocalLength(fn), nil
		}
	}
	return 0.0, ErrEmptyTag
}

// ISOSpeed convenience func. "IFD/Exif" ISOSpeed
func (res ExifResults) ISOSpeed() (int, error) {
	iso, err := res.GetTag(ifdexif.FqIfdExif, 0, ifdexif.ISOSpeedRatings).GetInt(res.exifReader)
	return iso, err
}

// Flash convenience func. "IFD/Exif" Flash
func (res ExifResults) Flash() (FlashMode, error) {
	f, err := res.GetTag(ifdexif.FqIfdExif, 0, ifdexif.Flash).GetInt(res.exifReader)
	return FlashMode(f), err
}

// Orientation convenience func. "IFD" Orientation
// WIP
func (res ExifResults) Orientation() (string, error) {
	// Orientation
	return "", nil
}

// ExposureBiasValue convenience func. "IFD/Exif" ExposureBiasValue
// WIP
func (res ExifResults) ExposureBiasValue() (string, error) {
	return "", nil
}

// XMLPacket convenience func. "IFD" XMLPacket
// WIP
func (res ExifResults) XMLPacket() (str string, err error) {
	defer func() {
		if state := recover(); state != nil {
			err = state.(error)
		}
	}()
	str, err = res.GetTag(ifd.IfdRoot, 0, ifd.XMLPacket).GetString(res.exifReader)
	if err != nil {
		return
	}
	str = strings.Replace(str, "\n", "", -1)
	return strings.Replace(str, "   ", "", -1), nil
}

// Thumbnail convenience func. "IFD0" StripOffsets and StripByteCounts
// WIP
// Errors with NEF images
func (res ExifResults) Thumbnail() (int, int, error) {
	// CR2 file IFDO
	offset, err := res.GetTag(ifd.IfdRoot, 0, ifd.StripOffsets).GetInt(res.exifReader)
	if err != nil {
		return 0, 0, err
	}
	size, err := res.GetTag(ifd.IfdRoot, 0, ifd.StripByteCounts).GetInt(res.exifReader)
	if err != nil {
		return 0, 0, err
	}
	return offset, size, nil
}
