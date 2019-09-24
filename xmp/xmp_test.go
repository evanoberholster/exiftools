package xmp_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/evanoberholster/exif/models"

	"github.com/evanoberholster/exif/xmp"

	xmp2 "trimmer.io/go-xmp/xmp"
)

func LoadTestFile() ([]byte, error) {
	f, err := os.Open("../test/test.xmp")
	if err != nil {
		return []byte{}, fmt.Errorf("Could not open test file: %v", err)
	}
	bb, err := ioutil.ReadAll(f)
	if err != nil {
		return []byte{}, fmt.Errorf("Could not read test file: %v", err)
	}
	return bb, nil
}

// TestReadXMPDocument - Need better test
func TestReadXMPDocument(t *testing.T) {
	bb, err := LoadTestFile()
	if err != nil {
		t.Error(err)
	}
	doc := &xmp2.Document{}
	if err := xmp2.Unmarshal(bb, doc); err != nil {
		t.Fatalf("Could not Unmarshal test file: %v", err)
	}
}

func TestBase(t *testing.T) {
	bb, err := LoadTestFile()
	if err != nil {
		t.Error(err)
	}
	doc := &xmp2.Document{}
	if err := xmp2.Unmarshal(bb, doc); err != nil {
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
	bb, err := LoadTestFile()
	if err != nil {
		t.Error(err)
	}
	doc := &xmp2.Document{}
	if err := xmp2.Unmarshal(bb, doc); err != nil {
		t.Fatalf("Could not Unmarshal test file: %v", err)
	}

	c := models.DublinCore{
		Creator:     "John Doe",
		Description: "",
		Format:      "image/x-canon-cr2",
		Rights:      "John Doe",
		Source:      "",
		Subject:     []string{"National Park"},
		Title:       "",
	}
	dc := xmp.DublinCore(doc)
	if c.Creator != dc.Creator {
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
