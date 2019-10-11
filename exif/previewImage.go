package exif

import "fmt"

var exifCompressionValues = map[uint16]string{
	6:     "JPEG (old-style)",
	7:     "JPEG",
	99:    "JPEG",
	34712: "JPEG 2000",
	34892: "Lossy JPEG",
	34934: "JPEG XR",
	34927: "WebP",
	34933: "PNG",
}

// PreviewImageTag -
type PreviewImageTag struct {
	StartTag    FieldName
	LengthTag   FieldName
	Compression FieldName
	Start       int
	Length      int
}

// NewPreviewImageTag -
func NewPreviewImageTag(start FieldName, length FieldName, compression FieldName) PreviewImageTag {
	return PreviewImageTag{start, length, compression, 0, 0}
}

// PreviewImage returns the byte start location and length of the preview Image.
func (x Exif) PreviewImage(tags ...PreviewImageTag) (int64, int64, error) {
	tags = append(tags,
		NewPreviewImageTag(PreviewImageStart, PreviewImageLength, FieldName("None")),                        // IFD0 PreviewImage
		NewPreviewImageTag(ThumbJPEGInterchangeFormat, ThumbJPEGInterchangeFormatLength, FieldName("None")), // IFD0 ThumbnailImage
	)
	for i, tag := range tags {
		if tag.Compression != FieldName("None") {
			compression, err := x.Get(tag.Compression)
			if err == nil {
				c, err := compression.Int(0)
				if err != nil {
					continue
				}
				_, ok := exifCompressionValues[uint16(c)]
				if !ok {
					continue
				}
			}
		}
		offset, err := x.Get(tag.StartTag)
		if err != nil {
			continue
		}
		tags[i].Start, err = offset.Int(0)
		if err != nil {
			continue
		}
		length, err := x.Get(tag.LengthTag)
		if err != nil {
			continue
		}
		tags[i].Length, err = length.Int(0)
		if err != nil {
			continue
		}
	}

	var maxTag PreviewImageTag
	for i := range tags {
		if tags[i].Length > maxTag.Length {
			maxTag = tags[i]
		}
	}
	fmt.Println(maxTag)
	return int64(maxTag.Start), int64(maxTag.Length), nil
}
