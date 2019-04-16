package exif
import (
	"time"
	"fmt"
)

func (x *Exif) GPSAltitude() (float32, error) {
	alt, err := x.Get(GPSAltitude)
	if err != nil { return 0, err }

	altRef, err := x.Get(GPSAltitudeRef)
	if err != nil { return 0, err }

	ref, err := altRef.Int(0)
	if err != nil { return 0, fmt.Errorf("Cannot parse GPS Altitude: %v", err) }

	aN, aD, err := alt.Rat2(0)
	if err != nil { return 0, fmt.Errorf("Cannot parse GPS Altitude: %v", err) }
	
	a := float32(aN/aD)
	if ref == 1 { a = a*-1 } // Adjust for "Below Sea Level"
	return a, nil
}


func (x *Exif) GPSTimeStamp() (time.Time, error) {
	var dt time.Time
	dS, err := x.Get(GPSDateStamp)
	if err != nil {
		return dt, err
	}
	tS, err := x.Get(GPSTimeStamp)
	if err != nil {
		return dt, err
	}
	exifTimeLayout := "2006:01:02"

	dateStr, err := dS.StringVal()
	hourN, hourD, err := tS.Rat2(0)
	minN, minD, err := tS.Rat2(1)
	secN, secD, err := tS.Rat2(2)
	if err != nil {
		return time.ParseInLocation(exifTimeLayout, dateStr, time.UTC)
	} else {
		exifTimeLayout = "2006:01:02 15:04:05.999"
		sec := float32(secN)/float32(secD)
		if sec < 10 {
			dateStr = fmt.Sprintf("%v %d:%d:0%.3f", dateStr, int(hourN/hourD), int(minN/minD), sec)
		} else {
			dateStr = fmt.Sprintf("%v %d:%d:%.3f", dateStr, int(hourN/hourD), int(minN/minD), sec)
		}
	}
	
	return time.ParseInLocation(exifTimeLayout, dateStr, time.UTC)
}