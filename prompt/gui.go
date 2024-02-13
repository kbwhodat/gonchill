package prompt

import (
	"fmt"
	"strings"
	"github.com/manifoldco/promptui"
)

type link struct {
	Name string
	Url  string
}

func Selection(selections []string, content string) string {
	options := selections

	links := []link{}
	var prefix string

	if content == "episodes" {
		prefix = "https://en.rarbg-official.com/episodes/"

	}else if content == "seasons" {
		prefix = "https://en.rarbg-official.com/seasons/"

	}else if content == "series" {
		prefix = "https://en.rarbg-official.com/series/"

	}else if content == "movies" {
		prefix = "https://en.rarbg-official.com/movies/"
	}

	for i := 0; i < len(options); i++ {
		endpoint := strings.TrimPrefix(options[i], prefix)
		newLinks := link{Name: endpoint, Url: options[i]}
		links = append(links, newLinks)

	}

	var shortName []string
	for _, p := range links {
		shortName = append(shortName, p.Name)
	}

	prompt := promptui.Select{
		Label: "Please select an option",
		Items: shortName,
	}

	index, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return "Something went wrong..."
	}

	fmt.Printf("You selected: %s\n", result)

	return links[index].Url
}
