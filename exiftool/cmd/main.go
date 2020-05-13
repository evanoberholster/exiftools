package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/evanoberholster/exiftools/exiftool"
	"github.com/evanoberholster/exiftools/exiftool/api"
	"github.com/evanoberholster/exiftools/exiftool/buffer"
	"github.com/evanoberholster/exiftools/exiftool/tags/ifd"
	"github.com/evanoberholster/exiftools/exiftool/tags/ifdexif"
	"github.com/evanoberholster/exiftools/exiftool/tags/ifdmknote"
)

func main() {
	var path string
	path = "../../../test/img/1.heic"
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	ti := exiftool.NewTagIndex()
	ti.Add("IFD", ifd.RootIfdTags)
	ti.Add("IFD/Exif", ifdexif.ExifIfdTags)
	ti.Add("IFD/Exif/Makernotes.Canon", ifdmknote.CanonIfdTags)
	ti.Add("IFD/GPS", ifd.GPSIfdTags)

	cb := buffer.NewCacheBuffer(f, 128*1024)

	er, err := exiftool.ParseExif2(cb)
	if err != nil {
		panic(err)
	}

	im := exiftool.NewIfdMapping()

	if _, err = im.LoadIfds(ifd.RootIfd, ifdexif.ExifIfd, ifd.GPSIfd); err != nil {
		fmt.Println(err)
	}
	//if _, err = im.LoadIfds(mknote.CanonMakernoteIfd); err != nil {
	//	fmt.Println(err)
	//}

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

		fmt.Printf("Path: %s \t| TagID: 0x%04x | %s   \t| %s \n", fqIfdPath, ite.TagID(), t.Name, ite.TagType())
		return nil
	}

	if err = er.Visit(ifd.RootIfd.Name, im, ti, visitor); err != nil {
		fmt.Println(err)
	}

	// Thumbnails
	for inum, i := range tags.GetIfd("IFD") {
		fmt.Println(inum, "__")
		fmt.Println(i[ifd.ImageWidth].GetInt(er))
		fmt.Println(i[ifd.ImageLength].GetInt(er))
		fmt.Println(i[ifd.StripByteCounts].GetInt(er))
		fmt.Println(i[ifd.StripOffsets].GetInt(er))
		fmt.Println(i[ifd.Compression].GetInt(er))
	}

	fmt.Println(tags.GetTag("IFD/Exif", 0, ifdexif.FocalLength).GetRational(er))
	fmt.Println(tags.FocalLength())
	fmt.Println(tags.Aperture())
	//fmt.Println(tags.GetIfds("IFD"))

	// Variables
	//fmt.Println(tags.XMLPacket())
	//fmt.Println(tags.GPSInfo())
	//fmt.Println(tags.Copyright())
	//fmt.Println(tags.Artist())
	//mk, err := tags.CameraMake()
	//
	//fmt.Println(mk, err, len(mk))
	//fmt.Println(tags.CameraModel())
	//fmt.Println(tags.CameraSerial())
	//fmt.Println(tags.LensMake())
	//fmt.Println(tags.LensModel())
	//fmt.Println(tags.LensSerial())
	//fmt.Println(tags.Dimensions())
	//fmt.Println(tags.ExposureProgram())
	//fmt.Println(tags.MeteringMode())
	//fmt.Println(tags.ShutterSpeed())
	//fmt.Println(tags.Aperture())
	//fmt.Println(tags.ISOSpeed())
	//fmt.Println(tags.FocalLength())
	//fmt.Println(tags.FocalLengthIn35mmFilm())
	//fmt.Println(tags.DateTime())
	//fmt.Println(tags.GPSTime())
	//fmt.Println(tags.ModifyDate())
	//fmt.Println(tags.CanonCameraSettings())
	//fmt.Println(tags.CanonShotInfo())
	//fmt.Println(tags.CanonFileInfo())
	//fmt.Println(tags.CanonAFInfo())

	offset, size, err := tags.Thumbnail()
	if err != nil {
		os.Exit(0)
	}
	thumb := make([]byte, size)
	if _, err := er.ReadAt(thumb, int64(offset)); err != nil {
		fmt.Println(err)
	} else {
		ioutil.WriteFile("image.jpg", thumb, 0644)
	}
}
