package exiftool

import (
	"errors"
	"fmt"

	"github.com/evanoberholster/exiftools/exiftool/exif"
)

// Enumerator Errors
var (
	ErrNoThumbnail     = errors.New("No thumbnail")
	ErrNoGpsTags       = errors.New("No GPS Tags")
	ErrTagTypeNotValid = errors.New("Tag Type invalid")
)

const (
	// ExifAddressableAreaStart is the absolute offset in the file that all
	// offsets are relative to.
	ExifAddressableAreaStart = uint32(0x0)
)

// IfdTagEnumerator knows how to decode an IFD and all of the tags it
// describes.
//
// The IFDs and the actual values can float throughout the EXIF block, but the
// IFD itself is just a minor header followed by a set of repeating,
// statically-sized records. So, the tags (though notnecessarily their values)
// are fairly simple to enumerate.
type IfdTagEnumerator struct {
	exifReader *ExifReader
	//byteOrder  binary.ByteOrder // Prefer exifReader.ByteOrder
	ifdOffset uint32
	rawBytes  []byte
}

// IfdEnumerate -
type IfdEnumerate struct {
	exifReader    *ExifReader
	currentOffset uint32
	tagIndex      *TagIndex
	ifdMapping    *IfdMapping
}

func newIfdEnumerate(er *ExifReader, ifdMapping *IfdMapping, tagIndex *TagIndex) *IfdEnumerate {
	return &IfdEnumerate{
		exifReader: er,
		ifdMapping: ifdMapping,
		tagIndex:   tagIndex,
	}
}

// getIfdEnumerate creates a new IFD Enumerate
func (er *ExifReader) getIfdEnumerate(ifdMapping *IfdMapping, tagIndex *TagIndex) *IfdEnumerate {
	return &IfdEnumerate{
		exifReader: er,
		ifdMapping: ifdMapping,
		tagIndex:   tagIndex,
	}
}

// getTagEnumerator creates a new IFD Tag Enumerator
func (ie *IfdEnumerate) getTagEnumerator(ifdOffset uint32) (enumerator *IfdTagEnumerator) {
	return &IfdTagEnumerator{
		exifReader: ie.exifReader.SubReader(int64(ifdOffset)),
		ifdOffset:  ifdOffset,
		rawBytes:   make([]byte, 4),
	}
}

// Scan enumerates the different EXIF blocks (called IFDs). `rootIfdName` will
// be "IFD" in the TIFF standard.
func (ie *IfdEnumerate) Scan(rootIfdName string, ifdOffset uint32, visitor TagVisitorFn) (err error) {
	defer func() {
		if state := recover(); state != nil {
			err = state.(error)
		}
	}()

	return ie.scan(rootIfdName, ifdOffset, visitor)
}

// Scan enumerates the different EXIF's IFD blocks.
func (ie *IfdEnumerate) scan(fqIfdName string, ifdOffset uint32, visitor TagVisitorFn) (err error) {
	defer func() {
		if state := recover(); state != nil {
			err = state.(error)
		}
	}()
	for ifdIndex := 0; ; ifdIndex++ {
		//fmt.Printf("Parsing IFD [%s] (%d) at offset (%04x).\n", fqIfdName, ifdIndex, ifdOffset)
		//ifdEnumerateLogger.Debugf(nil, "Parsing IFD [%s] (%d) at offset (%04x).", fqIfdName, ifdIndex, ifdOffset)

		enumerator := ie.getTagEnumerator(ifdOffset)
		nextIfdOffset, _, _, err := ie.ParseIfd2(fqIfdName, ifdIndex, enumerator, visitor, true)
		if err != nil {
			return err
		}

		if nextIfdOffset == 0 {
			break
		}

		ifdOffset = nextIfdOffset
	}

	return nil
}

