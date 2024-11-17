package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func MakeDisk(fn_out, fn_kernel, fn_system string, fn_files []string) error {
	out, err := os.OpenFile(fn_out, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0644)
	if err != nil {
		return err
	}
	defer out.Close()

	kernel, err := os.ReadFile(fn_kernel)
	if err != nil {
		return err
	}
	err = WriteKernel(out, kernel)
	if err != nil {
		return err
	}

	system, err := os.ReadFile(fn_system)
	if err != nil {
		return err
	}
	err = WriteSystem(out, fn_system, system)
	if err != nil {
		return err
	}

	fs := NewFileSystem(out)
	for i := range fn_files {
		data, err := os.ReadFile(fn_files[i])
		if err != nil {
			return err
		}
		if isText(data) {
			data = convertText(data)
		}
		err = fs.CreateFile(filepath.Base(fn_files[i]), data)
		if err != nil {
			return err
		}
	}
	return fs.Flush()
}

func isText(data []byte) bool {
	for i := range data {
		if data[i] == 0 || data[i] >= 127 {
			return false
		}
	}
	return true
}

func convertText(text []byte) []byte {
	var buf []byte
	for i := range text {
		buf = append(buf, text[i], 0)
	}
	return buf
}

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: make-disk disk.img kernel system.code files...")
		os.Exit(1)
	}

	err := MakeDisk(os.Args[1], os.Args[2], os.Args[3], os.Args[4:])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
