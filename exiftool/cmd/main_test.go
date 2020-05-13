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
	"github.com/evanoberholster/exiftools/exiftool/tags/mknote"
)

const (
	testPath = "../../../test/img/2.CR2"
)

var ti *exiftool.TagIndex

func init() {
	LoadTagIndex()
}

func LoadTagIndex() {
	ti = exiftool.NewTagIndex()
	ti.Add("IFD", ifd.RootIfdTags)
	ti.Add("IFD/Exif", ifdexif.ExifIfdTags)
	ti.Add("IFD/Exif/Makernotes.Canon", mknote.CanonIfdTags)
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
	im.LoadIfds(mknote.LoadMakernotesIfd("Canon"))

	b.ReportAllocs()
	b.ResetTimer()
	f.Seek(0, 0)
	cb := buffer.NewCacheBuffer(f, 128*1024)
	er, err := exiftool.ParseExif2(cb)
	if err != nil {
		b.Fatal(err)
	}

	//res := api.NewResults()
	//tags := make([]exif.Tag, 0, 100)
	tags := api.NewIfdTagMap(er)
	visitor := func(fqIfdPath string, ifdIndex int, ite *exiftool.IfdTagEntry) (err error) {
		// GetTag
		t, err := ti.Get(fqIfdPath, ite.TagID())
		if err != nil {
			return nil
		}
		// TagValue
		//value, err := ite.Value()
		//if err != nil {
		//	return nil
		//}
		ite.SetTag(&t)
		tags.AddTag(t, int8(ifdIndex), fqIfdPath, ite.TagID())

		return nil
	}

	if err = er.Visit(ifd.RootIfd.Name, im, ti, visitor); err != nil {
		fmt.Println(err)
		//b.Fatal(err)
	}
	for i := 0; i < b.N; i++ {
		//for _, i := range tags.GetIfd("IFD") {
		//	//fmt.Println(inum, "__")
		//	i[ifd.ImageWidth].GetInt(er)
		//	i[ifd.ImageLength].GetInt(er)
		//	i[ifd.StripByteCounts].GetInt(er)
		//	i[ifd.StripOffsets].GetInt(er)
		//	i[ifd.Compression].GetInt(er)
		//	//
		//	// StripOffsets
		//	// Compression
		//}
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
		//_, _ = tags.GetTag("IFD", 0, ifd.ImageLength).GetInt(er)
	}
}
