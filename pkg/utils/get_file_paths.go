package utils

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/mdyssr/azmena/pkg/store"
	"github.com/mdyssr/azmena/pkg/types"
)

func trimExtension(ext string) string {
	return strings.TrimPrefix(ext, ".")
}

func getFilesInFlatPath(src string, options types.RunOptions) ([]string, error) {
	paths := []string{}

	dirs, err := os.ReadDir(src)
	if err != nil {
		return []string{}, err
	}

	for _, dir := range dirs {
		if dir.IsDir() {
			continue
		}

		for _, e := range options.Extensions {
			ext := trimExtension(filepath.Ext(dir.Name()))
			if e == ext {
				paths = append(paths, filepath.Join(src, dir.Name()))
			}
		}
	}

	return paths, nil

}

func GetFilePaths(src string, options types.RunOptions) ([]string, error) {

	if len(options.Extensions) == 0 {
		options.Extensions = store.Extensions
	}

	if options.IsFlat {
		return getFilesInFlatPath(src, options)
	}

	paths := []string{}
	err := filepath.WalkDir(src, func(path string, d fs.DirEntry, err error) error {
		if !d.IsDir() {

			ext := trimExtension(filepath.Ext(d.Name()))
			for _, e := range options.Extensions {
				if e == ext {
					paths = append(paths, path)
					break
				}
			}
		}
		return nil
	})

	if err != nil {
		return []string{}, err
	}
	return paths, nil
}
