package main

import (
	"encoding/binary"
	"fmt"
	"os"
)

type FileSystem struct {
	file    *os.File
	catalog Catalog
	diskmap DiskMap
}

func NewFileSystem(file *os.File) *FileSystem {
	return &FileSystem{
		file:    file,
		catalog: Catalog{},
		diskmap: NewDiskMap(),
	}
}

func (fs *FileSystem) Flush() error {
	const diskmapStart = 10
	const diskmapSectors = 4
	const catalogStart = 14
	const catalogSectors = 12

	dmBuf := make([]byte, diskmapSectors*SectorBytes)
	size, err := binary.Encode(dmBuf, binary.LittleEndian, fs.diskmap)
	if err != nil || size != cap(dmBuf) {
		return fmt.Errorf("encode diskmap failed: %v %v", size, err)
	}

	catBuf := make([]byte, catalogSectors*SectorBytes)
	size, err = binary.Encode(catBuf, binary.LittleEndian, fs.catalog)
	if err != nil || size != cap(catBuf) {
		return fmt.Errorf("encode catalog failed: %v %v", size, err)
	}

	for i := 0; i < diskmapSectors; i++ {
		start := i * SectorBytes
		end := start + SectorBytes
		err := WriteSector(fs.file, dmBuf[start:end], diskmapStart+i)
		if err != nil {
			return err
		}
	}
	for i := 0; i < catalogSectors; i++ {
		start := i * SectorBytes
		end := start + SectorBytes
		err = WriteSector(fs.file, catBuf[start:end], catalogStart+i)
		if err != nil {
			return err
		}
	}
	return nil
}

func (fs *FileSystem) CreateFile(filename string, data []byte) error {
	name := NewName(filename)
	for i := 0; i < int(fs.catalog.size); i++ {
		if name == fs.catalog.contents[i].name {
			return fmt.Errorf("Duplicate name: %s", NameToString(name))
		}
	}

	if fs.catalog.size == MaxItem {
		return fmt.Errorf("Too many files")
	}
	fileno := fs.catalog.size
	fs.catalog.size++

	var start int16 = endList
	var pages = 0
	var words = 0
	for i := 0; i < len(data); i += PageBytes {
		pages++

		address, err := fs.diskmap.Extend(&start)
		if err != nil {
			return err
		}

		page := make([]byte, PageBytes)
		words = copy(page, data[i:])
		err = WritePage(fs.file, page, int(address))
		if err != nil {
			return err
		}
	}

	fs.catalog.contents[fileno] = Item{
		name: name,
		attr: Attributes{
			addr: start,
			length: Position{
				pages: int16(pages),
				words: int16(words / 2),
			},
			protected: 0,
		},
	}
	return nil
}

func (fs *FileSystem) ReadFile(filename string) ([]byte, error) {
	name := NewName(filename)
	var attr *Attributes
	for i := 0; i < int(fs.catalog.size); i++ {
		if name == fs.catalog.contents[i].name {
			attr = &fs.catalog.contents[i].attr
			break
		}
	}
	if attr == nil {
		return nil, fmt.Errorf("File not found")
	}

	data := make([]byte, attr.length.pages * PageBytes + attr.length.words * 2)
	pos := 0
	for addr := attr.addr; addr != endList; addr = fs.diskmap.contents[addr-1] {
		err := ReadPage(fs.file, data[pos:], int(addr))
		if err != nil {
			return nil, err
		}
		pos += PageBytes
	}
	return data,nil
}

//
// Disk maps
//

type DiskMap struct {
	free, next int16
	contents   [LastPage]int16
	filler     [5]int16
}

const (
	FirstPage = 27
	LastPage  = 249

	endList   = 0
	available = 32767
)

func NewDiskMap() DiskMap {
	dm := DiskMap{}
	for i := 1; i <= LastPage; i++ {
		if i < FirstPage {
			dm.contents[i-1] = endList
		} else {
			dm.contents[i-1] = available
		}
	}
	dm.free = LastPage - FirstPage + 1
	dm.next = FirstPage
	return dm
}

func (dm *DiskMap) Extend(start *int16) (int16, error) {
	if dm.free == 0 {
		return 0, fmt.Errorf("Disk full")
	}
	for dm.contents[dm.next-1] != available {
		dm.next = (dm.next % LastPage) + 1
	}
	if *start == endList {
		*start = dm.next
	} else {
		elem := *start
		succ := dm.contents[elem-1]
		for succ != endList {
			elem = succ
			succ = dm.contents[elem-1]
		}
		dm.contents[elem-1] = dm.next
	}
	dm.contents[dm.next-1] = endList
	dm.free--
	return dm.next, nil
}

//
// Catalogs
//

type Catalog struct {
	size     int16
	contents [MaxItem]Item
	filler   [15]int16
}

const MaxItem = 47

type Item struct {
	name Name
	attr Attributes
}

type Attributes struct {
	addr      int16
	length    Position
	protected int16
}

type Position struct {
	pages, words int16
}

//
// Names
//

type Name = [12]int16

func NewName(s string) (n Name) {
	i := 0
	for _, r := range s {
		if r >= 'a' && r <= 'z' || r >= '0' && r <= '9' || r == '_' {
			n[i] = int16(r)
			i++
		} else if r >= 'A' && r <= 'Z' {
			n[i] = int16(r + 32)
			i++
		}
		if i == len(n) {
			break
		}
	}
	if i == 0 {
		n[i] = 'a'
		i++
	}
	for ; i < len(n); i++ {
		n[i] = ' '
	}
	return
}

func EncodedName(s string) []byte {
	var buf []byte
	n := NewName(s)
	for i := range n {
		buf = append(buf, byte(n[i]), 0)
	}
	return buf
}

func NameToString(n Name) string {
	var s []rune
	for i := range n {
		if n[i] == ' ' {
			break
		} else {
			s = append(s, rune(n[i]))
		}
	}
	return string(s)
}
