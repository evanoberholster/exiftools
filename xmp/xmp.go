// Package xmp aliases the package foung in (trimmer.io/go-xmp/xmp)
package xmp

import (
	"io"

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

// GetBase - Extract XmpBase from XMP Document
func GetBase(m *xmp.Document) XmpBase {
	if c := xmpbase.FindModel(m); c != nil {
		return XmpBase{
			CreateDate:   c.CreateDate.Value(),
			MetadataDate: c.MetadataDate.Value(),
			ModifyDate:   c.ModifyDate.Value(),
			Label:        string(c.Label),
			Rating:       int(c.Rating),
			CreatorTool:  c.CreatorTool.String(),
		}
	}

	return XmpBase{}
}

// GetDublinCore - Extract DublinCore from XMP Document
func GetDublinCore(m *xmp.Document) DublinCore {
	if c := dc.FindModel(m); c != nil {
		return DublinCore{
			Creator:     []string(c.Creator),
			Description: c.Description.Default(),
			Format:      string(c.Format),
			Rights:      c.Rights.Default(),
			Source:      "",
			Subject:     []string(c.Subject),
			Title:       c.Title.Default(),
		}
	}
	return DublinCore{}
}
