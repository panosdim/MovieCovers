package main

import (
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

const (
	// MP4 File type
	MP4 = iota
	// MKV File type
	MKV
)

type movie struct {
	path     string
	fileType int
	name     string
	year     string
}

func newMovie(file string) *movie {
	m := new(movie)
	movieName := regexp.MustCompile(`([ .\w']+?)\W(\d{4})\W?.*`)

	m.path = file
	if filepath.Ext(file) == ".mkv" {
		m.fileType = MKV
	}
	if filepath.Ext(file) == ".mp4" {
		m.fileType = MP4
	}
	m.name = filepath.Base(file)

	matches := movieName.FindStringSubmatch(m.name)
	if matches != nil {
		m.name = strings.Title(strings.ToLower(strings.ReplaceAll(matches[1], ".", " ")))
		m.year = matches[2]
	} else {
		return nil
	}

	return m
}

func (m *movie) hasCover() bool {
	log.Println("Check for cover " + m.path)
	if m.fileType == MKV {
		info := mkvInfo(m.path)
		return strings.Contains(info, "Attachment") && strings.Contains(info, "image")
	}
	if m.fileType == MP4 {
		info := mp4Info(m.path)
		return strings.Contains(info, "covr")
	}
	return false
}

func (m *movie) attachCover(cover *os.File) {
	if m.fileType == MKV {
		mkvSaveCover(m.path, cover)
	}
	if m.fileType == MP4 {
		mp4SaveCover(m.path, cover)
	}
}
