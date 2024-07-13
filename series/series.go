package series

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
	Doc     *goquery.Document
}

func SearchSeries(query string, option string) {
	stack := []State{}
	stack = append(stack, State{URL: buildSearchURL(query), Content: "series"})

	for len(stack) > 0 {
		currentState := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		switch currentState.Content {
		case "series":
			selected_series := getSeries(currentState.URL)
			if selected_series == "Go Back" {
				continue
			} 
			stack = append(stack, currentState, State{URL: selected_series, Content: "seasons"})

		case "seasons":
			doc, selected_season := getSeasons(currentState.URL)
			if selected_season == "Go Back" {
				continue
			}
			stack = append(stack, currentState, State{URL: transformSeason(selected_season), Content: "episodes", Doc: doc})

		case "episodes":
			selected_episode := getEpisodes(currentState.URL, currentState.Doc)
			if selected_episode == "Go Back" {
				continue
			}
			stack = append(stack, currentState, State{URL: selected_episode, Content: "magnets"})

		case "magnets":
			selected_magnet := getMagnets(currentState.URL)
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
	return fmt.Sprintf("https://en.rarbg-official.com/series?keyword=%s&genre=&rating=0&order_by=latest", encodedQuery)
}

func showPrompt(selections []string, content string) string {
	options := append([]string{"Go Back"}, selections...)
	return prompt.Selection(options, content)
}

func getSeries(searchURL string) string {
	cookiesFilePath := filepath.Join("/tmp", "cookies.json")
	cookies, _ := scripts.ReadCookies(cookiesFilePath)

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
		log.Fatalf("Error fetching series: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Server returned non-200 status code: %d", resp.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatalf("Unable to parse HTML: %v", err)
	}

	var hold []string
	doc.Find("a[href]").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if exists && strings.Contains(href, "https://en.rarbg-official.com/series/") {
			hold = append(hold, href)
		}
	})

	selected_series := showPrompt(util.RemoveDuplicates(hold), "series")
	return selected_series
}

func getSeasons(selected_series string) (*goquery.Document, string) {
	cookiesFilePath := filepath.Join("/tmp", "cookies.json")
	cookies, _ := scripts.ReadCookies(cookiesFilePath)

	client := &http.Client{}
	req, err := http.NewRequest("GET", selected_series, nil)
	if err != nil {
		log.Fatalf("Failed to create a New Request")
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36")

	for _, cookie := range cookies {
		req.AddCookie(cookie)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Unable to search selected series: %s", selected_series)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatalf("Unable to parse document: %v", err)
	}

	var seasons []string
	doc.Find("a.season-item").Each(func(i int, s *goquery.Selection) {
		s.Contents().Each(func(_ int, sel *goquery.Selection) {
			if goquery.NodeName(sel) == "#text" {
				text := strings.TrimSpace(sel.Text())
				if text != "" {
					seasons = append(seasons, text)
				}
			}
		})
	})

	selected_season := showPrompt(util.RemoveDuplicates(seasons), "seasons")
	return doc, selected_season
}

func getEpisodes(selected_season string, doc *goquery.Document) string {
	var episodes []string
	doc.Find("a[href]").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if exists && strings.Contains(href, transformSeason(selected_season)) {
			episodes = append(episodes, href)
		}
	})

	selected_episode := showPrompt(util.RemoveDuplicates(episodes), "episodes")
	return selected_episode
}

func getMagnets(selected_episode string) string {
	cookiesFilePath := filepath.Join("/tmp", "cookies.json")
	cookies, _ := scripts.ReadCookies(cookiesFilePath)
	client := &http.Client{}
	req, err := http.NewRequest("GET", selected_episode, nil)
	if err != nil {
		log.Fatalf("Failed to create a New Request")
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36")

	for _, cookie := range cookies {
		req.AddCookie(cookie)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Unable to fetch episode: %v", err)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatalf("Unable to parse HTML: %v", err)
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

func transformSeason(season string) string {
	lowercase := strings.ToLower(season)
	result := strings.Replace(lowercase, " ", "-", 1)

	if season == "Specials" {
		result = "season-0"
	}

	return result
}
