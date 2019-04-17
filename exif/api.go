package exif
import (
	"fmt"
	"time"
	"strconv"
	"github.com/evanoberholster/exif/tiff"
)

func (x *Exif) Software() (string, error) {
	tag, err := x.Get(Software)
	if err != nil {
		return "", err
	}
	return tag.StringVal()
}

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

func calcTimeHelper(n int64, d int64, err error) (string) {
	a := int(n/d)
	if a > 10 {
		return strconv.Itoa(a)
	}
	return fmt.Sprintf("0%d",a)
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

	hour := calcTimeHelper(tS.Rat2(0))
 	min := calcTimeHelper(tS.Rat2(1))

	secN, secD, err := tS.Rat2(2)
	if err != nil {
		return time.ParseInLocation(exifTimeLayout, dateStr, time.UTC)
	} else {
		exifTimeLayout = "2006:01:02 15:04:05.999"
		sec := float32(secN)/float32(secD)
		if sec < 10 {
			dateStr = fmt.Sprintf("%v %v:%v:0%.3f", dateStr, hour, min, sec)
		} else {
			dateStr = fmt.Sprintf("%v %v:%v:%.3f", dateStr, hour, min, sec)
		}
	}
	return time.ParseInLocation(exifTimeLayout, dateStr, time.UTC)
}

func (x *Exif)FocalLength() (float32, error) {
	tag, err := x.Get(FocalLength)
	if err != nil {
		return 0, fmt.Errorf("Cannot parse Focal Length: %v", err)
	}
	
	switch tag.Type {

	case tiff.DTRational:
			num, denom, err := tag.Rat2(0)
			if err != nil { return 0, err }

			return float32(num)/float32(denom), nil
		
	case tiff.DTShort:
			a, _ := tag.Int(0)
			b, err := tag.Int(1)
			if err != nil { return 0, err }
			l := len(strconv.Itoa(b))

			if a == 0 { return float32(b), nil }
			if a == 2 {
				if l == 4 {
					return float32(b) / float32(1000), nil
				}
				if l == 3 {
					return float32(b) / float32(100), nil
				}
			}		
	}
	return 0, fmt.Errorf("Cannot parse FocalLength")
}