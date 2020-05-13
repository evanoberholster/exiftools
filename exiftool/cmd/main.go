package main

import (
	"fmt"
	"io/ioutil"
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
	path = "../../../test/img/1.heic"
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
	//res := api.NewResults()
	start = time.Now()
	im := exiftool.NewIfdMapping()

	if _, err = im.LoadIfds(ifd.RootIfd, ifdexif.ExifIfd, ifd.GPSIfd); err != nil {
		fmt.Println(err)
	}
	//if _, err = im.LoadIfds(mknote.CanonMakernoteIfd); err != nil {
	//	fmt.Println(err)
	//}
	//tags := make([]exif.Tag, 0, 2)
	tags := api.NewIfdTagMap(er)
	visitor := func(fqIfdPath string, ifdIndex int, ite *exiftool.IfdTagEntry) (err error) {
		// GetTag
		//fmt.Println(fqIfdPath, ite.TagID())
		t, err := ti.Get(fqIfdPath, ite.TagID())
		if err != nil {
			//fmt.Printf("Path: %s \t| TagID: 0x%04x  \t| %s %d %s\n", fqIfdPath, ite.TagID(), ite.TagType(), ifdIndex, err.Error())
			return nil
		}

		// TagValue
		//value, err := ite.Value()
		//if err != nil {
		//	//	fmt.Printf("%s \t| Value Error: %s \n", t.Name, err.Error())
		//	return nil
		//}
		ite.SetTag(&t)
		tags.AddTag(t, int8(ifdIndex), fqIfdPath, ite.TagID())
		//if ite.TagID() == ifd.Make {
		//
		//	tags = append(tags, t)
		//}

		//if ifdIndex > 0 {
		//	fqIfdPath = fqIfdPath + strconv.Itoa(ifdIndex)
		//}
		//res.Add(fqIfdPath, ite.TagID(), t.Name, ite.TagType(), value)

		fmt.Printf("Path: %s \t| TagID: 0x%04x | %s   \t| %s \n", fqIfdPath, ite.TagID(), t.Name, ite.TagType())
		//if ite.TagID() == ifd.XMLPacket {
		//	//str := strings.Replace(string(value.([]byte)), "\n", "", -1)
		//	//str = strings.Replace(str, "   ", "", -1)
		//	//fmt.Println(str)
		//} else {
		//	//fmt.Println(value)
		//}

		return nil
	}

	//f.Seek(0, 0)
	//p, err := ioutil.ReadAll(f)

	if err = er.Visit(ifd.RootIfd.Name, im, ti, visitor); err != nil {
		fmt.Println(err)
	}
	fmt.Println("Decode Time: ", time.Since(start))
	start = time.Now()

	//fmt.Println(tags["IFD"][ifd.Make])

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
	fmt.Println("Get Time: ", time.Since(start))

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
