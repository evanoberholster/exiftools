// Package xmp aliases the package foung in (trimmer.io/go-xmp/xmp)
package xmp

import (
	"io"

	"github.com/evanoberholster/exiftools/models"
	"trimmer.io/go-xmp/models/dc"
	xmpbase "trimmer.io/go-xmp/models/xmp_base"
	"trimmer.io/go-xmp/xmp"
)

// ReadXMPDocument - Read from file to XMP Document
func ReadXMPDocument(r io.Reader) (*xmp.Document, error) {
	// ScanPackets for XMP
	bb, err := xmp.ScanPackets(r)
	if (err != nil && err != io.EOF) || len(bb) == 0 {
		return nil, err
	}

	// Initialize new XMP Document
	doc := &xmp.Document{}

	// Unmarshal XMP
	if err := xmp.Unmarshal(bb[0], doc); err != nil {
		return nil, err
	}

	return doc, nil
}

// Unmarshal - Unmarshal XMP Document
func Unmarshal(bb []byte) (*xmp.Document, error) {
	doc := &xmp.Document{}
	err := xmp.Unmarshal(bb, doc)
	return doc, err
}

// Base - Extract XmpBase from XMP Document
func Base(m *xmp.Document) models.XmpBase {
	c := xmpbase.FindModel(m)
	if c == nil {
		return models.XmpBase{}
	}

	return models.XmpBase{
		CreateDate:   c.CreateDate.Value(),
		MetadataDate: c.MetadataDate.Value(),
		ModifyDate:   c.ModifyDate.Value(),
		Label:        string(c.Label),
		Rating:       int(c.Rating),
		CreatorTool:  c.CreatorTool.String(),
	}
}

// DublinCore - Extract DublinCore from XMP Document
func DublinCore(m *xmp.Document) models.DublinCore {
	c := dc.FindModel(m)
	if c == nil {
		return models.DublinCore{}
	}

	creator := []string(c.Creator)
	if creator == nil {
		creator = []string{}
	}

	subject := []string(c.Subject)
	if subject == nil {
		subject = []string{}
	}

	return models.DublinCore{
		Creator:     creator,
		Description: c.Description.Default(),
		Format:      string(c.Format),
		Rights:      c.Rights.Default(),
		Source:      "",
		Subject:     subject,
		Title:       c.Title.Default(),
	}
}
