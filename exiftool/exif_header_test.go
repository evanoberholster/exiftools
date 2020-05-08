package exiftool

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/dsoprea/go-exif"
)

var path = "../../test/img/1.jpg"

func BenchmarkParseExifHeader200(b *testing.B) {
	var err error

	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		f.Seek(0, 0)
		_, _ = parseExifHeader(f)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkParseExif2Header200(b *testing.B) {
	var err error

	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		f.Seek(0, 0)
		_, _ = parseExifHeader2(f)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkCurrent200(b *testing.B) {
	var err error

	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		f.Seek(0, 0)
		data, _ := ioutil.ReadAll(f)
		_, _ = exif.SearchAndExtractExif(data)
		if err != nil {
			b.Fatal(err)
		}
	}
}
