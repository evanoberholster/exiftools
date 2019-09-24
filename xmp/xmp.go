package xmp

import (
	"io"
	"os"

	"github.com/evanoberholster/exif/models"
	"trimmer.io/go-xmp/models/dc"
	xmpbase "trimmer.io/go-xmp/models/xmp_base"
	"trimmer.io/go-xmp/xmp"
)

// ReadXMPDocument - Read from file to XMP Document
func ReadXMPDocument(f *os.File) (*xmp.Document, error) {
	f.Seek(0, 0)

	doc := &xmp.Document{}

	bb, err := xmp.ScanPackets(f)
	if (err != nil && err != io.EOF) || len(bb) == 0 {
		return doc, err
	}

	if err := xmp.Unmarshal(bb[0], doc); err != nil {
		return doc, err
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
	var b models.XmpBase
	c := xmpbase.FindModel(m)
	if c == nil {
		return models.XmpBase{}
	}

	b.CreateDate = c.CreateDate.Value()
	b.MetadataDate = c.MetadataDate.Value()
	b.ModifyDate = c.ModifyDate.Value()
	b.Label = string(c.Label)
	b.Rating = int(c.Rating)
	b.CreatorTool = c.CreatorTool.String()

	return b
}

// DublinCore - Extract DublinCore from XMP Document
func DublinCore(m *xmp.Document) models.DublinCore {
	var d models.DublinCore
	c := dc.FindModel(m)
	if c == nil {
		return d
	}

	creator := []string(c.Creator)
	if len(creator) > 0 {
		d.Creator = creator[0]
	}
	s := []string(c.Subject)
	if s == nil {
		s = []string{}
	}
	d.Subject = s
	d.Description = c.Description.Default()
	d.Format = string(c.Format)
	d.Rights = c.Rights.Default()
	d.Title = c.Title.Default()

	return d
}
