package movies

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"gonchill/prompt"
	"gonchill/scripts"
	"gonchill/util"
)

type State struct {
	URL     string
	Content string
}

func SearchMovies(query string, option string) {
	stack := []State{}
	stack = append(stack, State{URL: buildSearchURL(query), Content: "movies"})

	cookiesFilePath := filepath.Join("/tmp", "cookies.json")
	cookies, err := scripts.ReadCookies(cookiesFilePath)
	if err != nil {
		log.Fatalf("Error reading cookies: %s", err)
	}

	for len(stack) > 0 {
		currentState := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		switch currentState.Content {
		case "movies":
			selected_movie := getMovies(currentState.URL, cookies)
			if selected_movie == "Go Back" {
				continue
			}
			stack = append(stack, currentState, State{URL: selected_movie, Content: "magnets"})

		case "magnets":
			selected_magnet := getMagnets(currentState.URL, cookies)
			if selected_magnet == "Go Back" {
				continue
			}
			util.Watch(selected_magnet, option)
			return
		}
	}
}

func buildSearchURL(query string) string {
	encodedQuery := url.QueryEscape(query)
	return fmt.Sprintf("https://en.rarbg-official.com/movies?keyword=%s&quality=&genre=&rating=0&year=0&language=&order_by=latest", encodedQuery)
}

func showPrompt(selections []string, content string) string {
	options := append([]string{"Go Back"}, selections...)
	return prompt.Selection(options, content)
}

func getMovies(searchURL string, cookies []*http.Cookie) string {
	client := &http.Client{}
	req, err := http.NewRequest("GET", searchURL, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36")

	for _, cookie := range cookies {
		req.AddCookie(cookie)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Failed to fetch data: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Server returned non-200 status code: %d", resp.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatalf("Failed to parse HTML: %v", err)
	}

	var hold []string
	doc.Find("a[href]").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if exists && strings.Contains(href, "https://en.rarbg-official.com/movies/") {
			hold = append(hold, href)
		}
	})

	selected_movie := showPrompt(util.RemoveDuplicates(hold), "movies")
	return selected_movie
}

func getMagnets(selected_movie string, cookies []*http.Cookie) string {
	client := &http.Client{}
	req, err := http.NewRequest("GET", selected_movie, nil)
	if err != nil {
		log.Fatalf("Failed to create a New Request")
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36")

	for _, cookie := range cookies {
		req.AddCookie(cookie)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Failed to reach selected movie: %s", selected_movie)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatalf("Failed to parse HTML: %v", err)
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

	selected_magnet := showPrompt(util.RemoveDuplicates(magnetLinks), "magnets")
	return selected_magnet
}
