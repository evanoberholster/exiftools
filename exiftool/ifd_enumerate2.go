package exiftool

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/evanoberholster/exiftools/exiftool/exif"
)

// IfdEnumerate -
type IfdEnumerate struct {
	exifReader    *ExifReader // New
	exifData      []byte
	buffer        *bytes.Buffer
	byteOrder     binary.ByteOrder
	currentOffset uint32
	tagIndex      *TagIndex
	ifdMapping    *IfdMapping
}

func newIfdEnumerate(er *ExifReader, ifdMapping *IfdMapping, tagIndex *TagIndex, byteOrder binary.ByteOrder) *IfdEnumerate {
	return &IfdEnumerate{
		exifReader: er,
		buffer:     new(bytes.Buffer),
		//bufReader:  bufio.NewReader(reader),
		byteOrder:  byteOrder,
		ifdMapping: ifdMapping,
		tagIndex:   tagIndex,
	}
}

// Scan enumerates the different EXIF's IFD blocks.
func (ie *IfdEnumerate) scan2(fqIfdName string, ifdOffset uint32, visitor TagVisitorFn) (err error) {
	defer func() {
		if state := recover(); state != nil {
		}
	}()
	//fmt.Println("Offset", ifdOffset)
	for ifdIndex := 0; ; ifdIndex++ {
		//ifdEnumerateLogger.Debugf(nil, "Parsing IFD [%s] (%d) at offset (%04x).", fqIfdName, ifdIndex, ifdOffset)
		enumerator := ie.getTagEnumerator2(ifdOffset)
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

// IfdTagEnumerator knows how to decode an IFD and all of the tags it
// describes.
//
// The IFDs and the actual values can float throughout the EXIF block, but the
// IFD itself is just a minor header followed by a set of repeating,
// statically-sized records. So, the tags (though notnecessarily their values)
// are fairly simple to enumerate.
type IfdTagEnumerator struct {
	exifReader      *ExifReader
	byteOrder       binary.ByteOrder
	addressableData []byte
	ifdOffset       uint32
	buffer          *bytes.Buffer
}

// NewIfdTagEnumerator creates a new IFD Tag Enumerator
func NewIfdTagEnumerator2(reader *ExifReader, byteOrder binary.ByteOrder, ifdOffset uint32) (enumerator *IfdTagEnumerator) {
	enumerator = &IfdTagEnumerator{
		exifReader: reader.NewReader(int64(ifdOffset)),
		byteOrder:  byteOrder,
		buffer:     new(bytes.Buffer),
		//buffer:    bytes.NewBuffer(addressableData[ifdOffset:]),
	}

	return enumerator
}

func (ie *IfdEnumerate) getTagEnumerator2(ifdOffset uint32) (enumerator *IfdTagEnumerator) {
	enumerator = NewIfdTagEnumerator2(
		ie.exifReader,
		ie.byteOrder,
		ifdOffset)

	return enumerator
}

// ParseIfd decodes the IFD block that we're currently sitting on the first
// byte of.
func (ie *IfdEnumerate) ParseIfd2(fqIfdPath string, ifdIndex int, enumerator *IfdTagEnumerator, visitor TagVisitorFn, doDescend bool) (nextIfdOffset uint32, entries []*IfdTagEntry, thumbnailData []byte, err error) {
	defer func() {
		if state := recover(); state != nil {
			err = state.(error)
		}
	}()

	tagCount, _, err := enumerator.get2Uint16()
	if err != nil {
		panic(err)
	}
	//fmt.Println("Tag Count", tagCount)
	//ifdEnumerateLogger.Debugf(nil, "Current IFD tag-count: (%d)", tagCount)

	entries = make([]*IfdTagEntry, 0)

	var enumeratorThumbnailOffset *IfdTagEntry
	var enumeratorThumbnailSize *IfdTagEntry

	for i := 0; i < int(tagCount); i++ {
		ite, err := ie.parseTag2(fqIfdPath, i, enumerator)
		if err != nil {
			if err == ErrTagTypeNotValid {
				//if log.Is(err, ErrTagTypeNotValid) == true {
				//ifdEnumerateLogger.Warningf(nil, "Tag in IFD [%s] at position (%d) has invalid type and will be skipped.", fqIfdPath, i)
				//fmt.Println("Error Not Valid")
				continue
			}
			if err != nil {
				fmt.Println("Error ehre")
				panic(err)
			}
		}

		tagID := ite.TagID()
		if tagID == ThumbnailOffsetTagId {
			enumeratorThumbnailOffset = ite
			//fmt.Println("Thumbnail Offset")
			//continue
		} else if tagID == ThumbnailSizeTagId {
			enumeratorThumbnailSize = ite
			//fmt.Println("Thumbnail Size")
			//continue
		}

		// LoadMakernotes accoring to Make
		//if ite.TagID() == ifd.Make && fqIfdPath == ifd.IfdRoot {
		//	if value, err := ite.Value(); err == nil {
		//		ie.ifdMapping.LoadIfds(mknote.LoadMakernotesIfd(value.(string)))
		//	}
		//}

		if visitor != nil && ite.ChildIfdPath() == "" {
			if err := visitor(fqIfdPath, ifdIndex, ite); err != nil {
				panic(err)
			}
		}

		// If it's an IFD but not a standard one, it'll just be seen as a LONG
		// (the standard IFD tag type), later, unless we skip it because it's
		// [likely] not even in the standard list of known tags.
		if ite.ChildIfdPath() != "" {
			//fmt.Println(ite.childIfdPath)
			if doDescend == true {
				//fmt.Println(ite.getValueOffset())
				//ifdEnumerateLogger.Debugf(nil, "Descending to IFD [%s].", ite.ChildIfdPath())
				//fmt.Printf("Descending to IFD [%s].\n", ite.ChildIfdPath())
				if err := ie.scan2(ite.ChildFqIfdPath(), ite.getValueOffset(), visitor); err != nil {
					panic(err)
				}
				//log.PanicIf(err)
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

	nextIfdOffset, _, err = enumerator.get2Uint32()
	if err != nil {
		panic(err)
	}
	//fmt.Println(nextIfdOffset, enumerator.ifdOffset)
	//panic(fmt.Errorf("Here"))

	//ifdEnumerateLogger.Debugf(nil, "Next IFD at offset: (%08x)", nextIfdOffset)

	return nextIfdOffset, entries, thumbnailData, nil
}

func (ie *IfdEnumerate) parseTag2(fqIfdPath string, tagPosition int, enumerator *IfdTagEnumerator) (ite *IfdTagEntry, err error) {
	defer func() {
		if state := recover(); state != nil {
			err = state.(error)
		}
	}()

	id, _, err := enumerator.get2Uint16()
	if err != nil {
		panic(err)
	}
	tagID := exif.TagID(id)

	tagTypeRaw, _, err := enumerator.get2Uint16()
	if err != nil {
		panic(err)
	}
	tagType := exif.TagType(tagTypeRaw)

	unitCount, _, err := enumerator.get2Uint32()
	if err != nil {
		panic(err)
	}

	valueOffset, rawValueOffset, err := enumerator.get2Uint32()
	if err != nil {
		panic(err)
	}
	//fmt.Println(valueOffset, rawValueOffset)

	if tagType.IsValid() == false {
		panic(ErrTagTypeNotValid)
	}

	ifdPath, err := ie.ifdMapping.StripPathPhraseIndices(fqIfdPath)
	if err != nil {
		panic(err)
	}

	ite = newIfdTagEntry(
		ifdPath,
		tagID,
		tagPosition,
		tagType,
		unitCount,
		valueOffset,
		rawValueOffset,
		ie.exifReader,
		ie.byteOrder)

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
	} else if err != ErrChildIfdNotMapped { //log.Is(err, ErrChildIfdNotMapped) == false {
		panic(err)
	}

	return ite, nil
}

// getUint16 reads a uint16 and advances both our current and our current
// accumulator (which allows us to know how far to seek to the beginning of the
// next IFD when it's time to jump).
func (ife *IfdTagEnumerator) get2Uint16() (value uint16, raw []byte, err error) {
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

	value = ife.byteOrder.Uint16(raw)

	return value, raw, nil
}

// getUint32 reads a uint32 and advances both our current and our current
// accumulator (which allows us to know how far to seek to the beginning of the
// next IFD when it's time to jump).
func (ife *IfdTagEnumerator) get2Uint32() (value uint32, raw []byte, err error) {
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
		//n, err := ife.buffer.Read(raw[offset:])
		if err != nil {
			panic(err)
		}

		offset += n
	}

	value = ife.byteOrder.Uint32(raw)

	return value, raw, nil
}
