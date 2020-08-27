package main

import (
	"flag"
	"log"
	"os"
	"os/exec"

	"github.com/magiconair/properties"
	"github.com/ryanbradynd05/go-tmdb"
)

var tmdbAPI *tmdb.TMDb
var imageBaseURL string
var atomicParsley string
var mkvMerge string
var mkvPropEdit string

func main() {
	// Read token from properties file
	properties.ErrorHandler = properties.PanicHandler
	p := properties.MustLoadFile("app.properties", properties.UTF8)

	token := p.GetString("token", "")

	// Configure TMDB
	config := tmdb.Config{
		APIKey:   token,
		Proxies:  nil,
		UseProxy: false,
	}

	tmdbAPI = tmdb.Init(config)

	tmdbConf, err := tmdbAPI.GetConfiguration()
	if err != nil {
		log.Fatal(err)
	}

	var posterSize string
	for i, size := range tmdbConf.Images.PosterSizes {
		if i == 0 {
			posterSize = size
		}
		if size == "w500" {
			posterSize = size
		}
	}

	imageBaseURL = tmdbConf.Images.SecureBaseURL + posterSize

	// Find if AtomicParsley, mkvmerge and mkvpropedit are installed
	atomicParsley, err = exec.LookPath("AtomicParsley")
	if err != nil {
		log.Fatal("AtomicParsley not found in path. Please install it.")
	}

	mkvPropEdit, err = exec.LookPath("mkvpropedit")
	if err != nil {
		log.Fatal("mkvpropedit not found in path. Please install it.")
	}

	mkvMerge, err = exec.LookPath("mkvmerge")
	if err != nil {
		log.Fatal("mkvmerge not found in path. Please install it.")
	}

	// Check for flags
	var path string
	var exclude string
	flag.StringVar(&path, "p", "", "Path to scan for movies (required)")
	flag.StringVar(&exclude, "e", "", "Regex to exclude from scan")
	flag.Parse()

	if path == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	files := scanMovies(path, exclude)

	var options = make(map[string]string)

	for _, file := range files {
		options["year"] = file.year
		result, err := tmdbAPI.SearchMovie(file.name, options)
		if err != nil {
			log.Fatal(err)
		}
		if result.TotalResults >= 1 {
			posterFile := download(imageBaseURL + result.Results[0].PosterPath)
			defer os.Remove(posterFile.Name()) // clean up
			log.Println("Saving cover for " + file.path)
			file.attachCover(posterFile)
		}

	}
}
