package api

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"testing"

	exifold "github.com/evanoberholster/exiftools/exif"
	"github.com/evanoberholster/exiftools/exiftool"
	"github.com/evanoberholster/exiftools/exiftool/tags/ifd"
	"github.com/evanoberholster/exiftools/exiftool/tags/ifdexif"
	"github.com/evanoberholster/exiftools/exiftool/tags/mknote"
	mknoteold "github.com/evanoberholster/exiftools/mknote"
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

func BenchmarkExifDecode200(b *testing.B) {
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

	res := NewResults()

	visitor := func(fqIfdPath string, ifdIndex int, ite *exiftool.IfdTagEntry) (err error) {
		// GetTag
		t, err := ti.Get(fqIfdPath, ite.TagID())
		if err != nil {
			return nil
		}

		// TagValue
		value, err := ite.Value()
		if err != nil {
			return nil
		}
		if ifdIndex > 0 {
			fqIfdPath = fqIfdPath + strconv.Itoa(ifdIndex)
		}
		res.Add(fqIfdPath, ite.TagID(), t.Name, ite.TagType(), value)

		return nil
	}

	b.ReportAllocs()
	b.ResetTimer()
	var eh *exiftool.ExifHeader
	var p []byte
	for i := 0; i < b.N; i++ {
		f.Seek(0, 0)
		eh, err = exiftool.ParseExif(f)
		if err != nil {
			panic(err)
		}
		f.Seek(0, 0)
		p, err = ioutil.ReadAll(f)
		if err = eh.Decode(ifd.RootIfd.Name, im, ti, p, visitor); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkExifDecodeOld200(b *testing.B) {
	var err error

	f, err := os.Open(testPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	exifold.RegisterParsers(mknoteold.All...)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		f.Seek(0, 0)
		if _, err := exifold.DecodeWithParseHeader(f); err != nil {
			b.Fatal(err)
		}
	}
}
