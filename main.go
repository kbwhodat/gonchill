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
		fmt.Println("Usage: movies|series \"search query\"")
		return
	}

	category := os.Args[1]
	query := strings.Join(os.Args[2:], " ")

	switch category {
	case "movies":
		movies.SearchMovies(query)
	case "series":
		series.SearchSeries(query)
	default:
		fmt.Println("Invalid category. Use 'movies' or 'series'.")
	}

}

