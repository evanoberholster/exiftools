package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/evanoberholster/exiftools/exiftool"
	"github.com/evanoberholster/exiftools/exiftool/tags/ifd"
	"github.com/evanoberholster/exiftools/exiftool/tags/ifdexif"
	"github.com/evanoberholster/exiftools/exiftool/tags/mknote"
)

func main() {
	var path string
	path = "../../test/img/13.jpg"
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	start := time.Now()
	eh, err := exiftool.ParseExif(f)
	fmt.Println("Exif time: ", time.Since(start), eh)
	if err != nil {
		panic(err)
		fmt.Println(err)
	}

	im := exiftool.NewIfdMapping()

	if _, err = im.LoadIfds(ifd.RootIfd); err != nil {
		fmt.Println(err)
	}
	if _, err = im.LoadIfds(ifdexif.ExifIfd); err != nil {
		fmt.Println(err)
	}
	if _, err = im.LoadIfds(ifd.GPSIfd); err != nil {
		fmt.Println(err)
	}
	if _, err = im.LoadIfds(ifd.IopIfd); err != nil {
		fmt.Println(err)
	}
	//if _, err = im.LoadIfds(mknote.CanonMakernoteIfd); err != nil {
	//	fmt.Println(err)
	//}
	ti := exiftool.NewTagIndex()
	ti.Add("IFD", ifd.RootIfdTags)
	ti.Add("IFD/Exif", ifdexif.ExifIfdTags)
	ti.Add("IFD/Exif/CanonMKnote", mknote.CanonIfdTags)
	ti.Add("IFD/GPS", ifd.GPSIfdTags)

	visitor := func(fqIfdPath string, ifdIndex int, ite *exiftool.IfdTagEntry) (err error) {

		// GetTag
		t, err := ti.Get(fqIfdPath, ite.TagID())
		if err != nil {
			//fmt.Printf("Path: %s \t| TagID: 0x%04x  \t| %s %d %s\n", fqIfdPath, ite.TagID(), ite.TagType(), ifdIndex, err.Error())
			return nil
		}

		// TagValue
		value, err := ite.Value()
		if err != nil {
			fmt.Printf("%s \t| Value Error: %s \n", t.Name, err.Error())
			return nil
		}
		fmt.Printf("Path: %s \t| TagID: 0x%04x | %s   \t| %s ", fqIfdPath, ite.TagID(), t.Name, ite.TagType())
		fmt.Printf("\t %s \n", value)

		return nil
	}
	f.Seek(0, 0)
	p, err := ioutil.ReadAll(f)
	start = time.Now()
	if err = eh.Decode(ifd.RootIfd.Name, im, ti, p, visitor); err != nil {
		fmt.Println(err)
	}
	fmt.Println("Decode Time: ", time.Since(start))

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
