package exif

import (
	"encoding/binary"
	"io"
)

type ExifReader interface {
	io.ReaderAt
	ByteOrder() binary.ByteOrder
}
