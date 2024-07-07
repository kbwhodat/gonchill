package main

import (
	"fmt"
  "bytes"
	"os"
  "log"
  "errors"
  "os/exec"
	"strings"

	"gonchill/movies"
	"gonchill/series"
	"gonchill/scripts"
)

func executePythonTask() {
  cmd := exec.Command("python", "scripts/setcookies.py")

  var out bytes.Buffer
  var stderr bytes.Buffer

  cmd.Stdout = &out
  cmd.Stderr = &stderr

  err := cmd.Run()
  if err != nil {
    log.Printf("cmd.Run() failed with %s\n", err)
    log.Printf("stderr: %v", stderr.String())
  }
}

func main() {

	if len(os.Args) < 3 {
		fmt.Println("Usage: go run . movies|series [Options] \"search query\"")
		fmt.Println("Options: -v: To use vlc ")
		fmt.Println("         -m: To use mpv ")
		return
	}

	category := os.Args[1]
	option := os.Args[2]
	query := strings.Join(os.Args[3:], " ")

  _, err := os.Stat("scripts/cookies.json")
  if errors.Is(err, os.ErrNotExist) {
    log.Println("No cookies file found, generating fresh cookies...")
    executePythonTask()
  }
  
  isCookieExpired := scripts.CheckCookieExpiry("cf_clearance")
  if isCookieExpired == true {
    log.Println("Cookies are expired...generating fresh cookies")
    log.Println("Please wait...")

    executePythonTask()

  } else {
    log.Println("using current cookies...")
  }

	switch category {
	case "movies":
		movies.SearchMovies(query, option)
	case "series":
		series.SearchSeries(query, option)
	default:
		fmt.Println("Invalid category. Use 'movies' or 'series'.")
	}

}

