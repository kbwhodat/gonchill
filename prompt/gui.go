package prompt

import (
    "fmt"
    "strings"
    "github.com/AlecAivazis/survey/v2"
)

type link struct {
    Name string
    Url  string
}

func Selection(selections []string, content string) string {
    options := selections

    links := []link{}
    var prefix string

    switch content {
    case "episodes":
        prefix = "https://en.rarbg-official.com/episodes/"
    case "seasons":
        prefix = "https://en.rarbg-official.com/seasons/"
    case "series":
        prefix = "https://en.rarbg-official.com/series/"
    case "movies":
        prefix = "https://en.rarbg-official.com/movies/"
    }

    for _, option := range options {
        endpoint := strings.TrimPrefix(option, prefix)
        newLinks := link{Name: endpoint, Url: option}
        links = append(links, newLinks)
    }

    var shortName []string
    for _, p := range links {
        shortName = append(shortName, p.Name)
    }


    var selected string
    prompt := &survey.Select{
        Message: "Please select an option:",
        Options: shortName,
    }

    err := survey.AskOne(prompt, &selected)
    if err != nil {
        fmt.Printf("Prompt failed %v\n", err)
        return "Something went wrong..."
    }

    fmt.Printf("You selected: %s\n", selected)

    index := -1
    for i, name := range shortName {
        if name == selected {
            index = i
            break
        }
    }

    if index == -1 {
        fmt.Println("Selection not found")
        return "Something went wrong..."
    }

    return links[index].Url
}
