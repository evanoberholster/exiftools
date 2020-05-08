package ifd

import (
	"github.com/evanoberholster/exiftools/exiftool"
	"github.com/evanoberholster/exiftools/exiftool/exif"
)

// IfdGPS Name and TagID
const (
	IfdGPS              = "GPS"
	IfdGPSID exif.TagID = 0x8825
)

// GPSPath is the GPS Ifd Path
var (
	GPSPath = exif.IfdPath{IfdRootID}
)

// GPSIfd is the GPS IFD "IFD/GPS" for GPSInfo
var GPSIfd = exiftool.IfdItem{GPSPath, IfdGPSID, IfdGPS}

// GPSIfdTags is a map of the the exif.TagID to exiftool.IndexedTag
var GPSIfdTags = map[exif.TagID]exiftool.IndexedTag{
	GPSVersionID: {GPSVersionID, "GPSVersionID", exif.TypeShort}, // double Check
}

var GPSIfdTags2 = map[exif.TagID]exif.Tag{
	GPSVersionID:        {"GPSVersionID", exif.TypeByte},
	GPSLatitudeRef:      {"GPSLatitudeRef", exif.TypeASCII},
	GPSLatitude:         {"GPSLatitude", exif.TypeRational},
	GPSLongitudeRef:     {"GPSLongitudeRef", exif.TypeASCII},
	GPSLongitude:        {"GPSLongitude", exif.TypeRational},
	GPSAltitudeRef:      {"GPSAltitudeRef", exif.TypeByte},
	GPSAltitude:         {"GPSAltitude", exif.TypeRational},
	GPSTimeStamp:        {"GPSTimeStamp", exif.TypeRational},
	GPSSatellites:       {"GPSSatellites", exif.TypeASCII},
	GPSStatus:           {"GPSStatus", exif.TypeASCII},
	GPSMeasureMode:      {"GPSMeasureMode", exif.TypeASCII},
	GPSDOP:              {"GPSDOP", exif.TypeRational},
	GPSSpeedRef:         {"GPSSpeedRef", exif.TypeASCII},
	GPSSpeed:            {"GPSSpeed", exif.TypeRational},
	GPSTrackRef:         {"GPSTrackRef", exif.TypeASCII},
	GPSTrack:            {"GPSTrack", exif.TypeRational},
	GPSImgDirectionRef:  {"GPSImgDirectionRef", exif.TypeASCII},
	GPSImgDirection:     {"GPSImgDirection", exif.TypeRational},
	GPSMapDatum:         {"GPSMapDatum", exif.TypeASCII},
	GPSDestLatitudeRef:  {"GPSDestLatitudeRef", exif.TypeASCII},
	GPSDestLatitude:     {"GPSDestLatitude", exif.TypeRational},
	GPSDestLongitudeRef: {"GPSDestLongitudeRef", exif.TypeASCII},
	GPSDestLongitude:    {"GPSDestLongitude", exif.TypeRational},
	GPSDestBearingRef:   {"GPSDestBearingRef", exif.TypeASCII},
	GPSDestBearing:      {"GPSDestBearing", exif.TypeRational},
	GPSDestDistanceRef:  {"GPSDestDistanceRef", exif.TypeASCII},
	GPSDestDistance:     {"GPSDestDistance", exif.TypeRational},
	GPSProcessingMethod: {"GPSProcessingMethod", exif.TypeUndefined},
	GPSAreaInformation:  {"GPSAreaInformation", exif.TypeUndefined},
	GPSDateStamp:        {"GPSDateStamp", exif.TypeASCII},
	GPSDifferential:     {"GPSDifferential", exif.TypeShort},
}

// GPSInfo Tags
const (
	GPSVersionID        exif.TagID = 0x0000
	GPSLatitudeRef      exif.TagID = 0x0001
	GPSLatitude         exif.TagID = 0x0002
	GPSLongitudeRef     exif.TagID = 0x0003
	GPSLongitude        exif.TagID = 0x0004
	GPSAltitudeRef      exif.TagID = 0x0005
	GPSAltitude         exif.TagID = 0x0006
	GPSTimeStamp        exif.TagID = 0x0007
	GPSSatellites       exif.TagID = 0x0008
	GPSStatus           exif.TagID = 0x0009
	GPSMeasureMode      exif.TagID = 0x000a
	GPSDOP              exif.TagID = 0x000b
	GPSSpeedRef         exif.TagID = 0x000c
	GPSSpeed            exif.TagID = 0x000d
	GPSTrackRef         exif.TagID = 0x000e
	GPSTrack            exif.TagID = 0x000f
	GPSImgDirectionRef  exif.TagID = 0x0010
	GPSImgDirection     exif.TagID = 0x0011
	GPSMapDatum         exif.TagID = 0x0012
	GPSDestLatitudeRef  exif.TagID = 0x0013
	GPSDestLatitude     exif.TagID = 0x0014
	GPSDestLongitudeRef exif.TagID = 0x0015
	GPSDestLongitude    exif.TagID = 0x0016
	GPSDestBearingRef   exif.TagID = 0x0017
	GPSDestBearing      exif.TagID = 0x0018
	GPSDestDistanceRef  exif.TagID = 0x0019
	GPSDestDistance     exif.TagID = 0x001a
	GPSProcessingMethod exif.TagID = 0x001b
	GPSAreaInformation  exif.TagID = 0x001c
	GPSDateStamp        exif.TagID = 0x001d
	GPSDifferential     exif.TagID = 0x001e
)

// GPSInfoIfdTagNames -
var GPSInfoIfdTagNames = map[exif.TagID]string{
	GPSVersionID:        "GPSVersionID",
	GPSLatitudeRef:      "GPSLatitudeRef",
	GPSLatitude:         "GPSLatitude",
	GPSLongitudeRef:     "GPSLongitudeRef",
	GPSLongitude:        "GPSLongitude",
	GPSAltitudeRef:      "GPSAltitudeRef",
	GPSAltitude:         "GPSAltitude",
	GPSTimeStamp:        "GPSTimeStamp",
	GPSSatellites:       "GPSSatellites",
	GPSStatus:           "GPSStatus",
	GPSMeasureMode:      "GPSMeasureMode",
	GPSDOP:              "GPSDOP",
	GPSSpeedRef:         "GPSSpeedRef",
	GPSSpeed:            "GPSSpeed",
	GPSTrackRef:         "GPSTrackRef",
	GPSTrack:            "GPSTrack",
	GPSImgDirectionRef:  "GPSImgDirectionRef",
	GPSImgDirection:     "GPSImgDirection",
	GPSMapDatum:         "GPSMapDatum",
	GPSDestLatitudeRef:  "GPSDestLatitudeRef",
	GPSDestLatitude:     "GPSDestLatitude",
	GPSDestLongitudeRef: "GPSDestLongitudeRef",
	GPSDestLongitude:    "GPSDestLongitude",
	GPSDestBearingRef:   "GPSDestBearingRef",
	GPSDestBearing:      "GPSDestBearing",
	GPSDestDistanceRef:  "GPSDestDistanceRef",
	GPSDestDistance:     "GPSDestDistance",
	GPSProcessingMethod: "GPSProcessingMethod",
	GPSAreaInformation:  "GPSAreaInformation",
	GPSDateStamp:        "GPSDateStamp",
	GPSDifferential:     "GPSDifferential",
}
