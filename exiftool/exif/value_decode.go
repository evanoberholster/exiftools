package exif

// isEmbedded returns whether the value is embedded or a reference. This can't
// be precalculated since the size is not defined for all types (namely the
// "undefined" types).
func (vc *ValueContext) isEmbedded() bool {
	tagType := vc.effectiveValueType()

	return (tagType.Size() * int(vc.unitCount)) <= 4
}

// readRawEncoded returns the encoded bytes for the value that we represent.
//func (vc *ValueContext) readRawEncoded() (rawBytes []byte, err error) {
//	defer func() {
//		if state := recover(); state != nil {
//			err = state.(error)
//		}
//	}()
//
//	tagType := vc.effectiveValueType()
//
//	unitSizeRaw := uint32(tagType.Size())
//
//	if vc.isEmbedded() {
//		byteLength := unitSizeRaw * vc.unitCount
//		return vc.rawValueOffset[:byteLength], nil
//	}
//	//fmt.Println(vc.valueOffset, vc.valueOffset+vc.unitCount*unitSizeRaw)
//	return vc.addressableData[vc.valueOffset : vc.valueOffset+vc.unitCount*unitSizeRaw], nil
//}

// readRawEncoded returns the encoded bytes for the value that we represent.
func (vc *ValueContext) readRawEncoded() (rawBytes []byte, err error) {
	defer func() {
		if state := recover(); state != nil {
			err = state.(error)
		}
	}()

	tagType := vc.effectiveValueType()

	unitSizeRaw := uint32(tagType.Size())

	//fmt.Println(vc.valueOffset, vc.unitCount*unitSizeRaw)
	if vc.isEmbedded() {
		byteLength := unitSizeRaw * vc.unitCount
		return vc.rawValueOffset[:byteLength], nil
	}

	data := make([]byte, vc.unitCount*unitSizeRaw)
	_, err = vc.exifReader.ReadAt(data, int64(vc.valueOffset))
	if err != nil {
		panic(err)
	}
	//fmt.Println(vc.valueOffset, vc.valueOffset+vc.unitCount*unitSizeRaw, n)
	//return vc.addressableData[vc.valueOffset : vc.valueOffset+vc.unitCount*unitSizeRaw], nil
	return data, nil
}

// ReadBytes parses the encoded byte-array from the value-context.
func (vc *ValueContext) ReadBytes() (value []byte, err error) {
	rawValue, err := vc.readRawEncoded()
	if err != nil {
		return
	}

	return parser.ParseBytes(rawValue, vc.unitCount)
}

// ReadASCII parses the encoded NUL-terminated ASCII string from the value-
// context.
func (vc *ValueContext) ReadASCII() (value string, err error) {
	rawValue, err := vc.readRawEncoded()
	if err != nil {
		return
	}

	return parser.ParseASCII(rawValue, vc.unitCount)
}

// ReadASCIINoNul parses the non-NUL-terminated encoded ASCII string from the
// value-context.
func (vc *ValueContext) ReadASCIINoNul() (value string, err error) {
	rawValue, err := vc.readRawEncoded()
	if err != nil {
		return
	}

	return parser.ParseASCIINoNul(rawValue, vc.unitCount)
}

// ReadShorts parses the list of encoded shorts from the value-context.
func (vc *ValueContext) ReadShorts() (value []uint16, err error) {
	rawValue, err := vc.readRawEncoded()
	if err != nil {
		return
	}

	return parser.ParseShorts(rawValue, vc.unitCount, vc.byteOrder)
}

// ReadLongs parses the list of encoded, unsigned longs from the value-context.
func (vc *ValueContext) ReadLongs() (value []uint32, err error) {
	rawValue, err := vc.readRawEncoded()
	if err != nil {
		return
	}

	return parser.ParseLongs(rawValue, vc.unitCount, vc.byteOrder)
}

// ReadRationals parses the list of encoded, unsigned rationals from the value-
// context.
func (vc *ValueContext) ReadRationals() (value []Rational, err error) {
	rawValue, err := vc.readRawEncoded()
	if err != nil {
		return
	}

	return parser.ParseRationals(rawValue, vc.unitCount, vc.byteOrder)
}

// ReadSignedLongs parses the list of encoded, signed longs from the value-context.
func (vc *ValueContext) ReadSignedLongs() (value []int32, err error) {
	rawValue, err := vc.readRawEncoded()
	if err != nil {
		return
	}

	return parser.ParseSignedLongs(rawValue, vc.unitCount, vc.byteOrder)
}

// ReadSignedRationals parses the list of encoded, signed rationals from the
// value-context.
func (vc *ValueContext) ReadSignedRationals() (value []SignedRational, err error) {
	rawValue, err := vc.readRawEncoded()
	if err != nil {
		return
	}

	return parser.ParseSignedRationals(rawValue, vc.unitCount, vc.byteOrder)
}
