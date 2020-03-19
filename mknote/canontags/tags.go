package canontags

import (
	"errors"

	"github.com/evanoberholster/exiftools/models"
)

//go:generate go run genCanon.go

const (
	UnknownCameraModel = models.CameraModel("Unknown")
)

var (
	UnknownCameraLensType = CanonLensType([]string{"Unknown"})
	ErrModelNotFound      = errors.New("Model Not Found")
	ErrLensTypeNotFound   = errors.New("Lens type Not Found")
)

// CanonLensType -
type CanonLensType []string

// CanonModel -
func CanonModel(id uint32) (models.CameraModel, error) {
	m, ok := canonModelIDValues[id]
	if !ok {
		return UnknownCameraModel, ErrModelNotFound
	}
	return m, nil
}

// CanonLens -
func CanonLens(lensType int) string {
	m, ok := canonLensTypeValues[lensType]
	if !ok {
		return ErrLensTypeNotFound.Error()
	}
	return m[0]
}

// CanonImageSize -
var canonImageSize = map[int]string{
	-1:  "n/a",
	0:   "Large",
	1:   "Medium",
	2:   "Small",
	5:   "Medium 1",
	6:   "Medium 2",
	7:   "Medium 3",
	8:   "Postcard",
	9:   "Widescreen",
	10:  "Medium Widescreen",
	14:  "Small 1",
	15:  "Small 2",
	16:  "Small 3",
	128: "640x480 Movie",
	129: "Medium Movie",
	130: "Small Movie",
	137: "1280x720 Movie",
	142: "1920x1080 Movie",
	143: "4096x2160 Movie",
}

// %canonQuality
// %canonWhiteBalance
// %pictureStyles
// %userDefStyles
