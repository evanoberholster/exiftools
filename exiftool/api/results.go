package api

import (
	"encoding/binary"
	"fmt"

	"github.com/evanoberholster/exiftools/exiftool"
	"github.com/evanoberholster/exiftools/exiftool/exif"
)

// ExifResults contains an exifReader, the tags and offsets
// for parsing the values from the exifReader
type ExifResults struct {
	exifReader *exiftool.ExifReader
	ifdTagMap  map[string][]exif.TagMap
	byteOrder  binary.ByteOrder
}

func (itm ExifResults) String() string {
	return fmt.Sprintln(itm.ifdTagMap)
}

// NewTagMap - initialize each TagMap with allocations for 10 Tags
func NewTagMap() exif.TagMap {
	return make(exif.TagMap, 10)
}

// NewExifResults creates a new ExifResults with an
// active ExifReader. The ByteOrder is taken from the ExifReader
func NewExifResults(er *exiftool.ExifReader) ExifResults {
	return ExifResults{
		exifReader: er,
		byteOrder:  er.ByteOrder(),
		ifdTagMap:  make(map[string][]exif.TagMap, 4),
	}
}

// GetTag returns a Tag from the ExifResults
func (itm ExifResults) GetTag(fqIfdPath string, ifdIndex int8, tagID exif.TagID) exif.Tag {
	if ifd, ok := itm.ifdTagMap[fqIfdPath]; ok {
		if tag, f := ifd[ifdIndex][tagID]; f {
			return tag
		}
	}
	return exif.Tag{}
}

// GetIfd returns an Ifd of Tags from the ExifResults
func (itm ExifResults) GetIfd(fqIfdPath string) []exif.TagMap {
	return itm.ifdTagMap[fqIfdPath]
}

// AddTag adds a Tag to the ExifResults
func (itm ExifResults) AddTag(tag exif.Tag, ifdIndex int8, fqIfdPath string, tagID exif.TagID) {
	if _, ok := itm.ifdTagMap[fqIfdPath]; !ok {
		itm.ifdTagMap[fqIfdPath] = append(itm.ifdTagMap[fqIfdPath], NewTagMap())
	} else {
		if len(itm.ifdTagMap[fqIfdPath]) == int(ifdIndex) {
			itm.ifdTagMap[fqIfdPath] = append(itm.ifdTagMap[fqIfdPath], NewTagMap())
		}
	}
	itm.ifdTagMap[fqIfdPath][ifdIndex][tagID] = tag
}
