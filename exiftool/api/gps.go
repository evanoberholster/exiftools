package api

import (
	"fmt"
	"time"

	"github.com/evanoberholster/exiftools/exiftool/exif"
	"github.com/evanoberholster/exiftools/exiftool/tags/ifd"
	"github.com/golang/geo/s2"
)

// GPS API constants
const (
	// GPSIfd "IFD/GPS"
	GPSIfdString = "IFD/GPS"
)

// gpsCoordsFromRationals returns a decimal given the EXIF-encoded information.
// The refValue is the N/E/S/W direction that this position is relative to.
func gpsCoordsFromRationals(refValue string, rawCoordinate []exif.Rational) (decimal float64, err error) {
	defer func() {
		if state := recover(); state != nil {
			err = state.(error)
		}
	}()

	if len(rawCoordinate) != 3 {
		err = fmt.Errorf("new GPS Coords requires a raw-coordinate with exactly three rationals")
		return
	}

	decimal = (float64(rawCoordinate[0].Numerator) / float64(rawCoordinate[0].Denominator))
	decimal += (float64(rawCoordinate[1].Numerator) / float64(rawCoordinate[1].Denominator) / 60.0)
	decimal += (float64(rawCoordinate[2].Numerator) / float64(rawCoordinate[2].Denominator) / 3600.0)

	// Decimal is a negative value for a South or West Orientation
	if refValue[0] == 'S' || refValue[0] == 'W' {
		decimal = -decimal
		return
	}
	return
}

// GpsInfo encapsulates all of the geographic information in one place.
type GpsInfo struct {
	Latitude, Longitude float64
	Altitude            int
	Timestamp           time.Time
}

// String returns a descriptive string.
func (gi *GpsInfo) String() string {
	return fmt.Sprintf("GpsInfo | LAT=(%.05f) LON=(%.05f) ALT=(%d) TIME=[%s] |",
		gi.Latitude, gi.Longitude, gi.Altitude, gi.Timestamp)
}

// S2CellID returns the cell-ID of the geographic location on the earth.
func (gi *GpsInfo) S2CellID() (cellID s2.CellID, err error) {
	defer func() {
		if state := recover(); state != nil {
			err = state.(error)
		}
	}()

	latLng := s2.LatLngFromDegrees(gi.Latitude, gi.Longitude)
	cellID = s2.CellIDFromLatLng(latLng)

	if !cellID.IsValid() {
		panic(ErrGpsCoordsNotValid)
	}

	return cellID, nil
}

// GPSInfo convenience func. "IFD/GPS" GPSLatitude and GPSLongitude
func (res ExifResults) GPSInfo() (lat, lng float64, err error) {
	defer func() {
		if state := recover(); state != nil {
			err = state.(error)
		}
	}()
	var ref string
	var raw []exif.Rational

	// Latitude
	ref, err = res.GetTag(GPSIfdString, 0, ifd.GPSLatitudeRef).GetString(res.exifReader)
	if err != nil {
		return
	}
	raw, err = res.GetTag(GPSIfdString, 0, ifd.GPSLatitude).GetRationals(res.exifReader)
	if err != nil {
		return
	}

	lat, err = gpsCoordsFromRationals(ref, raw)
	if err != nil {
		return
	}

	// Longitude
	ref, err = res.GetTag(GPSIfdString, 0, ifd.GPSLongitudeRef).GetString(res.exifReader)
	if err != nil {
		return
	}
	raw, err = res.GetTag(GPSIfdString, 0, ifd.GPSLongitude).GetRationals(res.exifReader)
	if err != nil {
		return
	}
	lng, err = gpsCoordsFromRationals(ref, raw)
	if err != nil {
		return
	}

	return
}

// GPSTime convenience func. "IFD/GPS" GPSDateStamp and GPSTimeStamp
func (res ExifResults) GPSTime() (timestamp time.Time, err error) {
	dateRaw, err := res.GetTag(GPSIfdString, 0, ifd.GPSDateStamp).GetString(res.exifReader)
	if err != nil {
		return
	}
	timeRaw, err := res.GetTag(GPSIfdString, 0, ifd.GPSTimeStamp).GetRationals(res.exifReader)
	if err != nil {
		return
	}
	hour := int(timeRaw[0].Numerator / timeRaw[0].Denominator)
	min := int(timeRaw[1].Numerator / timeRaw[1].Denominator)
	sec := int(timeRaw[2].Numerator / timeRaw[2].Denominator)

	timestamp, err = parseTimestamp(dateRaw, hour, min, sec)
	return
}
