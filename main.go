package main

import (
	"fmt"
	"os"
	"strings"

	"gonchill/movies"
	"gonchill/series"
)

func main() {

	if len(os.Args) < 3 {
		fmt.Println("Usage: movies|series [Options] \"search query\"")
		fmt.Println("Options: -v: To use vlc ")
		fmt.Println("         -m: To use mpv ")
		return
	}

	category := os.Args[1]
	option := os.Args[2]
	query := strings.Join(os.Args[3:], " ")

	switch category {
	case "movies":
		movies.SearchMovies(query, option)
	case "series":
		series.SearchSeries(query, option)
	default:
		fmt.Println("Invalid category. Use 'movies' or 'series'.")
	}

}

