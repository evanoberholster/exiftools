package main

import (
	"fmt"
	"log"
	"os"

	"github.com/evanoberholster/exif/exif"
	"github.com/evanoberholster/exif/mknote"
)

type Exif struct {
	FileSize					int64 		`json:"FileSize"`
	MIMEType					string 		`json:"MIMEType"`
	ImageWidth              	int     	`json:"ImageWidth"`
	ImageHeight             	int     	`json:"ImageHeight"`
	CameraMake                  string  	`json:"CameraMake"`
	CameraModel                 string  	`json:"CameraModel"`
	CameraSerial				string 		`json:"CameraSerial"`
	LensModel					string 		`json:"LensModel"`
	LensSerial					string 		`json:"LensSerial"`
}

// MIME

// EXIF

// JPEG Size