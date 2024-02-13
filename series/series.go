package series

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

func SearchSeries(query string) {

	encodedQuery := url.QueryEscape(query)

	searchURL := fmt.Sprintf("https://en.rarbg-official.com/series?keyword=%s&genre=&rating=0&order_by=latest", encodedQuery)

	resp, err := http.Get(searchURL)
	if err != nil {
		log.Fatalf("Error search for this url: %v", err)
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

	selected_series := prompt.Selection(util.RemoveDuplicates(hold), "series")


	resp, err = http.Get(selected_series)
	if err != nil {
		log.Fatalf("Unable to search selected series: %s", selected_series)
	}
	defer resp.Body.Close()

	doc, err = goquery.NewDocumentFromReader(resp.Body)
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

	selected_season := prompt.Selection(util.RemoveDuplicates(seasons), "seasons")
	fmt.Println(transformSeason(selected_season))

	var episodes []string
	doc.Find("a[href]").Each(func(i int, s *goquery.Selection){
		href, exists := s.Attr("href")
		if exists && strings.Contains(href, transformSeason(selected_season)){
			episodes = append(episodes, href)

		}
	})


	selected_episode := prompt.Selection(util.RemoveDuplicates(episodes), "episodes")

	resp, err = http.Get(selected_episode)
	if err != nil {
		log.Fatalf("Unable to fetch episode: %v", err)
	}
	defer resp.Body.Close()

	doc, err = goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatalf("Unable to parse html: %v", err)
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

func transformSeason(season string) string {

	lowercase := strings.ToLower(season)
	result := strings.Replace(lowercase, " ", "-", 1)

	return result
}
