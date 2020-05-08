package exiftool

import (
	"encoding/binary"

	"github.com/evanoberholster/exiftools/exiftool/exif"
)

// IfdTagEntry refers to a tag in the loaded EXIF block.
type IfdTagEntry struct {
	tagID          exif.TagID
	tagIndex       int
	tagType        exif.TagType
	unitCount      uint32
	valueOffset    uint32
	rawValueOffset []byte

	// childIfdName is the right most atom in the IFD-path. We need this to
	// construct the fully-qualified IFD-path.
	childIfdName string

	// childIfdPath is the IFD-path of the child if this tag represents a child
	// IFD.
	childIfdPath string

	// childFqIfdPath is the IFD-path of the child if this tag represents a
	// child IFD. Includes indices.
	childFqIfdPath string

	// TODO(dustin): !! IB's host the child-IBs directly in the tag, but that's not the case here. Refactor to accomodate it for a consistent experience.

	// ifdPath is the IFD that this tag belongs to.
	ifdPath string

	isUnhandledUnknown bool

	addressableData []byte
	exifReader      *ExifReader
	byteOrder       binary.ByteOrder
}

func newIfdTagEntry(ifdPath string, tagID exif.TagID, tagIndex int, tagType exif.TagType, unitCount uint32, valueOffset uint32, rawValueOffset []byte, exifReader *ExifReader, byteOrder binary.ByteOrder) *IfdTagEntry {
	return &IfdTagEntry{
		ifdPath:        ifdPath,
		tagID:          tagID,
		tagIndex:       tagIndex,
		tagType:        tagType,
		unitCount:      unitCount,
		valueOffset:    valueOffset,
		rawValueOffset: rawValueOffset,
		exifReader:     exifReader,
		byteOrder:      byteOrder,
	}
}

// ChildFqIfdPath returns the complete path of the child IFD along with the
// numeric suffixes differentiating sibling occurrences of the same type. "0"
// indices are omitted.
func (ite *IfdTagEntry) ChildFqIfdPath() string {
	return ite.childFqIfdPath
}

// getValueOffset is the four-byte offset converted to an integer to point to
// the location of its value in the EXIF block. The "get" parameter is obviously
// used in order to differentiate the naming of the method from the field.
func (ite *IfdTagEntry) getValueOffset() uint32 {
	return ite.valueOffset
}

// SetChildIfd sets child-IFD information (if we represent a child IFD).
func (ite *IfdTagEntry) SetChildIfd(childFqIfdPath, childIfdPath, childIfdName string) {
	ite.childFqIfdPath = childFqIfdPath
	ite.childIfdPath = childIfdPath
	ite.childIfdName = childIfdName
}

// ChildIfdName returns the name of the child IFD
func (ite *IfdTagEntry) ChildIfdName() string {
	return ite.childIfdName
}

// ChildIfdPath returns the path of the child IFD.
func (ite *IfdTagEntry) ChildIfdPath() string {
	return ite.childIfdPath
}

// TagID returns the ID of the tag that we represent. The combination of
// (IfdPath(), TagId()) is unique.
func (ite *IfdTagEntry) TagID() exif.TagID {
	return ite.tagID
}

// TagType is the type of value for this tag.
func (ite *IfdTagEntry) TagType() exif.TagType {
	return ite.tagType
}

// Value returns the specific, parsed, typed value from the tag.
func (ite *IfdTagEntry) Value() (value interface{}, err error) {
	defer func() {
		if state := recover(); state != nil {
			err = state.(error)
		}
	}()

	valueContext := ite.getValueContext()

	if ite.tagType == exif.TypeUndefined {
		//var err error

		//value, err = exif.Decode(valueContext)
		//if err != nil {
		//	if err == exif.ErrUnhandledUndefinedTypedTag || err == exif.ErrUnparseableValue {
		//		return nil, err
		//	}
		//
		//	log.Panic(err)
		//}
	} else {
		value, err = valueContext.Values()
		if err != nil {
			panic(err)
		}
	}

	return value, nil
}

func (ite *IfdTagEntry) String() (string, error) {
	valueContext := ite.getValueContext()
	return valueContext.ReadASCII()
}

func (ite *IfdTagEntry) getValueContext() *ValueContext {
	return NewValueContext(
		ite.ifdPath,
		ite.tagID,
		ite.unitCount,
		ite.valueOffset,
		ite.rawValueOffset,
		ite.exifReader,
		ite.tagType,
		ite.byteOrder)
}
