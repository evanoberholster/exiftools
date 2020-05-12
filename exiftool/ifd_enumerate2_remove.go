package exiftool

// NewIfdTagEnumerator creates a new IFD Tag Enumerator
//func NewIfdTagEnumerator2(reader *ExifReader, byteOrder binary.ByteOrder, ifdOffset uint32) (enumerator *IfdTagEnumerator) {
//	enumerator = &IfdTagEnumerator{
//		exifReader: reader.SubReader(int64(ifdOffset)),
//		byteOrder:  byteOrder,
//		buffer:     new(bytes.Buffer),
//		//buffer:    bytes.NewBuffer(addressableData[ifdOffset:]),
//	}
//
//	return enumerator
//}

//func (ie *IfdEnumerate) getTagEnumerator2(ifdOffset uint32) (enumerator *IfdTagEnumerator) {
//	enumerator = NewIfdTagEnumerator2(
//		ie.exifReader,
//		ie.exifReader.byteOrder,
//		ifdOffset)
//
//	return enumerator
//}

// ParseIfd decodes the IFD block that we're currently sitting on the first
// byte of.
