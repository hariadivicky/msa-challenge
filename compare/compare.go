package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// Dir .
type Dir struct {
	Prefix   string
	Path     string
	Contents []interface{}
}

// NewDir constructs Dir structure.
func NewDir(prefix string, path string) *Dir {
	prefix = normalizePath(prefix, false)
	path = normalizePath(path, true)

	return &Dir{
		Prefix: prefix,
		Path:   path,
	}
}

// Compare directory contents.
func Compare(source, target string) ([]string, error) {
	var result []string

	// check source existence
	sourceInfo, err := os.Stat(source)
	if err != nil {
		return result, err
	}

	// only accept directory.
	if !sourceInfo.IsDir() {
		return result, fmt.Errorf("source is not a directory")
	}

	// check target existence
	targetInfo, err := os.Stat(target)
	if err != nil {
		return result, err
	}

	// only accept directory.
	if !targetInfo.IsDir() {
		return result, fmt.Errorf("target is not a directory")
	}

	// read source contents.
	sourceDir := NewDir(source, "")
	sourceContents, err := sourceDir.Read()
	if err != nil {
		return result, err
	}

	// read target contents.
	targetDir := NewDir(target, "")
	targetContents, err := targetDir.Read()
	if err != nil {
		return result, err
	}

	for _, sourceContent := range sourceContents {
		if !inSlices(targetContents, sourceContent) {
			result = append(result, sourceContent+" [DELETED]")
		} else {
			s, err := ioutil.ReadFile(sourceDir.Prefix + sourceContent)
			if err != nil {
				return result, err
			}

			t, err := ioutil.ReadFile(targetDir.Prefix + sourceContent)
			if err != nil {
				return result, err
			}

			if !bytes.Equal(s, t) {
				result = append(result, sourceContent+" [MODIFIED]")
			}
		}
	}

	for _, targetContent := range targetContents {
		if !inSlices(sourceContents, targetContent) {
			result = append(result, targetContent+" [NEW]")
		}
	}

	return result, nil
}

// subDir creates Dir structure with parent prefix.
func (d *Dir) subDir(path string) *Dir {
	return NewDir(d.Prefix, d.Path+path)
}

// Read direcotry contents.
func (d *Dir) Read() ([]string, error) {
	var result []string
	dir, err := os.Open(d.Prefix + d.Path)
	if err != nil {
		return nil, err
	}

	contents, err := dir.Readdir(-1)
	if err != nil {
		return nil, err
	}

	dir.Close()

	for _, content := range contents {
		if content.IsDir() {
			subDir := d.subDir(content.Name())

			subContents, err := subDir.Read()
			if err != nil {
				return nil, err
			}

			result = append(result, subContents...)
		} else {
			path := d.Path + content.Name()
			result = append(result, path)
		}

	}

	return result, nil
}

// inSlices checks if search element is in slices.
func inSlices(list []string, search string) bool {
	for _, val := range list {
		if val == search {
			return true
		}
	}

	return false
}

func normalizePath(path string, isRelative bool) string {
	// normalize trailing slash.
	sep := string(filepath.Separator)
	if isRelative {
		path = strings.TrimRight(path, sep)
	} else {
		path = strings.Trim(path, sep)
	}
	return path + "/"
}
