package exif

// IfdItem is a container to define each IFD
type IfdItem struct {
	IfdPath IfdPath
	TagID   TagID
	Name    string
}

// Valid returns true if the IfdItem has a Name
func (ifd IfdItem) Valid() bool {
	return ifd.Name != ""
}