// ParseIfd decodes the IFD block that we're currently sitting on the first
// byte of.
// WIP: Test & Benchmark
func (ie *IfdEnumerate) ParseIfd(fqIfdPath string, ifdIndex int, enumerator *IfdTagEnumerator, visitor TagVisitorFn, doDescend bool) (nextIfdOffset uint32, entries []*IfdTagEntry, thumbnailData []byte, err error) {
	defer func() {
		if state := recover(); state != nil {
			err = state.(error)
		}
	}()

	tagCount, _, err := enumerator.getUint16()
	if err != nil {
		panic(err)
	}

	//ifdEnumerateLogger.Debugf(nil, "Current IFD tag-count: (%d)", tagCount)

	entries = make([]*IfdTagEntry, 0, 20)

	var enumeratorThumbnailOffset *IfdTagEntry
	var enumeratorThumbnailSize *IfdTagEntry
	//fmt.Println(tagCount)

	for i := 0; i < int(tagCount); i++ {
		ite, err := ie.parseTag(fqIfdPath, i, enumerator)
		if err != nil {
			if errors.Is(err, ErrTagTypeNotValid) {
				// Log TagNotValid Error
				//ifdEnumerateLogger.Warningf(nil, "Tag in IFD [%s] at position (%d) has invalid type and will be skipped.", fqIfdPath, i)
				continue
			}
			if err != nil {
				panic(err)
			}
		}

		tagID := ite.TagID()
		if tagID == ThumbnailOffsetTagId {
			enumeratorThumbnailOffset = ite
			continue
		} else if tagID == ThumbnailSizeTagId {
			enumeratorThumbnailSize = ite
			continue
		}

		if visitor != nil && ite.ChildIfdPath() == "" {
			if err := visitor(fqIfdPath, ifdIndex, ite); err != nil {
				panic(err)
			}
		}

		// If it's an IFD but not a standard one, it'll just be seen as a LONG
		// (the standard IFD tag type), later, unless we skip it because it's
		// [likely] not even in the standard list of known tags.
		if ite.ChildIfdPath() != "" {
			if doDescend == true {
				//ifdEnumerateLogger.Debugf(nil, "Descending to IFD [%s].", ite.ChildIfdPath())
				//fmt.Printf("Descending to IFD [%s].\n", ite.ChildIfdPath())
				if err := ie.scan(ite.ChildFqIfdPath(), ite.getValueOffset(), visitor); err != nil {
					panic(err)
				}
			}
		}

		entries = append(entries, ite)
	}

	// Needs fixing!!!
	if enumeratorThumbnailOffset != nil && enumeratorThumbnailSize != nil {
		//	thumbnailData, err = ie.parseThumbnail(enumeratorThumbnailOffset, enumeratorThumbnailSize)
		//	if err != nil {
		//		panic(err)
		//	}
		//	//log.PanicIf(err)
	}

	// NextIfdOffset
	if nextIfdOffset, _, err = enumerator.getUint32(); err != nil {
		panic(err)
	}

	//ifdEnumerateLogger.Debugf(nil, "Next IFD at offset: (%08x)", nextIfdOffset)

	return nextIfdOffset, entries, thumbnailData, nil
}

// getUint16 reads a uint16 and advances both our current and our current
// accumulator (which allows us to know how far to seek to the beginning of the
// next IFD when it's time to jump).
func (ife *IfdTagEnumerator) getUint16() (value uint16, raw []byte, err error) {
	defer func() {
		if state := recover(); state != nil {
			err = state.(error)
			//err = log.Wrap(state.(error))
		}
	}()

	needBytes := 2
	offset := 0
	raw = make([]byte, needBytes)

	for offset < needBytes {
		n, err := ife.exifReader.Read(raw[offset:])
		//n, err := ife.buffer.Read(raw[offset:])
		if err != nil {
			panic(err)
		}

		offset += n
	}
	value = ife.exifReader.byteOrder.Uint16(raw)
	//value = ife.byteOrder.Uint16(raw)

	return value, raw, nil
}

// getUint32 reads a uint32 and advances both our current and our current
// accumulator (which allows us to know how far to seek to the beginning of the
// next IFD when it's time to jump).
func (ife *IfdTagEnumerator) getUint32() (value uint32, raw []byte, err error) {
	defer func() {
		if state := recover(); state != nil {
			err = state.(error)
		}
	}()

	needBytes := 4
	offset := 0
	raw = make([]byte, needBytes)

	for offset < needBytes {
		n, err := ife.exifReader.Read(raw[offset:])
		if err != nil {
			panic(err)
		}

		offset += n
	}

	value = ife.exifReader.byteOrder.Uint32(raw)

	return value, raw, nil
}

// uint32 reads a uint32 and advances both our current and our current
// accumulator (which allows us to know how far to seek to the beginning of the
// next IFD when it's time to jump).
func (ife *IfdTagEnumerator) uint32() (value uint32, err error) {
	defer func() {
		if state := recover(); state != nil {
			err = state.(error)
		}
	}()
	if len(ife.rawBytes) < 4 {
		return 0, fmt.Errorf("Not enough bytes")
	}

	if n, err := ife.exifReader.Read(ife.rawBytes[:4]); err != nil || n != 4 { // Uint32 = 4bytes
		panic(err)
	}

	value = ife.exifReader.byteOrder.Uint32(ife.rawBytes[:4])

	return value, nil
}

// uint16 reads a uint16 and advances both our current and our current
// accumulator (which allows us to know how far to seek to the beginning of the
// next IFD when it's time to jump).
func (ife *IfdTagEnumerator) uint16() (value uint16, err error) {
	defer func() {
		if state := recover(); state != nil {
			err = state.(error)
		}
	}()
	if len(ife.rawBytes) < 2 {
		return 0, fmt.Errorf("Not enough bytes")
	}

	if n, err := ife.exifReader.Read(ife.rawBytes[:2]); err != nil || n != 2 { // Uint16 = 2bytes
		panic(err)
	}

	value = ife.exifReader.byteOrder.Uint16(ife.rawBytes[:2])

	return value, nil
}

