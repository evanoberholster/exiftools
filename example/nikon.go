package main

import (
	"github.com/evanoberholster/exif/exif"
	"github.com/evanoberholster/exif/mknote"
)

// RawImage -
type RawImage interface {
	RawCameraSettings(x *exif.Exif) (CameraSettings, error)
}

// CanonRaw - Raw Image from a Canon Camera
type CanonRaw string

// RawCameraSettings - Get Canon camera Settings
func (cr *CanonRaw) RawCameraSettings(x *exif.Exif) (CameraSettings, error) {
	c := CameraSettings{}
	tag, err := x.Get(mknote.Canon_CameraSettings)
	if err != nil {
		return c, err
	}

	c.ContinuousDrive = ProcessCameraSettingsFields(tag, CanonContinuousDrive)
	c.RecordMode = ProcessCameraSettingsFields(tag, CanonRecordMode)
	c.FocusMode = ProcessCameraSettingsFields(tag, CanonFocusMode)
	c.ExposureMode = ProcessCameraSettingsFields(tag, CanonExposureMode)
	c.MeteringMode = ProcessCameraSettingsFields(tag, CanonMeteringMode)

	return c, nil
}
