package main

import (
	"log"
	"os"

	"github.com/magiconair/properties"
	"github.com/ryanbradynd05/go-tmdb"
)

var tmdbAPI *tmdb.TMDb
var imageBaseURL string

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

	// TODO: Add two parameters -p folder to scan -e folder to exclude
	root := `\\DS115\Movies`
	files := scanMovies(root, "Kids")

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
