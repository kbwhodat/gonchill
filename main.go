package main

import (
  "fmt"
  "bytes"
  "errors"
  "log"
  "os"
  "os/exec"
  "strings"
  "gonchill/movies"
  "gonchill/series"
  "gonchill/scripts"
  _ "embed"
)

//go:embed scripts/setcookies.py
var setcookiesPy string

func executePythonTask(script string) {
  tmpfile, err := os.CreateTemp("", "setcookies*.py")
  if err != nil {
    log.Printf("Error creating temp file: %s\n", err)
    return
  }
  defer os.Remove(tmpfile.Name())

  if _, err := tmpfile.Write([]byte(script)); err != nil {
    log.Printf("Error writing to temp file: %s\n", err)
    return
  }
  if err := tmpfile.Close(); err != nil {
    log.Printf("Error closing temp file: %s\n", err)
    return
  }

  cmd := exec.Command("python", tmpfile.Name())
  var out bytes.Buffer
  var stderr bytes.Buffer
  cmd.Stdout = &out
  cmd.Stderr = &stderr

  err = cmd.Run()
  if err != nil {
    log.Printf("cmd.Run() failed with %s\n", err)
    log.Printf("stderr: %v", stderr.String())
  }
  log.Println(out.String())
}

func main() {
  if len(os.Args) < 3 {
    fmt.Println("Usage: gonchill movies|series [Options] \"search query\"")
    fmt.Println("Options: -v: To use vlc ")
    fmt.Println("         -m: To use mpv ")
    return
  }

  category := os.Args[1]
  option := os.Args[2]
  // Fix: strings.join should be strings.Join (capital J) - Go function names are case-sensitive
  query := strings.Join(os.Args[3:], " ")

  _, err := os.Stat("/tmp/cookies.json")
  if errors.Is(err, os.ErrNotExist) {
    log.Println("No cookies file found, generating fresh cookies...")
    executePythonTask(setcookiesPy)
  }

  isCookieExpired := scripts.CheckCookieExpiry("cf_clearance")
  if isCookieExpired == true {
    log.Println("Cookies are expired...generating fresh cookies")
    log.Println("Please wait...")

    executePythonTask(setcookiesPy)

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
