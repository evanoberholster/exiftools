package exiftool

import "errors"

// ParseIfd2 decodes the IFD block that we're currently sitting on the first
// byte of.
// WIP: Test & Benchmark
func (ie *IfdEnumerate) ParseIfd2(fqIfdPath string, ifdIndex int, enumerator *IfdTagEnumerator, visitor TagVisitorFn, doDescend bool) (nextIfdOffset uint32, entries []*IfdTagEntry, thumbnailData []byte, err error) {
	defer func() {
		if state := recover(); state != nil {
			err = state.(error)
		}
	}()

	//rawBytes := make([]byte, 4)

	tagCount, err := enumerator.uint16()
	//tagCount, _, err := enumerator.getUint16()
	if err != nil {
		panic(err)
	}

	//ifdEnumerateLogger.Debugf(nil, "Current IFD tag-count: (%d)", tagCount)

	//entries = make([]*IfdTagEntry, 0, 20)

	//var enumeratorThumbnailOffset *IfdTagEntry
	//var enumeratorThumbnailSize *IfdTagEntry
	//fmt.Println(tagCount)

	for i := 0; i < int(tagCount); i++ {
		ite, err := ie.parseTag2(fqIfdPath, i, enumerator)
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

		//tagID := ite.TagID()
		//if tagID == ThumbnailOffsetTagId {
		//	enumeratorThumbnailOffset = ite
		//	continue
		//} else if tagID == ThumbnailSizeTagId {
		//	enumeratorThumbnailSize = ite
		//	continue
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
			if doDescend == true {
				//ifdEnumerateLogger.Debugf(nil, "Descending to IFD [%s].", ite.ChildIfdPath())
				//fmt.Printf("Descending to IFD [%s].\n", ite.ChildIfdPath())
				if err := ie.scan(ite.ChildFqIfdPath(), ite.getValueOffset(), visitor); err != nil {
					panic(err)
				}
			}
		}

		//entries = append(entries, ite)
	}

	// Needs fixing!!!
	//if enumeratorThumbnailOffset != nil && enumeratorThumbnailSize != nil {
	//	thumbnailData, err = ie.parseThumbnail(enumeratorThumbnailOffset, enumeratorThumbnailSize)
	//	if err != nil {
	//		panic(err)
	//	}
	//	//log.PanicIf(err)
	//}

	// NextIfdOffset
	if nextIfdOffset, err = enumerator.uint32(); err != nil {
		panic(err)
	}

	//ifdEnumerateLogger.Debugf(nil, "Next IFD at offset: (%08x)", nextIfdOffset)
	return nextIfdOffset, entries, thumbnailData, nil
}
