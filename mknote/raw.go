package mknote

import (
	"github.com/evanoberholster/exiftools/exif"
	"github.com/evanoberholster/exiftools/tiff"
)

// CameraSettingsField -
type CameraSettingsField map[int]string

// CameraSettings -
type CameraSettings struct {
	RecordMode      string `json:"RecordMode"`
	FocusMode       string `json:"FocusMode"`
	ExposureMode    string `json:"ExposureMode"`
	ContinuousDrive string `json:"ContinuousDrive"`
	MeteringMode    string `json:"MeteringMode"`
}

func processCameraSettingsFields(tag *tiff.Tag, i int) string {
	a, err := tag.Int(i)
	if err != nil {
		return ""
	}
	return CanonCameraSettingsFields[i][a]
}

// RawImage -
type RawImage interface {
	RawCameraSettings(x *exif.Exif) (CameraSettings, error)
}
