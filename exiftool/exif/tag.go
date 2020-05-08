package exif

// TagID is the uint16 representation of an IFD tag
type TagID uint16

// IfdPath is an array of TagID representing an IFD
type IfdPath []TagID

// Tag - is an Exif Tag
type Tag struct {
	Name string
	Type TagType
}

type TagMap map[TagID]Tag
