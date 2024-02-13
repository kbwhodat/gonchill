package movies

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"gonchill/prompt"
	"gonchill/util"
)

func SearchMovies(query string) {

	encodedQuery := url.QueryEscape(query)

	searchURL := fmt.Sprintf("https://en.rarbg-official.com/movies?keyword=%s&quality=&genre=&rating=0&year=0&language=&order_by=latest", encodedQuery)

	resp, err := http.Get(searchURL)
	if err != nil {
		log.Fatalf("Failed to fetch data: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Server returned non-200 status code: %d", resp.StatusCode)
	}

	// Parse the HTML content
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatalf("Failed to parse HTML: %v", err)
	}

	var hold []string
	// Extract and print movie URLs
	fmt.Println("Movies found:")
	doc.Find("a[href]").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if exists && strings.Contains(href, "https://en.rarbg-official.com/movies/") {
			hold = append(hold, href)
		}
	})

	selected_movie := prompt.Selection(util.RemoveDuplicates(hold), "movies")

	resp, err = http.Get(selected_movie)
	if err != nil {
		log.Fatalf("Failed to reach selected movie: %s", selected_movie)
	}
	defer resp.Body.Close()

	doc, err = goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatalf("Failed to parse html: %v", err)
	}

	var magnetLinks []string
	pattern := regexp.MustCompile(`magnet:\?xt=urn:btih:[a-zA-Z0-9]*`)
	doc.Find("a[href]").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if exists && pattern.MatchString(href) {
			magnet := pattern.FindString(href)
			magnetLinks = append(magnetLinks, magnet)
		}
	})

	selected_magnet := prompt.Selection(util.RemoveDuplicates(magnetLinks), "magnets")

	util.Watch(selected_magnet)

}


