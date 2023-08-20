package utils

import (
	"go-ffprope/pkg/store"
	"io/fs"
	"path/filepath"
)

func GetFilePaths(src string) ([]string, error) {
	paths := []string{}
	err := filepath.WalkDir(src, func(path string, d fs.DirEntry, err error) error {
		if !d.IsDir() {

			ext := filepath.Ext(d.Name())

			for _, e := range store.Extensions {
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
