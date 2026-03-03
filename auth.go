package geekdo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
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
	req, err := http.NewRequest("POST", "https://boardgamegeek.com/login/api/v1", bytes.NewBuffer(bodyJSON))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	c.setBaseHeaders(req)
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Origin", "https://boardgamegeek.com")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Priority", "u=4")
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to log in: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("failed to log in: status code %v", resp.StatusCode)
	}
	// BGG's login response sets bggusername/bggpassword as host-only cookies,
	// then immediately sends deletion cookies (Max-Age=0) with domain=.boardgamegeek.com.
	// Go's cookiejar conflates the two domain scopes within the same registrable domain,
	// so the deletions wipe the valid cookies. Re-add non-deletion cookies to restore them.
	// In Go's cookie parsing, Max-Age=0 becomes MaxAge=-1; real cookies have MaxAge >= 0.
	var keep []*http.Cookie
	for _, ck := range resp.Cookies() {
		if ck.MaxAge >= 0 {
			keep = append(keep, ck)
		}
	}
	c.httpClient.Jar.SetCookies(req.URL, keep)
	return nil
}
