package xmp_test

import (
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/evanoberholster/exiftools/models"
	"github.com/evanoberholster/exiftools/xmp"
)

const JPEGTestFile = "../exif/samples/test.jpg"
const XMPTestFile = "../exif/samples/test.xmp"

// TestReadXMPDocument - Test needs improvement
func TestReadXMPDocument(t *testing.T) {
	f, err := os.Open(JPEGTestFile)
	if err != nil {
		t.Fatalf("Could not open test file: %v", err)
	}
	_, err = xmp.ReadXMPDocument(f)
	if err != nil {
		t.Fatalf("Could not read test XMP file: %v", err)
	}
}

// TestUnMarshal - Test needs improvement
func TestUnMarshal(t *testing.T) {
	f, err := os.Open(XMPTestFile)
	if err != nil {
		t.Fatalf("Could not open test file: %v", err)
	}
	bb, err := ioutil.ReadAll(f)
	if err != nil {
		t.Fatalf("Could not read test file: %v", err)
	}
	_, err = xmp.Unmarshal(bb)
	if err != nil {
		t.Fatalf("Could not Unmarshal test file: %v", err)
	}
}

func TestBase(t *testing.T) {
	f, err := os.Open(XMPTestFile)
	if err != nil {
		t.Fatalf("Could not open test file: %v", err)
	}
	bb, err := ioutil.ReadAll(f)
	if err != nil {
		t.Fatalf("Could not read test file: %v", err)
	}
	doc, err := xmp.Unmarshal(bb)
	if err != nil {
		t.Fatalf("Could not Unmarshal test file: %v", err)
	}

	createDate, _ := time.Parse(time.RFC3339, "2018-04-14T15:33:20.61Z")
	modifyDate, _ := time.Parse(time.RFC3339, "2018-04-14T15:33:20.61Z")
	metadataDate, _ := time.Parse(time.RFC3339, "2018-04-16T08:50:26.23Z")
	b := models.XmpBase{
		CreateDate:   createDate,
		CreatorTool:  "",
		Identifier:   "",
		Label:        "",
		MetadataDate: metadataDate,
		ModifyDate:   modifyDate,
		Rating:       0,
	}
	base := xmp.Base(doc)
	if b.CreateDate != base.CreateDate {
		t.Fatal("CreateDate does not match", base, b)
	}
	if b.CreatorTool != base.CreatorTool {
		t.Fatal("CreatorTool does not match", base, b)
	}
	if b.Identifier != base.Identifier {
		t.Fatal("Identifier does not match", base, b)
	}
	if b.Label != base.Label {
		t.Fatal("Label does not match", base, b)
	}
	if b.MetadataDate != base.MetadataDate {
		t.Fatal("MetadataDate does not match", base, b)
	}
	if b.ModifyDate != base.ModifyDate {
		t.Fatal("Label does not match", base, b)
	}
	if b.Rating != base.Rating {
		t.Fatal("Rating does not match", base, b)
	}
}

func TestDublinCore(t *testing.T) {
	f, err := os.Open(XMPTestFile)
	if err != nil {
		t.Fatalf("Could not open test file: %v", err)
	}
	bb, err := ioutil.ReadAll(f)
	if err != nil {
		t.Fatalf("Could not read test file: %v", err)
	}
	doc, err := xmp.Unmarshal(bb)
	if err != nil {
		t.Fatalf("Could not Unmarshal test file: %v", err)
	}

	c := models.DublinCore{
		Creator:     []string{"John Doe"},
		Description: "",
		Format:      "image/x-canon-cr2",
		Rights:      "John Doe",
		Source:      "",
		Subject:     []string{"National Park"},
		Title:       "",
	}
	dc := xmp.DublinCore(doc)
	if c.Creator[0] != dc.Creator[0] {
		t.Fatal("Creator does not match", dc.Creator, c.Creator)
	}
	if c.Description != dc.Description {
		t.Fatal("Description does not match", dc.Description, c.Description)
	}
	if c.Format != dc.Format {
		t.Fatal("Format does not match", dc.Format, c.Format)
	}
	if c.Rights != dc.Rights {
		t.Fatal("Rights does not match", dc.Rights, c.Rights)
	}
	if c.Source != dc.Source {
		t.Fatal("Source does not match", dc.Source, c.Source)
	}
	if c.Subject[0] != dc.Subject[0] {
		t.Fatal("Subject does not match", dc.Subject, c.Subject)
	}
	if c.Title != dc.Title {
		t.Fatal("Title does not match", dc.Title, c.Title)
	}
}
