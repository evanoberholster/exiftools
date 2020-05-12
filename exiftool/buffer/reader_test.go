package buffer

import (
	"os"
	"testing"

	"github.com/evanoberholster/exiftools/exiftool"
)

var path = "../../../test/img/20.jpg"

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
		_, _ = exiftool.ParseExif(f)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkParseExifHeaderNew200(b *testing.B) {
	var err error

	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	b.ReportAllocs()
	b.ResetTimer()
	var cb *CacheBuffer
	for i := 0; i < b.N; i++ {
		f.Seek(0, 0)
		cb = NewCacheBuffer(f, 128*1024)

		_, _ = exiftool.ParseExif(cb)
		if err != nil {
			b.Fatal(err)
		}
	}
}
