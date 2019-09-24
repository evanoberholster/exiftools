package models

// Canon specific makernote values
var (
	CanonContinuousDriveValues = map[int]string{
		0:  "Single",
		1:  "Continuous",
		2:  "Movie",
		3:  "Continuous, Speed Priority",
		4:  "Continuous, Low",
		5:  "Continuous, High",
		6:  "Silent Single",
		9:  "Single, Silent",
		10: "Continuous, Silent",
	}
	CanonFocusModeValues = map[int]string{
		0:   "One-shot AF",
		1:   "AI Servo AF",
		2:   "AI Focus AF",
		3:   "Manual Focus (3)",
		4:   "Single",
		5:   "Continuous",
		6:   "Manual Focus (6)",
		16:  "Pan Focus",
		256: "AF + MF",
		512: "Movie Snap Focus",
		519: "Movie Servo AF",
	}
	CanonExposureModeValues = map[int]string{
		0: "Easy",
		1: "Program AE",
		2: "Shutter speed priority AE",
		3: "Aperture-priority AE",
		4: "Manual",
		5: "Depth-of-field AE",
		6: "M-Dep",
		7: "Bulb",
		8: "Flexible-priority AE",
	}
	CanonRecordModeValues = map[int]string{
		1:  "JPEG",
		2:  "CRW+THM",
		3:  "AVI+THM",
		4:  "TIF",
		5:  "TIF+JPEG",
		6:  "CR2",
		7:  "CR2+JPEG",
		9:  "MOV",
		10: "MP4",
		11: "CRM",
		12: "CR3",
		13: "CR3+JPEG",
	}
	CanonMeteringModeValues = map[int]string{
		0: "Default",
		1: "Spot",
		2: "Average",
		3: "Evaluative",
		4: "Partial",
		5: "Center-weighted average",
	}
	CanonAESettingValues = map[int]string{
		0: "Normal AE",
		1: "Exposure Compensation",
		2: "AE Lock",
		3: "AE Lock + Exposure Compensation",
		4: "No AE",
	}
)
