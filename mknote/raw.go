package mknote

import (
	"github.com/evanoberholster/exiftools/exif"
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
	Lens            string `json:"Lens"`
}

// RawImage -
type RawImage interface {
	RawCameraSettings(x *exif.Exif) (CameraSettings, error)
}
