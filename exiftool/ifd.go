package exiftool

import (
	"errors"
	"fmt"

	"github.com/evanoberholster/exiftools/exiftool/exif"
)

// Errors
var (
	ErrChildIfdNotMapped = errors.New("no child-IFD for that tag-ID under parent")
)

func (im *IfdMapping) AddPath(ifdPath exif.IfdPath, ifdPointer exif.TagID, name string, tagNameMap map[exif.TagID]string) (err error) {
	defer func() {
		if state := recover(); state != nil {
			err = state.(error)
		}
	}()

	// TODO(dustin): !! It would be nicer to provide a list of names in the placement rather than tag-IDs.

	ptr, err := im.GetParentPlacement(ifdPath[:len(ifdPath)-1])
	if err != nil {
		panic(err)
	}

	path := make([]string, len(ifdPath)+1)
	if len(ifdPath) > 0 {
		copy(path, ptr.Path)
	}

	path[len(path)-1] = name

	placement := make(exif.IfdPath, len(ifdPath)+1)
	if len(placement) > 0 {
		copy(placement, ptr.Placement)
	}

	placement[len(placement)-1] = ifdPointer
	fmt.Println(name, path)
	childIfd := &MappedIfd{
		ParentTagID: ptr.TagID,
		Path:        path,
		Placement:   placement,
		Name:        name,
		TagID:       ifdPointer,
		Children:    make(map[exif.TagID]*MappedIfd),
	}

	if _, found := ptr.Children[ifdPointer]; found == true {
		panic(fmt.Errorf("Child IFD with tag-ID (%04x) already registered under IFD [%s] with tag-ID (%04x)", ifdPointer, ptr.Name, ptr.TagID))
	}

	ptr.Children[ifdPointer] = childIfd

	return nil
}

// Add puts the given IFD at the given position of the tree. The position of the
// tree is referred to as the placement and is represented by a set of tag-IDs,
// where the leftmost is the root tag and the tags going to the right are
// progressive descendants.
func (im *IfdMapping) Add(parentPlacement []uint16, tagID exif.TagID, name string) (err error) {
	defer func() {
		if state := recover(); state != nil {
			err = state.(error)
		}
	}()

	// TODO(dustin): !! It would be nicer to provide a list of names in the placement rather than tag-IDs.

	ptr, err := im.Get(parentPlacement)
	if err != nil {
		panic(err)
	}

	path := make([]string, len(parentPlacement)+1)
	if len(parentPlacement) > 0 {
		copy(path, ptr.Path)
	}

	path[len(path)-1] = name

	placement := make(exif.IfdPath, len(parentPlacement)+1)
	if len(placement) > 0 {
		copy(placement, ptr.Placement)
	}

	placement[len(placement)-1] = tagID

	childIfd := &MappedIfd{
		ParentTagID: ptr.TagID,
		Path:        path,
		Placement:   placement,
		Name:        name,
		TagID:       tagID,
		Children:    make(map[exif.TagID]*MappedIfd),
	}

	if _, found := ptr.Children[tagID]; found == true {
		panic(fmt.Errorf("Child IFD with tag-ID (%04x) already registered under IFD [%s] with tag-ID (%04x)", tagID, ptr.Name, ptr.TagID))
	}

	ptr.Children[tagID] = childIfd

	return nil
}

func (im *IfdMapping) Get(parentPlacement []uint16) (childIfd *MappedIfd, err error) {
	defer func() {
		if state := recover(); state != nil {
			//err = log.Wrap(state.(error))
			err = state.(error)
		}
	}()

	ptr := im.rootNode
	for _, tagID := range parentPlacement {
		if descendantPtr, found := ptr.Children[exif.TagID(tagID)]; found == false {
			panic(fmt.Errorf("IFD child with tag-ID (%04x) not registered: [%s]", tagID, ptr.PathPhrase()))
			//log.Panicf("ifd child with tag-ID (%04x) not registered: [%s]", tagId, ptr.PathPhrase())
		} else {
			ptr = descendantPtr
		}
	}

	return ptr, nil
}
