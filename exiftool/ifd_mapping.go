package exiftool

import (
	"errors"
	"fmt"
	"strings"

	"github.com/evanoberholster/exiftools/exiftool/exif"
	"github.com/evanoberholster/exiftools/exiftool/tags/ifd"
)

// Errors
var (
	ErrChildIfdNotMapped = errors.New("no child-IFD for that tag-ID under parent")
)

type IfdTagIdAndIndex struct {
	Name  string
	TagID exif.TagID
	Index int
}

// IfdMapping describes all of the IFDs that we currently recognize.
type IfdMapping struct {
	rootNode *MappedIfd
}

func (im *IfdMapping) PathPhraseFromLineage(lineage []IfdTagIdAndIndex) (pathPhrase string) {
	pathParts := make([]string, len(lineage))
	for i, itii := range lineage {
		pathParts[i] = itii.Name
	}

	return strings.Join(pathParts, "/")
}

// StripPathPhraseIndices returns a non-fully-qualified path-phrase (no
// indices).
func (im *IfdMapping) StripPathPhraseIndices(pathPhrase string) (strippedPathPhrase string, err error) {
	defer func() {
		if state := recover(); state != nil {
			err = state.(error)
		}
	}()

	lineage, err := im.ResolvePath(pathPhrase)
	if err != nil {
		panic(err)
	}

	strippedPathPhrase = im.PathPhraseFromLineage(lineage)
	return strippedPathPhrase, nil
}

// ResolvePath takes a list of names, which can also be suffixed with indices
// (to identify the second, third, etc.. sibling IFD) and returns a list of
// tag-IDs and those indices.
//
// Example:
//
// - IFD/Exif/Iop
// - IFD0/Exif/Iop
//
// This is the only call that supports adding the numeric indices.
func (im *IfdMapping) ResolvePath(pathPhrase string) (lineage []IfdTagIdAndIndex, err error) {
	defer func() {
		if state := recover(); state != nil {
			err = state.(error)
			//err = log.Wrap(state.(error))
		}
	}()

	pathPhrase = strings.TrimSpace(pathPhrase)

	if pathPhrase == "" {
		panic("Can not resolve empty path-phrase")
		//log.Panicf("can not resolve empty path-phrase")
	}

	path := strings.Split(pathPhrase, "/")
	lineage = make([]IfdTagIdAndIndex, len(path))

	ptr := im.rootNode
	empty := IfdTagIdAndIndex{}
	for i, name := range path {
		indexByte := name[len(name)-1]
		index := 0
		if indexByte >= '0' && indexByte <= '9' {
			index = int(indexByte - '0')
			name = name[:len(name)-1]
		}

		itii := IfdTagIdAndIndex{}
		for _, mi := range ptr.Children {
			if mi.Name != name {
				continue
			}

			itii.Name = name
			itii.TagID = mi.TagID
			itii.Index = index

			ptr = mi

			break
		}

		if itii == empty {
			panic(fmt.Errorf("Ifd child with name [%s] not registered: [%s]", name, pathPhrase))
			//log.Panicf("ifd child with name [%s] not registered: [%s]", name, pathPhrase)
		}

		lineage[i] = itii
	}

	return lineage, nil
}

// GetChild is a convenience function to get the child path for a given parent
// placement and child tag-ID.
func (im *IfdMapping) GetChild(parentPathPhrase string, tagID exif.TagID) (mi *MappedIfd, err error) {
	defer func() {
		if state := recover(); state != nil {
			err = state.(error)
		}
	}()

	if mi, err = im.GetWithPath(parentPathPhrase); err != nil {
		return nil, fmt.Errorf("GetChild Error: %s %w", parentPathPhrase, err)
	}

	for _, childMi := range mi.Children {
		if childMi.TagID == tagID {
			return childMi, nil
		}
	}

	// Whether or not an IFD is defined in data, such an IFD is not registered and would be unknown.
	return nil, ErrChildIfdNotMapped
}

func (im *IfdMapping) GetWithPath(pathPhrase string) (mi *MappedIfd, err error) {
	defer func() {
		if state := recover(); state != nil {
			err = state.(error)
			//err = log.Wrap(state.(error))
		}
	}()

	if pathPhrase == "" {
		return nil, fmt.Errorf("Path-phrase is empty")
	}

	path := strings.Split(pathPhrase, "/")
	ptr := im.rootNode

	for _, name := range path {
		var hit *MappedIfd
		for _, mi := range ptr.Children {
			if mi.Name == name {
				hit = mi
				break
			}
		}

		if hit == nil {
			panic(fmt.Errorf("ifd child with name [%s] not registered: [%s]", name, ptr.PathPhrase()))
			//log.Panicf("ifd child with name [%s] not registered: [%s]", name, ptr.PathPhrase())
		}

		ptr = hit
	}

	return ptr, nil
}

func (mi *MappedIfd) PathPhrase() string {
	return strings.Join(mi.Path, "/")
}

var (
	ErrIfdNotValid = fmt.Errorf("Ifd not Valid")
)

// LoadIfds loads ifdItems in the IfdMapping
func (im *IfdMapping) LoadIfds(ifds ...ifd.IfdItem) (*IfdMapping, error) {
	var err error
	for _, item := range ifds {
		if err = im.addIfdItem(item); err != nil {
			panic(err)
		}
	}

	return im, err
}

//func (im *IfdMapping) addIfdItem(ifdPath exif.IfdPath, ifdPointer exif.TagID, name string) (err error) {
func (im *IfdMapping) addIfdItem(ifdItem ifd.IfdItem) (err error) {
	defer func() {
		if state := recover(); state != nil {
			err = state.(error)
		}
	}()
	if !ifdItem.Valid() {
		return ErrIfdNotValid
	}

	ptr, err := im.GetParentPlacement(ifdItem.IfdPath)
	if err != nil {
		panic(err)
	}

	// Define Path
	path := make([]string, len(ifdItem.IfdPath)+1)
	if len(ifdItem.IfdPath) > 0 {
		copy(path, ptr.Path)
	}
	path[len(path)-1] = ifdItem.Name

	// Define Placement
	placement := make(ifd.IfdPath, len(ifdItem.IfdPath)+1)
	if len(placement) > 0 {
		copy(placement, ptr.Placement)
	}
	placement[len(placement)-1] = ifdItem.TagID

	childIfd := &MappedIfd{
		ParentTagID: ptr.TagID,
		Path:        path,
		Placement:   placement,
		Name:        ifdItem.Name,
		TagID:       ifdItem.TagID,
		Children:    make(map[exif.TagID]*MappedIfd),
	}

	if _, found := ptr.Children[ifdItem.TagID]; found == true {
		panic(fmt.Errorf("Child IFD [%s] with tag-ID (0x%04x) already registered under IFD [%s] with tag-ID (0x%04x)", ifdItem.Name, ifdItem.TagID, ptr.Name, ptr.TagID))
	}

	ptr.Children[ifdItem.TagID] = childIfd

	return nil
}

func (im *IfdMapping) GetParentPlacement(parentPlacement ifd.IfdPath) (childIfd *MappedIfd, err error) {
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
