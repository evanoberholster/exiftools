package exiftool

type ExifThumbnail struct {
	offset      uint64
	size        uint64
	compression int
	width       uint
	height      uint
}
