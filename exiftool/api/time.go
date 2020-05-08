package api

import (
	"fmt"
	"time"
	"strings"
	"strconv"
	"github.com/evanoberholster/exiftools/exiftool/tags/ifdexif"
)

var (
	ErrParseYear = fmt.Errorf("Error parsing Year")
	ErrParseMonth = fmt.Errorf("Error parsing Month")
	ErrParseDay = fmt.Errorf("Error parsing Day")
	ErrParseHour = fmt.Errorf("Error parsing Hour")
)

func (res Results) DateTime() (dateTime time.Time, err error) {
	dateRaw, err := res.GetTag("IFD/Exif", ifdexif.DateTimeOriginal).String()
	if err == nil || dateRaw != "" {
		dateTime, err = parseExifFullTimestamp(dateRaw)
		if err == nil || !dateTime.IsZero() {
			return dateTime, nil
		}
	}
	dateRaw, err = res.GetTag("IFD/Exif", ifdexif.DateTimeDigitized).String()
	if err == nil || dateRaw != "" {
		dateTime, err = parseExifFullTimestamp(dateRaw)
		if err == nil || !dateTime.IsZero() {
			return dateTime, nil
		}
	}
	return 
}

// parseExifFullTimestamp parses dates like "2018:11:30 13:01:49" into a UTC
// `time.Time` struct.
func parseExifFullTimestamp(fullTimestampPhrase string) (timestamp time.Time, err error) {
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
		err = ErrParseHour
		return
    }

    second, err := strconv.ParseUint(timeParts[2], 10, 8)
    if err != nil {
		err = ErrParseHour
		return
    }

    timestamp = time.Date(int(year), time.Month(month), int(day), int(hour), int(minute), int(second), 0, time.UTC)
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