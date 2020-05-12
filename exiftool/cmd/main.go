package main

import (
	"fmt"
	"os"
	"time"

	"github.com/evanoberholster/exiftools/exiftool"
	"github.com/evanoberholster/exiftools/exiftool/api"
	"github.com/evanoberholster/exiftools/exiftool/buffer"
	"github.com/evanoberholster/exiftools/exiftool/tags/ifd"
	"github.com/evanoberholster/exiftools/exiftool/tags/ifdexif"
	"github.com/evanoberholster/exiftools/exiftool/tags/mknote"
)

func main() {
	var path string
	path = "../../../test/img/4.CR2"
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	cb := buffer.NewCacheBuffer(f, 128*1024)

	start := time.Now()
	er, err := exiftool.ParseExif2(cb)
	fmt.Println("Exif time: ", time.Since(start), er)
	if err != nil {
		panic(err)
		fmt.Println(err)
	}

	//   ifd.IopIfd
	//if _, err = im.LoadIfds(mknote.CanonMakernoteIfd); err != nil {
	//	fmt.Println(err)
	//}
	ti := exiftool.NewTagIndex()
	ti.Add("IFD", ifd.RootIfdTags)
	ti.Add("IFD/Exif", ifdexif.ExifIfdTags)
	ti.Add("IFD/Exif/Makernotes.Canon", mknote.CanonIfdTags)
	ti.Add("IFD/GPS", ifd.GPSIfdTags)
	//
	res := api.NewResults()
	start = time.Now()
	im := exiftool.NewIfdMapping()

	if _, err = im.LoadIfds(ifd.RootIfd, ifdexif.ExifIfd, ifd.GPSIfd); err != nil {
		fmt.Println(err)
	}
	//if _, err = im.LoadIfds(mknote.CanonMakernoteIfd); err != nil {
	//	fmt.Println(err)
	//}
	visitor := func(fqIfdPath string, ifdIndex int, ite *exiftool.IfdTagEntry) (err error) {
		// GetTag
		//fmt.Println(fqIfdPath, ite.TagID())
		t, err := ti.Get(fqIfdPath, ite.TagID())
		if err != nil {
			//fmt.Printf("Path: %s \t| TagID: 0x%04x  \t| %s %d %s\n", fqIfdPath, ite.TagID(), ite.TagType(), ifdIndex, err.Error())
			return nil
		}

		// TagValue
		value, err := ite.Value()
		if err != nil {
			//	fmt.Printf("%s \t| Value Error: %s \n", t.Name, err.Error())
			return nil
		}
		//if ifdIndex > 0 {
		//	fqIfdPath = fqIfdPath + strconv.Itoa(ifdIndex)
		//}
		res.Add(fqIfdPath, ite.TagID(), t.Name, ite.TagType(), value)

		fmt.Printf("Path: %s \t| TagID: 0x%04x | %s   \t| %s ", fqIfdPath, ite.TagID(), t.Name, ite.TagType())
		if ite.TagID() == ifd.XMLPacket {
			//str := strings.Replace(string(value.([]byte)), "\n", "", -1)
			//str = strings.Replace(str, "   ", "", -1)
			//fmt.Println(str)
		} else {
			fmt.Println(value)
		}

		return nil
	}

	//f.Seek(0, 0)
	//p, err := ioutil.ReadAll(f)

	if err = er.Visit(ifd.RootIfd.Name, im, ti, visitor); err != nil {
		fmt.Println(err)
	}
	fmt.Println("Decode Time: ", time.Since(start))
	start = time.Now()

	// Variables
	//fmt.Println(res.XMLPacket())
	//fmt.Println(res.GPSInfo())
	//fmt.Println(res.Copyright())
	//fmt.Println(res.Artist())
	fmt.Println(res.CameraMake())
	//fmt.Println(res.CameraModel())
	//fmt.Println(res.CameraSerial())
	//fmt.Println(res.LensMake())
	//fmt.Println(res.LensModel())
	//fmt.Println(res.LensSerial())
	//fmt.Println(res.Dimensions())
	//fmt.Println(res.DateTime())
	//fmt.Println(res.GPSTime())
	//fmt.Println(res.ModifyDate())
	//fmt.Println(res.ExposureProgram())
	//fmt.Println(res.MeteringMode())
	//fmt.Println(res.ShutterSpeed())
	//fmt.Println(res.Aperture())
	//fmt.Println(res.ISOSpeed())
	//fmt.Println(res.FocalLength())
	//fmt.Println(res.FocalLengthIn35mmFilm())
	//fmt.Println(res.CanonCameraSettings())
	//fmt.Println(res.CanonShotInfo())
	//fmt.Println(res.CanonFileInfo())
	//fmt.Println(res.CanonAFInfo())

	fmt.Println("Get Time: ", time.Since(start))
	//fmt.Println("Total Time: ", time.Since(begin))
	//visitor := func(fqIfdPath string, ifdIndex int, ite *exiftool.IfdTagEntry) (err error) {
	//	tagID := ite.TagID()
	//	//agType := ite.TagType()
	//
	//	//fmt.Println(fqIfdPath, tagID)
	//	// TagName = IfdPointer, TagID
	//	t, _ := tags.RootIfdTagNames[tagID]
	//
	//	value, err := ite.Value()
	//
	//	//if err != nil {
	//	//	fmt.Println(err)
	//	//}
	//	if t == "Make" && value == "Canon" {
	//		canon.LoadCanonMakerNote(im)
	//		fmt.Println("Hello World")
	//	}
	//
	//	//fmt.Printf("Path: %s \t| TagID: 0x%04x  \t|  %s \t| Type:%s %s\n", tagType, tagID, t, fqIfdPath, value)
	//	//fmt.Println(, )
	//	return nil
	//}
	//
	//err = eh.Visit(exif.IfdRootPath, tags.IfdPathStandard, im, ti, p, visitor)
	//if err != nil {
	//	fmt.Println(err)
	//}

}
