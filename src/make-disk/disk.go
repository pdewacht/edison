package main

import (
	"fmt"
	"os"
	"path/filepath"
)

const SectorBytes = 128
const PageBytes = 1024
const PageSectors = PageBytes / SectorBytes

func WriteSector(f *os.File, sector []byte, sectorno int) error {
	if sectorno < 0 || sectorno > 2001 {
		return fmt.Errorf("sector out of range")
	}
	if len(sector) != SectorBytes {
		return fmt.Errorf("Invalid sector size")
	}
	n := sectorno - (sectorno % 26) + (sectorno * 3 % 26)
	_, err := f.WriteAt(sector, int64(n*SectorBytes))
	return err
}

func ReadSector(f *os.File, sector []byte, sectorno int) error {
	if sectorno < 0 || sectorno > 2001 {
		return fmt.Errorf("sector out of range")
	}
	if cap(sector) > SectorBytes {
		sector = sector[0:SectorBytes]
	}
	n := sectorno - (sectorno % 26) + (sectorno * 3 % 26)
	_, err := f.ReadAt(sector, int64(n*SectorBytes))
	return err
}

func WritePage(f *os.File, page []byte, pageno int) error {
	if len(page) != PageBytes {
		return fmt.Errorf("Invalid page size")
	}
	sectorno := PageSectors*pageno + 2
	for i := 0; i < PageSectors; i++ {
		start := i * SectorBytes
		end := start + SectorBytes
		err := WriteSector(f, page[start:end], sectorno+i)
		if err != nil {
			return err
		}
	}
	return nil
}

func ReadPage(f *os.File, page []byte, pageno int) error {
	if cap(page) > PageBytes {
		page = page[0:PageBytes]
	}
	sectorno := PageSectors*pageno + 2
	for i := 0; i < PageSectors; i++ {
		start := i * SectorBytes
		end := start + SectorBytes
		err := ReadSector(f, page[start:end], sectorno+i)
		if err != nil {
			return err
		}
	}
	return nil
}

func WriteRaw(f *os.File, data []byte, pageno int) error {
	for i := 0; i*PageBytes < len(data); i++ {
		page := make([]byte, PageBytes)
		copy(page, data[i*PageBytes:])
		err := WritePage(f, page, pageno+i)
		if err != nil {
			return err
		}
	}
	return nil
}

func WriteKernel(f *os.File, kernel []byte) error {
	if len(kernel) > 8*PageBytes {
		return fmt.Errorf("Kernel too large")
	}
	return WriteRaw(f, kernel, 3)
}

func WriteSystem(f *os.File, name string, system []byte) error {
	sysname := EncodedName(filepath.Base(name))
	data := append(sysname, system...)
	if len(data) > 16*PageBytes {
		return fmt.Errorf("System too large")
	}
	return WriteRaw(f, data, 11)
}