func (ie *IfdEnumerate) parseTag(fqIfdPath string, tagPosition int, enumerator *IfdTagEnumerator) (ite *IfdTagEntry, err error) {
	defer func() {
		if state := recover(); state != nil {
			err = state.(error)
			//err = log.Wrap(state.(error))
		}
	}()

	// TagID
	id, _, err := enumerator.getUint16()
	if err != nil {
		panic(err)
	}
	tagID := exif.TagID(id)

	// TagType
	tagTypeRaw, _, err := enumerator.getUint16()
	if err != nil {
		panic(err)
	}
	tagType := exif.TagType(tagTypeRaw)
	if tagType.IsValid() == false {
		panic(ErrTagTypeNotValid)
	}

	// UnitCount
	unitCount, _, err := enumerator.getUint32()
	if err != nil {
		panic(err)
	}

	// Offsets
	valueOffset, rawValueOffset, err := enumerator.getUint32()
	if err != nil {
		panic(err)
	}

	ifdPath, err := ie.ifdMapping.StripPathPhraseIndices(fqIfdPath)
	if err != nil {
		panic(err)
	}
	//fmt.Println(ifdPath, fqIfdPath)
	//ifdPath := fqIfdPath
	ite = newIfdTagEntry(
		ifdPath,
		tagID,
		tagPosition,
		tagType,
		unitCount,
		valueOffset,
		rawValueOffset,
		ie.exifReader,
		ie.exifReader.byteOrder)

	//fmt.Sprintln(unitCount, valueOffset, rawValueOffset)

	// If it's an IFD but not a standard one, it'll just be seen as a LONG
	// (the standard IFD tag type), later, unless we skip it because it's
	// [likely] not even in the standard list of known tags.
	mi, err := ie.ifdMapping.GetChild(ifdPath, tagID)
	if err == nil {
		ite.SetChildIfd(
			fmt.Sprintf("%s/%s", fqIfdPath, mi.Name),
			mi.PathPhrase(),
			mi.Name)

		// We also need to set `tag.ChildFqIfdPath` but can't do it here
		// because we don't have the IFD index.
	} else if !errors.Is(err, ErrChildIfdNotMapped) {
		panic(err)
	}

	return ite, nil
}

func (ie *IfdEnumerate) parseTag2(fqIfdPath string, tagPosition int, enumerator *IfdTagEnumerator) (ite *IfdTagEntry, err error) {
	defer func() {
		if state := recover(); state != nil {
			err = state.(error)
			//err = log.Wrap(state.(error))
		}
	}()

	// TagID
	id, err := enumerator.uint16()
	if err != nil {
		panic(err)
	}
	tagID := exif.TagID(id)

	// TagType
	tagTypeRaw, err := enumerator.uint16()
	if err != nil {
		panic(err)
	}
	tagType := exif.TagType(tagTypeRaw)
	if tagType.IsValid() == false {
		panic(ErrTagTypeNotValid)
	}

	// UnitCount
	unitCount, err := enumerator.uint32()
	if err != nil {
		panic(err)
	}
	//if n, err := enumerator.exifReader.Read(raw); err != nil || n != 4 { // Uint32 = 4bytes
	//	panic(err)
	//}
	//unitCount := enumerator.exifReader.byteOrder.Uint32(raw)

	// Offsets
	valueOffset, err := enumerator.uint32()
	if err != nil {
		panic(err)
	}
	rawValueOffset := make([]byte, 4)
	if n := copy(rawValueOffset, enumerator.rawBytes); n < 4 {
		panic(err)
	}
	//rawValueOffset := raw

	//ifdPath, err := ie.ifdMapping.StripPathPhraseIndices(fqIfdPath)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(ifdPath, fqIfdPath)
	ifdPath := fqIfdPath
	ite = newIfdTagEntry(
		ifdPath,
		tagID,
		tagPosition,
		tagType,
		unitCount,
		valueOffset,
		rawValueOffset,
		ie.exifReader,
		ie.exifReader.byteOrder)

	//fmt.Sprintln(unitCount, valueOffset, rawValueOffset)

	// If it's an IFD but not a standard one, it'll just be seen as a LONG
	// (the standard IFD tag type), later, unless we skip it because it's
	// [likely] not even in the standard list of known tags.
	mi, err := ie.ifdMapping.GetChild(ifdPath, tagID)
	if err == nil {
		ite.SetChildIfd(
			fmt.Sprintf("%s/%s", fqIfdPath, mi.Name),
			mi.PathPhrase(),
			mi.Name)

		// We also need to set `tag.ChildFqIfdPath` but can't do it here
		// because we don't have the IFD index.
	} else if !errors.Is(err, ErrChildIfdNotMapped) {
		panic(err)
	}

	return ite, nil
}
