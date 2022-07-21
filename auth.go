package geekdo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type loginBody struct {
	Credentials loginBodyCredentials `json:"credentials"`
}

type loginBodyCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (c *Client) Login(username, password string) error {
	body := loginBody{
		Credentials: loginBodyCredentials{
			Username: username,
			Password: password,
		},
	}
	bodyJSON, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("failed to encode body: %w", err)
	}
	resp, err := http.Post("https://boardgamegeek.com/login/api/v1", "application/json", bytes.NewBuffer(bodyJSON))
	if err != nil {
		return fmt.Errorf("failed to log in: %w", err)
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("failed to log in: status code %v", resp.StatusCode)
	}
	// Make sure we persist the cookies, as there is some funk in the BGG
	// cookies causing them to be deleted immediately after being set due to
	// duplicates.
	newCookies := []*http.Cookie{}
	for _, c := range resp.Cookies() {
		if c.Expires.After(time.Now()) {
			newCookies = append(newCookies, c)
		}
	}
	u, err := url.Parse("https://boardgamegeek.com/")
	if err != nil {
		return fmt.Errorf("failed to parse URL: %w", err)
	}
	c.httpClient.Jar.SetCookies(u, newCookies)
	return nil
}
