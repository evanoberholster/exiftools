package models

import (
	"time"
	//_ "trimmer.io/go-xmp/models"
)

// XMP

// DublinCore - Dublin Core metadata
///// Modified DublinCore 17/03/2020 https://godoc.org/trimmer.io/go-xmp/models/dc#DublinCore
type DublinCore struct {
	Creator     []string    `xmp:"dc:creator"`
	Date        []time.Time `xmp:"dc:date"`
	Description string      `xmp:"dc:description"`
	Format      string      `xmp:"dc:format"`
	Rights      string      `xmp:"dc:rights"`
	Source      string      `xmp:"dc:source"`
	Subject     []string    `xmp:"dc:subject"`
	Title       string      `xmp:"dc:title"`
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

// CC - Creative Commons Metadata
///// Modified XmpBase 18/03/2020
type CC struct {
	License         string `xmp:"cc:license"`
	MorePermissions string `xmp:"cc:morePermissions"`
	AttributionURL  string `xmp:"cc:attributionURL"`
	AttributionName string `xmp:"cc:attributionName"`
}

// XmpDJI - video cameras on DJI drones
///// Modified XmpDJI 18/03/2020
type XmpDJI struct {
	AbsoluteAltitude  float32 `xmp:"drone-dji:AbsoluteAltitude"`  //  "+543.44",
	FlightPitchDegree float32 `xmp:"drone-dji:FlightPitchDegree"` //  "+4.80",
	FlightRollDegree  float32 `xmp:"drone-dji:FlightRollDegree"`  //  "+1.30",
	FlightYawDegree   float32 `xmp:"drone-dji:FlightYawDegree"`   //  "-1.90",
	GimbalPitchDegree float32 `xmp:"drone-dji:GimbalPitchDegree"` //  "-90.00",
	GimbalRollDegree  float32 `xmp:"drone-dji:GimbalRollDegree"`  //  "+0.00",
	GimbalYawDegree   float32 `xmp:"drone-dji:GimbalYawDegree"`   //  "-2.00",
	RelativeAltitude  float32 `xmp:"drone-dji:RelativeAltitude"`  //  "+46.60"
	SpeedX            float32 `xmp:"drone-dji:SpeedX"`
	SpeedY            float32 `xmp:"drone-dji:SpeedY"` // - +0.00
	SpeedZ            float32 `xmp:"drone-dji:SpeedZ"` //  - +0.00,+0.40
	Model             string  `xmp:"tiff:Model"`
}
