package canontags

import (
	"errors"
	"fmt"

	"github.com/evanoberholster/exiftools/models"
	"github.com/evanoberholster/exiftools/tiff"
)

//go:generate go run genCanon.go

const (
	UnknownCameraModel = models.CameraModel("Unknown")
)

var (
	UnknownCameraLensType = CanonLensType([]string{"Unknown"})
	ErrModelNotFound      = errors.New("Model Not Found")
	ErrLensTypeNotFound   = errors.New("Lens type Not Found")
	ErrMakerNote          = errors.New("Error reading MakerNote")
)

// CanonAFInfo - Work In Progress
type CanonAFInfo struct {
	AFAreaMode       string // 1: AFAreaMode
	NumAFPoints      int    // 2:
	ValidAFPoints    int    // 3:
	CanonImageWidth  int    // 4:
	CanonImageHeight int    // 5:
	//AFImageWidth     int    // 6:
	//AFImageHeight    int    // 7:
	//AFAreaWidths     int    // 8:
	//AFAreaHeights    int    // 9:
	AFAreaPositions map[AFPosition]int
	//AFPointsInFocus  int
	//AFPointsSelected int
}

// Get CanonAFInfo
// WIP - March 20, 2020
func (cai *CanonAFInfo) Get(tag *tiff.Tag) error {
	if i, err := tag.Int(1); err == nil {
		cai.AFAreaMode = afAreaModeValues[i]
	}

	cai.NumAFPoints, _ = tag.Int(2)
	cai.ValidAFPoints, _ = tag.Int(3)
	cai.CanonImageHeight, _ = tag.Int(4)
	cai.CanonImageWidth, _ = tag.Int(5)

	cai.AFAreaPositions = make(map[AFPosition]int)
	for i := 8; i < cai.ValidAFPoints+8; i++ {
		posX, _ := tag.Int(i)
		posY, _ := tag.Int(i + cai.ValidAFPoints)
		cai.AFAreaPositions[AFPosition{posX, posY}] = 1
	}
	return nil
}

// AFPosition - Position of AutoFocus Points (x,y)
type AFPosition [2]int

// ToInt - Convert AFPosition to Int
func (ap *AFPosition) ToInt() (int, int) {
	return ap[0], ap[1] // x,y
}

// FocusDistance -
type FocusDistance [2]float32

func (fd *FocusDistance) String() string {
	return fmt.Sprintf("%f.2 - %f.2", fd[0], fd[1])
}

// CanonShotInfo - WIP
type CanonShotInfo struct {
	ExposureCompensation int
	CameraTemperature    int
	FlashGuideNumber     int
	SequenceNumber       int
	//FlashExposureComp    int
	//AutoExposureBra      int
	//AEBBracketVal        int
	//FocusDistance  FocusDistance
}

// Get the CanonShotInfo from a *tiff.Tag
func (csi *CanonShotInfo) Get(tag *tiff.Tag) error {
	// 6: ExposureCompensation
	if e, err := tag.Int(6); err == nil {
		csi.ExposureCompensation = e
	}
	// 9: SequenceNumber
	if e, err := tag.Int(9); err == nil {
		csi.SequenceNumber = e
	}
	// 12: CameraTemperature
	if e, err := tag.Int(12); err == nil {
		csi.CameraTemperature = e - 128 // Conversion
	}
	// 13: FlashGuideNumber
	if e, err := tag.Int(13); err == nil {
		csi.FlashGuideNumber = e / 32 // Conversion
	}
	// FocusDistance
	//if fdUpper, err := tag.Int(19); err == nil { // 19: FocusDistanceUpper
	//	if fdLower, err := tag.Int(20); err == nil { // 20: FocusDistanceLower
	//		csi.FocusDistance = FocusDistance{float32(fdUpper), float32(fdLower)}
	//	}
	//}
	return nil
}

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
var canonImageSizeValues = map[int]string{
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

// AFAreaMode -
// Updated on March 19,2020
var afAreaModeValues = map[int]string{
	0:  "Off (Manual Focus)",
	1:  "AF Point Expansion (surround)",
	2:  "Single-point AF",
	3:  "n/a",
	4:  "Auto",
	5:  "Face Detect AF",
	6:  "Face + Tracking",
	7:  "Zone AF",
	8:  "AF Point Expansion (4 point)",
	9:  "Spot AF",
	10: "AF Point Expansion (8 point)",
	11: "Flexizone Multi (49 point)",
	12: "Flexizone Multi (9 point)",
	13: "Flexizone Single",
	14: "Large Zone AF",
}
