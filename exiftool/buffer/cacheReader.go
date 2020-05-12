package buffer

import (
	"errors"
	"io"
)

// CacheBuffer -
type CacheBuffer struct {
	// buf is the Buffer of the CacheBuffer
	buf []byte
	// the Read() offset for the Buffer
	offset int

	// Source for the CacheBuffer
	src io.Reader
	// Source EOF is true when the buffer contains all the src reader
	srcEOF bool
	// Position of the CacheBuffer
	bytesBuffered int

	// bufSize is the amount that should be read on each read
	bufSize int
}

// NewCacheBuffer creates a new CacheBuffer with a Reader and bufferSize
func NewCacheBuffer(src io.Reader, bufSize int) *CacheBuffer {
	if bufSize <= 0 {
		bufSize = 128 * 1024 // 64Kb
	}
	// fillBuffer inital
	return &CacheBuffer{
		buf:     make([]byte, 0, 0),
		src:     src,
		bufSize: bufSize,
	}
}

func (cb *CacheBuffer) fillBuffer(offset int, size int) (n int, err error) {
	// Fetch in the size of a buffer
	l := size + offset - cb.bytesBuffered
	fetchSize := l / cb.bufSize
	if l%cb.bufSize > 0 {
		fetchSize++
	}

	temp := make([]byte, fetchSize*cb.bufSize)
	n, err = cb.src.Read(temp)
	cb.bytesBuffered += n
	cb.buf = append(cb.buf, temp[:n]...)
	if err != nil {
		if err == io.EOF {
			cb.srcEOF = true
			if n > 0 {
				return n, nil
			}
		}
		return n, err
	}
	return n, nil
}

// Buffer returns the []byte from CacheBuffer
// Warning: only returns what has been read so far.
func (cb CacheBuffer) Buffer() []byte {
	return cb.buf[:cb.bytesBuffered]
}

func (cb *CacheBuffer) Read(p []byte) (n int, err error) {
	if len(p) <= 0 { // Read an empty byte slice
		return 0, nil
	}

	if len(p)+cb.offset-cb.bytesBuffered > 0 && !cb.srcEOF {
		// fill Buffer
		if n, err = cb.fillBuffer(cb.offset, len(p)); err != nil {
			return 0, err
		}
	}

	n = copy(p, cb.buf[cb.offset:])
	cb.offset += n

	return n, nil
}

func (cb *CacheBuffer) ReadAt(p []byte, off int64) (n int, err error) {
	if len(p) <= 0 { // Read an empty byte slice
		return 0, nil
	}
	if off < 0 {
		return 0, errors.New("CacheBuffer.ReadAt: negative offset")
	}

	if len(p)+int(off)-cb.bytesBuffered > 0 && !cb.srcEOF {
		// fill Buffer
		if n, err = cb.fillBuffer(int(off), len(p)); err != nil {
			//return 0, err
		}
	}

	n = copy(p, cb.buf[off:])
	if n < len(p) {
		return n, io.EOF
	}

	return n, nil
}
