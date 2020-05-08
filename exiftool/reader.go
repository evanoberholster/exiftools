package exiftool

import (
	"errors"
	"io"
	"os"
)

// Cachedata
type Cachedata struct {
	buf []byte
	//reader io.Reader
	//pos int64
}

type ExifReader struct {
	header   *ExifHeader
	cache    *Cachedata
	index    int64
	offset   int64
	isClosed bool
}

func NewExifReader(eh *ExifHeader, buf []byte) *ExifReader {
	return &ExifReader{
		header: eh,
		cache: &Cachedata{
			buf: buf[eh.foundAt:],
		},
		offset: 0,
	}
}

func NewReader(buf []byte) *ExifReader {
	cache := &Cachedata{
		buf: buf,
	}
	return &ExifReader{
		cache: cache,
	}
}

func (er ExifReader) NewReader(offset int64) *ExifReader {
	if er.isClosed {
		return nil
	}
	return &ExifReader{
		cache:  er.cache,
		offset: offset,
	}
}

// Read reads the next len(p) bytes from the CacheReader buffer or until the
// buffer is drained. The return value n is the number of bytes read. If the
// buffer has no data to return, err is io.EOF (unless len(p) is zero);
// otherwise it is nil.
func (er *ExifReader) Read(p []byte) (n int, err error) {
	if er.empty() {
		// Buffer is empty
		if len(p) == 0 {
			return 0, nil
		}
		return 0, io.EOF
	}

	n = copy(p, er.cache.buf[er.offset:])
	er.offset += int64(n)

	return n, nil
}

func (er *ExifReader) ReadAt(p []byte, off int64) (n int, err error) {
	if er.isClosed {
		return 0, os.ErrClosed
	}
	if off < 0 {
		return 0, errors.New("ExifReader.ReadAt: negative offset")
	}
	reqLen := len(p)
	buffLen := int64(len(er.cache.buf))
	//buffLen := int64(f.Buff.Len())
	if off >= buffLen {
		return 0, io.EOF
	}

	n = copy(p, er.cache.buf[off:])
	if n < reqLen {
		err = io.EOF
	}
	return n, err
}

// empty reports whether the unread portion of the buffer is empty.
func (cr *ExifReader) empty() bool { return len(cr.cache.buf) <= int(cr.offset) }
