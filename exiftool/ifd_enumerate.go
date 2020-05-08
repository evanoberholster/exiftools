package exiftool

import (
	"bytes"
	"encoding/binary"
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

func NewIfdEnumerate(ifdMapping *IfdMapping, tagIndex *TagIndex, exifData []byte, byteOrder binary.ByteOrder) *IfdEnumerate {
	return &IfdEnumerate{
		exifData:   exifData,
		buffer:     bytes.NewBuffer(exifData),
		byteOrder:  byteOrder,
		ifdMapping: ifdMapping,
		tagIndex:   tagIndex,
	}
}

// NewIfdTagEnumerator creates a new IFD Tag Enumerator
func NewIfdTagEnumerator(addressableData []byte, byteOrder binary.ByteOrder, ifdOffset uint32) (enumerator *IfdTagEnumerator) {
	enumerator = &IfdTagEnumerator{
		addressableData: addressableData,
		byteOrder:       byteOrder,
		buffer:          bytes.NewBuffer(addressableData[ifdOffset:]),
	}

	return enumerator
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
		}
	}()
	for ifdIndex := 0; ; ifdIndex++ {
		//ifdEnumerateLogger.Debugf(nil, "Parsing IFD [%s] (%d) at offset (%04x).", fqIfdName, ifdIndex, ifdOffset)
		enumerator := ie.getTagEnumerator(ifdOffset)
		nextIfdOffset, _, _, err := ie.ParseIfd(fqIfdName, ifdIndex, enumerator, visitor, true)
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

func (ie *IfdEnumerate) getTagEnumerator(ifdOffset uint32) (enumerator *IfdTagEnumerator) {
	enumerator = NewIfdTagEnumerator(
		ie.exifData[ExifAddressableAreaStart:],
		ie.byteOrder,
		ifdOffset)

	return enumerator
}

// ParseIfd decodes the IFD block that we're currently sitting on the first
// byte of.
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

	entries = make([]*IfdTagEntry, 0)

	var enumeratorThumbnailOffset *IfdTagEntry
	var enumeratorThumbnailSize *IfdTagEntry

	for i := 0; i < int(tagCount); i++ {
		ite, err := ie.parseTag(fqIfdPath, i, enumerator)
		if err != nil {
			if err == ErrTagTypeNotValid {
				//if log.Is(err, ErrTagTypeNotValid) == true {
				//ifdEnumerateLogger.Warningf(nil, "Tag in IFD [%s] at position (%d) has invalid type and will be skipped.", fqIfdPath, i)
				continue
			}
			if err != nil {
				panic(err)
			}
			//log.Panic(err)
		}

		tagID := ite.TagID()
		if tagID == ThumbnailOffsetTagId {
			enumeratorThumbnailOffset = ite

			continue
		} else if tagID == ThumbnailSizeTagId {
			enumeratorThumbnailSize = ite
			continue
		}

		if visitor != nil {
			if err := visitor(fqIfdPath, ifdIndex, ite); err != nil {
				panic(err)
			}
			//log.PanicIf(err)
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

	nextIfdOffset, _, err = enumerator.getUint32()
	if err != nil {
		panic(err)
	}
	//log.PanicIf(err)

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
		n, err := ife.buffer.Read(raw[offset:])
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
func (ife *IfdTagEnumerator) getUint32() (value uint32, raw []byte, err error) {
	defer func() {
		if state := recover(); state != nil {
			err = state.(error)
			//err = log.Wrap(state.(error))
		}
	}()

	needBytes := 4
	offset := 0
	raw = make([]byte, needBytes)

	for offset < needBytes {
		n, err := ife.buffer.Read(raw[offset:])
		if err != nil {
			panic(err)
		}

		offset += n
	}

	value = ife.byteOrder.Uint32(raw)

	return value, raw, nil
}

func (ie *IfdEnumerate) parseTag(fqIfdPath string, tagPosition int, enumerator *IfdTagEnumerator) (ite *IfdTagEntry, err error) {
	defer func() {
		if state := recover(); state != nil {
			err = state.(error)
			//err = log.Wrap(state.(error))
		}
	}()

	id, _, err := enumerator.getUint16()
	if err != nil {
		panic(err)
	}
	tagID := exif.TagID(id)
	//log.PanicIf(err)

	tagTypeRaw, _, err := enumerator.getUint16()
	if err != nil {
		panic(err)
	}
	tagType := exif.TagType(tagTypeRaw)

	unitCount, _, err := enumerator.getUint32()
	if err != nil {
		panic(err)
	}

	valueOffset, rawValueOffset, err := enumerator.getUint32()
	if err != nil {
		panic(err)
	}

	if tagType.IsValid() == false {
		panic(ErrTagTypeNotValid)
	}

	ifdPath, err := ie.ifdMapping.StripPathPhraseIndices(fqIfdPath)
	if err != nil {
		panic(err)
	}
	//log.PanicIf(err)

	ite = newIfdTagEntry(
		ifdPath,
		tagID,
		tagPosition,
		tagType,
		unitCount,
		valueOffset,
		rawValueOffset,
		ie.exifData[ExifAddressableAreaStart:],
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
