package ifd

import (
	"github.com/evanoberholster/exiftools/exiftool/exif"
)

// IfdGPS Name and TagID
const (
	ifdGPS              = "GPS"
	ifdGPSID exif.TagID = 0x8825
)

// GPSPath is the GPS Ifd Path
var (
	GPSPath = IfdPath{IfdRootID}
)

// GPSIfd is the GPS IFD "IFD/GPS" for GPSInfo
var GPSIfd = IfdItem{GPSPath, ifdGPSID, ifdGPS}

// GPSIfdTags is a map of the the exif.TagID to exif.Tag contained in the GPSInfo ifd
var GPSIfdTags = map[exif.TagID]exif.Tag{
	GPSVersionID:        exif.NewTag("GPSVersionID", exif.TypeByte),
	GPSLatitudeRef:      exif.NewTag("GPSLatitudeRef", exif.TypeASCII),
	GPSLatitude:         exif.NewTag("GPSLatitude", exif.TypeRational),
	GPSLongitudeRef:     exif.NewTag("GPSLongitudeRef", exif.TypeASCII),
	GPSLongitude:        exif.NewTag("GPSLongitude", exif.TypeRational),
	GPSAltitudeRef:      exif.NewTag("GPSAltitudeRef", exif.TypeByte),
	GPSAltitude:         exif.NewTag("GPSAltitude", exif.TypeRational),
	GPSTimeStamp:        exif.NewTag("GPSTimeStamp", exif.TypeRational),
	GPSSatellites:       exif.NewTag("GPSSatellites", exif.TypeASCII),
	GPSStatus:           exif.NewTag("GPSStatus", exif.TypeASCII),
	GPSMeasureMode:      exif.NewTag("GPSMeasureMode", exif.TypeASCII),
	GPSDOP:              exif.NewTag("GPSDOP", exif.TypeRational),
	GPSSpeedRef:         exif.NewTag("GPSSpeedRef", exif.TypeASCII),
	GPSSpeed:            exif.NewTag("GPSSpeed", exif.TypeRational),
	GPSTrackRef:         exif.NewTag("GPSTrackRef", exif.TypeASCII),
	GPSTrack:            exif.NewTag("GPSTrack", exif.TypeRational),
	GPSImgDirectionRef:  exif.NewTag("GPSImgDirectionRef", exif.TypeASCII),
	GPSImgDirection:     exif.NewTag("GPSImgDirection", exif.TypeRational),
	GPSMapDatum:         exif.NewTag("GPSMapDatum", exif.TypeASCII),
	GPSDestLatitudeRef:  exif.NewTag("GPSDestLatitudeRef", exif.TypeASCII),
	GPSDestLatitude:     exif.NewTag("GPSDestLatitude", exif.TypeRational),
	GPSDestLongitudeRef: exif.NewTag("GPSDestLongitudeRef", exif.TypeASCII),
	GPSDestLongitude:    exif.NewTag("GPSDestLongitude", exif.TypeRational),
	GPSDestBearingRef:   exif.NewTag("GPSDestBearingRef", exif.TypeASCII),
	GPSDestBearing:      exif.NewTag("GPSDestBearing", exif.TypeRational),
	GPSDestDistanceRef:  exif.NewTag("GPSDestDistanceRef", exif.TypeASCII),
	GPSDestDistance:     exif.NewTag("GPSDestDistance", exif.TypeRational),
	GPSProcessingMethod: exif.NewTag("GPSProcessingMethod", exif.TypeUndefined),
	GPSAreaInformation:  exif.NewTag("GPSAreaInformation", exif.TypeUndefined),
	GPSDateStamp:        exif.NewTag("GPSDateStamp", exif.TypeASCII),
	GPSDifferential:     exif.NewTag("GPSDifferential", exif.TypeShort),
}

// GPSInfo Tags; GPSInfo Ifd
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
