package api

import (
	"github.com/evanoberholster/exiftools/exiftool/exif"
)

func NewResults() Results {
	var res Results
	res = make(map[string]map[exif.TagID]*ExifTag)
	return res
}

type Results map[string]map[exif.TagID]*ExifTag

//func (res Results) String() string {
//	return fmt.Sprintf("Ifds: %d", len(res))
//}

// Add - adds an ExifTag to Results
func (res Results) Add(fqIfdPath string, tagID exif.TagID, tagName string, tagType exif.TagType, value interface{}) {
	if _, ok := res[fqIfdPath]; !ok {
		res[fqIfdPath] = make(map[exif.TagID]*ExifTag)
	}

	res[fqIfdPath][tagID] = &ExifTag{
		TagID:   tagID,
		tagName: tagName,
		tagType: tagType,
		value:   value}
}

func (res Results) GetTag(fqIfdPath string, tagID exif.TagID) *ExifTag {
	if ifd, ok := res[fqIfdPath]; ok {
		if tag, f := ifd[tagID]; f {
			return tag
		}
	}
	return nil
}
