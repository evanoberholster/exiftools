package exiftool

import (
	"fmt"

	"github.com/evanoberholster/exiftools/exiftool/exif"
)

var (
	ErrIfdNotValid = fmt.Errorf("Ifd not Valid")
)

// LoadIfds loads ifdItems in the IfdMapping
func (im *IfdMapping) LoadIfds(ifdItem IfdItem) (*IfdMapping, error) {
	var err error
	err = im.addIfdItem(ifdItem)
	// add IFD
	//im.addIfd()
	return im, err
}

//func (im *IfdMapping) addIfdItem(ifdPath exif.IfdPath, ifdPointer exif.TagID, name string) (err error) {
func (im *IfdMapping) addIfdItem(ifd IfdItem) (err error) {
	defer func() {
		if state := recover(); state != nil {
			err = state.(error)
		}
	}()
	if !ifd.Valid() {
		return ErrIfdNotValid
	}

	ptr, err := im.GetParentPlacement(ifd.IfdPath)
	if err != nil {
		panic(err)
	}

	// Define Path
	path := make([]string, len(ifd.IfdPath)+1)
	if len(ifd.IfdPath) > 0 {
		copy(path, ptr.Path)
	}
	path[len(path)-1] = ifd.Name

	// Define Placement
	placement := make(exif.IfdPath, len(ifd.IfdPath)+1)
	if len(placement) > 0 {
		copy(placement, ptr.Placement)
	}
	placement[len(placement)-1] = ifd.TagID

	fmt.Println(ifd.Name, path, ifd.IfdPath, ifd.TagID)
	childIfd := &MappedIfd{
		ParentTagID: ptr.TagID,
		Path:        path,
		Placement:   placement,
		Name:        ifd.Name,
		TagID:       ifd.TagID,
		Children:    make(map[exif.TagID]*MappedIfd),
	}

	if _, found := ptr.Children[ifd.TagID]; found == true {
		panic(fmt.Errorf("Child IFD [%s] with tag-ID (0x%04x) already registered under IFD [%s] with tag-ID (0x%04x)", ifd.Name, ifd.TagID, ptr.Name, ptr.TagID))
	}

	ptr.Children[ifd.TagID] = childIfd

	return nil
}

func (im *IfdMapping) GetParentPlacement(parentPlacement exif.IfdPath) (childIfd *MappedIfd, err error) {
	defer func() {
		if state := recover(); state != nil {
			err = state.(error)
		}
	}()

	ptr := im.rootNode
	for _, ifdTagID := range parentPlacement {
		if descendantPtr, found := ptr.Children[ifdTagID]; found == false {
			panic(fmt.Errorf("IFD [%s] with tag-ID (0x%04x) not registered ", ptr.PathPhrase(), ifdTagID))
		} else {
			ptr = descendantPtr
		}
	}

	return ptr, nil
}
