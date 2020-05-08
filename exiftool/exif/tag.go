package exif

type TagID uint16

type IfdPath []TagID

// Tag - is an Exif Tag
type Tag struct {
	//ID      uint16
	Name string
	//IfdPath IfdPath
	Type TagType
}
