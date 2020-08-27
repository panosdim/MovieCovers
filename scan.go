package main

import (
	"os"
	"path/filepath"
	"regexp"
)

func scanMovies(path string, exclude string) []*movie {
	var files []*movie

	exclDirs := regexp.MustCompile(exclude)
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if exclude != "" {
			if info.IsDir() && exclDirs.MatchString(info.Name()) {
				return filepath.SkipDir
			}
		}

		if filepath.Ext(path) == ".mkv" || filepath.Ext(path) == ".mp4" {
			mov := newMovie(path)
			if mov != nil && !mov.hasCover() {
				files = append(files, mov)
			}
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	return files
}
