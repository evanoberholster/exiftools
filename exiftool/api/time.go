package api

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/evanoberholster/exiftools/exiftool/tags/ifd"
	"github.com/evanoberholster/exiftools/exiftool/tags/ifdexif"
)

// Errors for Parsing of Time
var (
	ErrParseYear  = fmt.Errorf("Error parsing Year")
	ErrParseMonth = fmt.Errorf("Error parsing Month")
	ErrParseDay   = fmt.Errorf("Error parsing Day")
	ErrParseHour  = fmt.Errorf("Error parsing Hour")
	ErrParseMin   = fmt.Errorf("Error parsing Min")
)

// ModifyDate - the date and time at which the Exif file was modified
func (res ExifResults) ModifyDate() (time.Time, error) {
	// "IFD" DateTime
	if dateRaw, err := res.GetTag("IFD", 0, ifd.DateTime).GetString(res.exifReader); err == nil && dateRaw != "" {
		// "IFD/Exif" SubSecTime
		subSecRaw, _ := res.GetTag("IFD/Exif", 0, ifdexif.SubSecTime).GetString(res.exifReader)
		if dateTime, err := parseExifFullTimestamp(dateRaw, subSecRaw); err == nil && !dateTime.IsZero() {
			return dateTime, nil
		}
	}
	return time.Time{}, ErrEmptyTag
}

// DateTime - the date and time at which the EXIF file was created
// with sub-second precision
func (res ExifResults) DateTime() (time.Time, error) {
	// "IFD/Exif" DateTimeOriginal
	if dateRaw, err := res.GetTag("IFD/Exif", 0, ifdexif.DateTimeOriginal).GetString(res.exifReader); err == nil && dateRaw != "" {
		// "IFD/Exif" SubSecTimeOriginal
		subSecRaw, _ := res.GetTag("IFD/Exif", 0, ifdexif.SubSecTimeOriginal).GetString(res.exifReader)

		if dateTime, err := parseExifFullTimestamp(dateRaw, subSecRaw); err == nil && !dateTime.IsZero() {
			return dateTime, nil
		}
	}

	// "IFD/Exif" DateTimeDigitized
	if dateRaw, err := res.GetTag("IFD/Exif", 0, ifdexif.DateTimeDigitized).GetString(res.exifReader); err == nil && dateRaw != "" {
		// "IFD/Exif" SubSecTimeDigitized
		subSecRaw, _ := res.GetTag("IFD/Exif", 0, ifdexif.SubSecTimeDigitized).GetString(res.exifReader)
		if dateTime, err := parseExifFullTimestamp(dateRaw, subSecRaw); err == nil && !dateTime.IsZero() {
			return dateTime, nil
		}
	}
	return time.Time{}, ErrEmptyTag
}

// parseExifFullTimestamp parses dates like "2018:11:30 13:01:49" into a UTC
// `time.Time` struct.
func parseExifFullTimestamp(fullTimestampPhrase string, subSecString string) (timestamp time.Time, err error) {
	defer func() {
		if state := recover(); state != nil {
			err = state.(error)
		}
	}()

	parts := strings.Split(fullTimestampPhrase, " ")
	datestampValue, timestampValue := parts[0], parts[1]

	dateParts := strings.Split(datestampValue, ":")

	year, err := strconv.ParseUint(dateParts[0], 10, 16)
	if err != nil {
		err = ErrParseYear
		return
	}

	month, err := strconv.ParseUint(dateParts[1], 10, 8)
	if err != nil {
		err = ErrParseMonth
		return
	}

	day, err := strconv.ParseUint(dateParts[2], 10, 8)
	if err != nil {
		err = ErrParseDay
		return
	}

	timeParts := strings.Split(timestampValue, ":")

	hour, err := strconv.ParseUint(timeParts[0], 10, 8)
	if err != nil {
		err = ErrParseHour
		return
	}

	minute, err := strconv.ParseUint(timeParts[1], 10, 8)
	if err != nil {
		err = ErrParseMin
		return
	}

	second, err := strconv.ParseUint(timeParts[2], 10, 8)
	if err != nil {
		err = ErrParseMin
		return
	}

	subSec, err := strconv.ParseUint(subSecString, 10, 16)
	if err != nil {
		subSec = 0
	}

	timestamp = time.Date(int(year), time.Month(month), int(day), int(hour), int(minute), int(second), int(subSec*1000000), time.UTC)
	return timestamp, nil
}

// ParseTimestamp parses dates like "2018:11:30" into a UTC `time.Time` struct.
func parseTimestamp(dateStamp string, hour, min, sec int) (timestamp time.Time, err error) {
	defer func() {
		if state := recover(); state != nil {
			err = state.(error)
		}
	}()

	dateParts := strings.Split(dateStamp, ":")

	year, err := strconv.ParseUint(dateParts[0], 10, 16)
	if err != nil {
		err = ErrParseYear
		return
	}

	month, err := strconv.ParseUint(dateParts[1], 10, 8)
	if err != nil {
		err = ErrParseMonth
		return
	}

	day, err := strconv.ParseUint(dateParts[2], 10, 8)
	if err != nil {
		err = ErrParseDay
		return
	}

	timestamp = time.Date(int(year), time.Month(month), int(day), hour, min, sec, 0, time.UTC)
	return timestamp, nil
}
