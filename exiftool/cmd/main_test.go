package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/evanoberholster/exiftools/exiftool"
	"github.com/evanoberholster/exiftools/exiftool/api"
	"github.com/evanoberholster/exiftools/exiftool/buffer"
	"github.com/evanoberholster/exiftools/exiftool/tags/ifd"
	"github.com/evanoberholster/exiftools/exiftool/tags/ifdexif"
	"github.com/evanoberholster/exiftools/exiftool/tags/ifdmknote"
)

const (
	testPath = "../../../test/img/2.CR2"
)

var ti exiftool.TagIndex

func init() {
	LoadTagIndex()
}

func LoadTagIndex() {
	ti = exiftool.NewTagIndex()
	ti.Add("IFD", ifd.RootIfdTags)
	ti.Add("IFD/Exif", ifdexif.ExifIfdTags)
	ti.Add("IFD/Exif/Makernotes.Canon", ifdmknote.CanonIfdTags)
	ti.Add("IFD/GPS", ifd.GPSIfdTags)
}

func BenchmarkExif200(b *testing.B) {
	var err error

	f, err := os.Open(testPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	im := exiftool.NewIfdMapping()
	if _, err = im.LoadIfds(ifd.RootIfd, ifdexif.ExifIfd, ifd.GPSIfd, ifd.IopIfd); err != nil {
		fmt.Println(err)
	}
	im.LoadIfds(ifdmknote.LoadMakernotesIfd("Canon"))

	b.ReportAllocs()
	b.ResetTimer()
	f.Seek(0, 0)
	cb := buffer.NewCacheBuffer(f, 128*1024)
	er, err := exiftool.ParseExif2(cb)
	if err != nil {
		b.Fatal(err)
	}

	tags := api.NewExifResults(er)
	visitor := func(fqIfdPath string, ifdIndex int, ite *exiftool.IfdTagEntry) (err error) {
		// GetTag
		t, err := ti.Get(fqIfdPath, ite.TagID())
		if err != nil {
			return nil
		}
		// SetTag
		ite.SetTag(&t)

		// AddTag
		tags.AddTag(t, int8(ifdIndex), fqIfdPath, ite.TagID())
		return nil
	}

	if err = er.Visit(ifd.RootIfd.Name, im, ti, visitor); err != nil {
		fmt.Println(err)
	}
	for i := 0; i < b.N; i++ {
		tags.XMLPacket()
		tags.GPSInfo()
		tags.Copyright()
		tags.Artist()
		tags.CameraMake()
		tags.CameraModel()
		tags.CameraSerial()
		tags.LensMake()
		tags.LensModel()
		tags.LensSerial()
		tags.Dimensions()
		tags.ExposureProgram()
		tags.MeteringMode()
		tags.ShutterSpeed()
		tags.Aperture()
		tags.ISOSpeed()
		tags.FocalLength()
		tags.FocalLengthIn35mmFilm()
		tags.DateTime()
		tags.GPSTime()
		tags.ModifyDate()
		tags.CanonCameraSettings()
		tags.CanonShotInfo()
		tags.CanonFileInfo()
		tags.CanonAFInfo()
	}
}
