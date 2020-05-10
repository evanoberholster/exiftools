package exif

import (
	"bytes"
	"encoding/binary"
	"errors"
)

var (
	// ErrNotEnoughData is used when there isn't enough data to accomodate what
	// we're trying to parse (sizeof(type) * unit_count).
	ErrNotEnoughData = errors.New("not enough data for type")

	// ErrUnhandledUndefinedTypedTag is used when we try to parse a tag that's
	// recorded as an "unknown" type but not a documented tag (therefore
	// leaving us not knowning how to read it).
	ErrUnhandledUndefinedTypedTag = errors.New("not a standard unknown-typed tag")

	// ErrUnparseableValue is the error for a value that we should have been
	// able to parse but were not able to.
	ErrUnparseableValue = errors.New("unparseable undefined tag")
)

var (
	parser *Parser
)

func init() {
	parser = new(Parser)
}

// Parser knows how to parse all well-defined, encoded EXIF types.
type Parser struct{}

// TODO: Add Test & Benchmark
func (p *Parser) ParseBytes(data []byte, unitCount uint32) (value []uint8, err error) {
	defer func() {
		if state := recover(); state != nil {
			err = state.(error)
		}
	}()

	count := int(unitCount)

	if len(data) < (TypeByte.Size() * count) {
		err = ErrNotEnoughData
		return
	}

	value = []uint8(data[:count])

	return value, nil
}

// ParseASCII returns a string and auto-strips the trailing NUL character that
// should be at the end of the encoding.
// TODO: Add Test & Benchmark
func (p *Parser) ParseASCII(data []byte, unitCount uint32) (value string, err error) {
	defer func() {
		if state := recover(); state != nil {
			err = state.(error)
		}
	}()

	count := int(unitCount)

	if len(data) < (TypeASCII.Size() * count) {
		err = ErrNotEnoughData
		return
	}

	if len(data) == 0 || data[count-1] != 0 {
		s := string(data[:count])
		//parserLogger.Warningf(nil, "ascii not terminated with nul as expected: [%v]", s)
		return s, nil
	}

	// Auto-strip the NUL from the end. It serves no purpose outside of
	// encoding semantics.
	return string(data[:count-1]), nil
}

// ParseASCIINoNul returns a string without any consideration for a trailing NUL
// character.
// TODO: Add Test & Benchmark
func (p *Parser) ParseASCIINoNul(data []byte, unitCount uint32) (value string, err error) {
	defer func() {
		if state := recover(); state != nil {
			err = state.(error)
		}
	}()
	count := int(unitCount)

	if len(data) < (TypeASCII.Size() * count) {
		err = ErrNotEnoughData
		return
	}

	return string(data[:count]), nil
}

// ParseShorts knows how to parse an encoded list of shorts.
// TODO: Add Test & Benchmark
func (p *Parser) ParseShorts(data []byte, unitCount uint32, byteOrder binary.ByteOrder) (value []uint16, err error) {
	defer func() {
		if state := recover(); state != nil {
			err = state.(error)
		}
	}()
	count := int(unitCount)

	if len(data) < (TypeShort.Size() * count) {
		panic(ErrNotEnoughData)
	}

	value = make([]uint16, count)
	for i := 0; i < count; i++ {
		value[i] = byteOrder.Uint16(data[i*2:])
	}

	return value, nil
}

// ParseLongs knows how to encode an encoded list of unsigned longs.
// TODO: Add Test & Benchmark
func (p *Parser) ParseLongs(data []byte, unitCount uint32, byteOrder binary.ByteOrder) (value []uint32, err error) {
	defer func() {
		if state := recover(); state != nil {
			err = state.(error)
		}
	}()

	count := int(unitCount)

	if len(data) < (TypeLong.Size() * count) {
		panic(ErrNotEnoughData)
	}

	value = make([]uint32, count)
	for i := 0; i < count; i++ {
		value[i] = byteOrder.Uint32(data[i*4:])
	}

	return value, nil
}

// ParseRationals knows how to parse an encoded list of unsigned rationals.
// TODO: Add Test & Benchmark
func (p *Parser) ParseRationals(data []byte, unitCount uint32, byteOrder binary.ByteOrder) (value []Rational, err error) {
	defer func() {
		if state := recover(); state != nil {
			err = state.(error)
		}
	}()

	count := int(unitCount)

	if len(data) < (TypeRational.Size() * count) {
		panic(ErrNotEnoughData)
	}

	value = make([]Rational, count)
	for i := 0; i < count; i++ {
		value[i].Numerator = byteOrder.Uint32(data[i*8:])
		value[i].Denominator = byteOrder.Uint32(data[i*8+4:])
	}

	return value, nil
}

// ParseSignedLongs knows how to parse an encoded list of signed longs.
// TODO: Add Test & Benchmark
func (p *Parser) ParseSignedLongs(data []byte, unitCount uint32, byteOrder binary.ByteOrder) (value []int32, err error) {
	defer func() {
		if state := recover(); state != nil {
			err = state.(error)
		}
	}()

	count := int(unitCount)

	if len(data) < (TypeSignedLong.Size() * count) {
		panic(ErrNotEnoughData)
	}

	b := bytes.NewBuffer(data)

	value = make([]int32, count)
	for i := 0; i < count; i++ {
		if err = binary.Read(b, byteOrder, &value[i]); err != nil {
			return
		}
	}

	return value, nil
}

// ParseSignedRationals knows how to parse an encoded list of signed
// rationals.
// TODO: Add Test & Benchmark
func (p *Parser) ParseSignedRationals(data []byte, unitCount uint32, byteOrder binary.ByteOrder) (value []SignedRational, err error) {
	defer func() {
		if state := recover(); state != nil {
			err = state.(error)
		}
	}()
	count := int(unitCount)

	if len(data) < (TypeSignedRational.Size() * count) {
		panic(ErrNotEnoughData)
	}

	b := bytes.NewBuffer(data)

	value = make([]SignedRational, count)
	for i := 0; i < count; i++ {
		if err = binary.Read(b, byteOrder, &value[i].Numerator); err != nil {
			return
		}

		if err = binary.Read(b, byteOrder, &value[i].Denominator); err != nil {
			return
		}
	}

	return value, nil
}
