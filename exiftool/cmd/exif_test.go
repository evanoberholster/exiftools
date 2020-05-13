package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/dsoprea/go-exif/v2"
	exifcommon "github.com/dsoprea/go-exif/v2/common"
	log "github.com/dsoprea/go-logging"
	exifold "github.com/evanoberholster/exiftools/exif"
	"github.com/evanoberholster/exiftools/exiftool"
	"github.com/evanoberholster/exiftools/exiftool/api"
	"github.com/evanoberholster/exiftools/exiftool/buffer"
	"github.com/evanoberholster/exiftools/exiftool/tags/ifd"
	"github.com/evanoberholster/exiftools/exiftool/tags/ifdexif"
	"github.com/evanoberholster/exiftools/exiftool/tags/ifdmknote"
	mknoteold "github.com/evanoberholster/exiftools/mknote"
)

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
	im.LoadIfds(ifdmknote.LoadMakernotesIfd("Canon"))

	b.ReportAllocs()
	b.ResetTimer()
	var er *exiftool.ExifReader
	for i := 0; i < b.N; i++ {
		f.Seek(0, 0)
		cb := buffer.NewCacheBuffer(f, 128*1024)
		er, err = exiftool.ParseExif2(cb)
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

func BenchmarkExifDecodeOprea200(b *testing.B) {
	var err error

	type IfdEntry struct {
		IfdPath     string                      `json:"ifd_path"`
		FqIfdPath   string                      `json:"fq_ifd_path"`
		IfdIndex    int                         `json:"ifd_index"`
		TagId       uint16                      `json:"tag_id"`
		TagName     string                      `json:"tag_name"`
		TagTypeId   exifcommon.TagTypePrimitive `json:"tag_type_id"`
		TagTypeName string                      `json:"tag_type_name"`
		UnitCount   uint32                      `json:"unit_count"`
		Value       interface{}                 `json:"value"`
		ValueString string                      `json:"value_string"`
	}

	f, err := os.Open(testPath)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	// Run the parse.
	im := exif.NewIfdMappingWithStandard()
	ti := exif.NewTagIndex()

	//exifold.RegisterParsers(mknoteold.All...)
	f.Seek(0, 0)
	data, err := ioutil.ReadAll(f)
	log.PanicIf(err)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {

		rawExif, err := exif.SearchAndExtractExif(data)
		if err != nil {
			if err == exif.ErrNoExif {
				fmt.Printf("No EXIF data.\n")
			}
			b.Fatal(err)
		}

		entries := make([]IfdEntry, 20)
		visitor := func(fqIfdPath string, ifdIndex int, ite *exif.IfdTagEntry) (err error) {
			defer func() {
				if state := recover(); state != nil {
					err = log.Wrap(state.(error))
					b.Fatal(err)
				}
			}()
			tagId := ite.TagId()
			tagType := ite.TagType()

			ifdPath, err := im.StripPathPhraseIndices(fqIfdPath)
			if err != nil {
				b.Fatal(err)
			}

			it, err := ti.Get(ifdPath, tagId)
			if err != nil {
				if log.Is(err, exif.ErrTagNotFound) {
					//mainLogger.Warningf(nil, "Unknown tag: [%s] (%04x)", ifdPath, tagId)
					return nil
				} else {
					b.Fatal(err)
				}
			}

			value, err := ite.Value()
			if err != nil {
				if log.Is(err, exifcommon.ErrUnhandledUndefinedTypedTag) == true {
					//mainLogger.Warningf(nil, "Skipping non-standard undefined tag: [%s] (%04x)", ifdPath, tagId)
					return nil
				} //else if err == exifundefined.ErrUnparseableValue {
				//mainLogger.Warningf(nil, "Skipping unparseable undefined tag: [%s] (%04x)", ifdPath, tagId)
				//return nil
				//}
				b.Fatal(err)
			}

			valueString, err := ite.FormatFirst()
			if err != nil {
				b.Fatal(err)
			}

			entry := IfdEntry{
				IfdPath:     ifdPath,
				FqIfdPath:   fqIfdPath,
				IfdIndex:    ifdIndex,
				TagId:       tagId,
				TagName:     it.Name,
				TagTypeId:   tagType,
				TagTypeName: tagType.String(),
				UnitCount:   ite.UnitCount(),
				Value:       value,
				ValueString: valueString,
			}

			entries = append(entries, entry)

			return nil
		}

		_, err = exif.Visit(exifcommon.IfdStandard, im, ti, rawExif, visitor)
		if err != nil {
			b.Fatal(err)
		}
	}
}
