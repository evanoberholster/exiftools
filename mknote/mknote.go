// Package mknote provides makernote parsers that can be used with goexif/exif.
package mknote

import (
	"bytes"
	"fmt"

	"github.com/evanoberholster/exiftools/exif"
	"github.com/evanoberholster/exiftools/tiff"
)

var (
	// Canon is an exif.Parser for canon makernote data.
	Canon = &canon{}
	// NikonV3 is an exif.Parser for nikon makernote data.
	NikonV3 = &nikonV3{}
	// AdobeDNG is an exif.Parse for DNG subIfds.
	AdobeDNG = &adobeDNG{}
	// All is a list of all available makernote parsers
	All = []exif.Parser{Canon, NikonV3, AdobeDNG}
)

type adobeDNG struct{}

// Parse decodes all Adobe DNG subIFDS found in x and adds it to x.
func (_ *adobeDNG) Parse(x *exif.Exif) error {
	m, err := x.Get(exif.SubIfdsPointer)
	if err != nil {
		return nil
	}
	if !(m.Count > 0) {
		return nil
	}
	subIfds := []map[uint16]exif.FieldName{
		SubIFD0Fields,
		SubIFD1Fields,
		SubIFD2Fields,
	}
	r := bytes.NewReader(x.Raw)
	for i, sub := range subIfds {
		offset, err := m.Int64(i)
		if err != nil {
			return nil
		}
		_, err = r.Seek(offset, 0)
		if err != nil {
			return fmt.Errorf("exif: seek to sub-IFD %s failed: %v", exif.SubIfdsPointer, err)
		}
		subDir, _, err := tiff.DecodeDir(r, x.Tiff.Order)
		if err != nil {
			return fmt.Errorf("exif: sub-IFD %s decode failed: %v", exif.SubIfdsPointer, err)
		}
		x.LoadTags(subDir, sub, false)
	}

	return nil
}

type canon struct{}

// Parse decodes all Canon makernote data found in x and adds it to x.
func (_ *canon) Parse(x *exif.Exif) error {
	m, err := x.Get(exif.MakerNote)
	if err != nil {
		return nil
	}

	// Confirm that exif.Make is Canon
	if mk, err := x.Get(exif.Make); err != nil {
		if val, err := mk.StringVal(); err != nil || val != "Canon" {
			return nil
		}
	}

	// Canon notes are a single IFD directory with no header.
	// Reader offsets need to be w.r.t. the original tiff structure.
	cReader := bytes.NewReader(append(make([]byte, m.ValOffset), m.Val...))
	cReader.Seek(int64(m.ValOffset), 0)

	mkNotesDir, _, err := tiff.DecodeDir(cReader, x.Tiff.Order)
	if err != nil {
		return err
	}
	// Parse Canon MakerFields
	x.LoadTags(mkNotesDir, makerNoteCanonFields, false)

	return nil
}

type nikonV3 struct{}

// Parse decodes all Nikon makernote data found in x and adds it to x.
func (_ *nikonV3) Parse(x *exif.Exif) error {
	m, err := x.Get(exif.MakerNote)
	if err != nil {
		return nil
	}
	if len(m.Val) < 6 {
		return nil
	}
	if bytes.Compare(m.Val[:6], []byte("Nikon\000")) != 0 {
		return nil
	}

	// Nikon v3 maker note is a self-contained IFD (offsets are relative
	// to the start of the maker note)
	nReader := bytes.NewReader(m.Val[10:])
	mkNotes, err := tiff.Decode(nReader)
	if err != nil {
		return err
	}
	makerNoteOffset := m.ValOffset + 10
	x.LoadTags(mkNotes.Dirs[0], makerNoteNikon3Fields, false)

	if err := loadSubDir(x, nReader, NikonPreviewPtr, makerNoteNikon3PreviewFields); err != nil {
	}
	previewTag, err := x.Get(NikonPreviewImageStart)
	if err == nil {
		offset, _ := previewTag.Int64(0)
		previewTag.SetInt(0, offset+int64(makerNoteOffset))
		x.Update(NikonPreviewImageStart, previewTag)
	}

	return nil
}

func loadSubDir(x *exif.Exif, r *bytes.Reader, ptr exif.FieldName, fieldMap map[uint16]exif.FieldName) error {
	tag, err := x.Get(ptr)
	if err != nil {
		return nil
	}
	offset, err := tag.Int64(0)
	if err != nil {
		return nil
	}

	_, err = r.Seek(offset, 0)
	if err != nil {
		return fmt.Errorf("exif: seek to sub-IFD %s failed: %v", ptr, err)
	}
	subDir, _, err := tiff.DecodeDir(r, x.Tiff.Order)
	if err != nil {
		return fmt.Errorf("exif: sub-IFD %s decode failed: %v", ptr, err)
	}
	x.LoadTags(subDir, fieldMap, false)
	return nil
}
