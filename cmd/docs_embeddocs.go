// +build !nodocs
// +build !noembeddocs

package cmd

import (
	"os"
	"path/filepath"

	"github.com/rakyll/statik/fs"
)

// DocsDir is unused when chezmoi is built with embedded docs.
var DocsDir = ""

func getDocsFilenames() ([]string, error) {
	var filenames []string
	err := fs.Walk(statikFS, "/docs", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.Mode().IsRegular() {
			filenames = append(filenames, filepath.Base(path))
		}
		return nil
	})
	return filenames, err
}

func getDoc(filename string) ([]byte, error) {
	return fs.ReadFile(statikFS, filepath.Join("/docs", filename))
}
