package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {
	p := fmt.Println
	pth := filepath.Join("dir1", "dir2", "filename")
	fmt.Println("p:", pth)

	p(filepath.Join("dir1///", "filename"))
	p(filepath.Join("dir1/../dir1", "filename"))
	p("Dir:", filepath.Dir(pth))
	p("Base:", filepath.Base(pth))
	p(filepath.IsAbs("./dir/file"))
	p(filepath.IsAbs("/dir/file"))

	filename := "config.json"
	ext := filepath.Ext(filename)
	p(ext)

	// Get file name
	p(strings.TrimSuffix(filename, ext))

	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	p(rel)

	// relative path
	rel, err = filepath.Rel("a/b", "a/c/t/file")
	if err != nil {
		panic(err)
	}
	p(rel)
}
