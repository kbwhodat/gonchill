package scripts

import (
  "encoding/json"
  "net/http"
  "fmt"
  "log"
  "time"
  "os"
)

type Cookie struct {
    Domain   string `json:"domain"`
    Name     string `json:"name"`
    Value    string `json:"value"`
    Path     string `json:"path"`
    HttpOnly bool   `json:"httpOnly"`
    Secure   bool   `json:"secure"`
    SameSite string `json:"sameSite"`
    Expiry   int64  `json:"expiry"`
}


func ReadCookies(filePath string) ([]*http.Cookie, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %v", err)
	}

	var cookies []Cookie
	err = json.Unmarshal(data, &cookies)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal cookies: %v", err)
	}

	httpCookies := make([]*http.Cookie, len(cookies))
	for i, c := range cookies {
		httpCookies[i] = &http.Cookie{
			Name:     c.Name,
			Value:    c.Value,
			Path:     c.Path,
			Domain:   c.Domain,
			Secure:   c.Secure,
			HttpOnly: c.HttpOnly,
		}
	}
	return httpCookies, nil
}

func readCookiesFromFile() ([]Cookie, error) {
    var cookies []Cookie
    data, err := os.ReadFile("/tmp/cookies.json")
    if err != nil {
        return nil, err
    }
    err = json.Unmarshal(data, &cookies)
    if err != nil {
        return nil, err
    }
    return cookies, nil
}

func CheckCookieExpiry(cookieName string) (bool) {
  cookies, err := readCookiesFromFile()
  if err != nil {
    log.Fatal(err)
  }
  for _, cookie := range cookies {
    if cookie.Name == cookieName {
      expiryTime := time.Unix(cookie.Expiry, 0).AddDate(-1, 0, 0).Unix()
      log.Println(expiryTime)
      expiryTimeAfter30Mins := time.Unix(expiryTime, 0).Add(30 * time.Minute)

      if time.Now().After(expiryTimeAfter30Mins) {
        return true
      } else {
        return false
      }
    }
  }
  return true
}
