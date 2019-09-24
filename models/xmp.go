package models

import (
	"time"
	//_ "trimmer.io/go-xmp/models"
)

// XMP

// DublinCore -
///// Modified DublinCore 17/04/2019 https://godoc.org/trimmer.io/go-xmp/models/dc#DublinCore
type DublinCore struct {
	Creator     string   `xmp:"dc:creator"`
	Description string   `xmp:"dc:description"`
	Format      string   `xmp:"dc:format"`
	Rights      string   `xmp:"dc:rights"`
	Source      string   `xmp:"dc:source"`
	Subject     []string `xmp:"dc:subject"`
	Title       string   `xmp:"dc:title"`
}

// XmpBase -
///// Modified XmpBase 17/04/2019 https://godoc.org/trimmer.io/go-xmp/models/xmp_base#XmpBase
type XmpBase struct {
	CreateDate   time.Time `xmp:"xmp:CreateDate"`
	CreatorTool  string    `xmp:"xmp:CreatorTool"`
	Identifier   string    `xmp:"xmp:Identifier"`
	Label        string    `xmp:"xmp:Label"`
	MetadataDate time.Time `xmp:"xmp:MetadataDate"`
	ModifyDate   time.Time `xmp:"xmp:ModifyDate"`
	Rating       int       `xmp:"xmp:Rating"`
}
