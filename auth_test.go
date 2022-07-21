package geekdo

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_Login(t *testing.T) {
	username := os.Getenv("BGG_USERNAME")
	password := os.Getenv("BGG_PASSWORD")
	if username == "" || password == "" {
		t.Skip("Need to pass BGG_USERNAME and BGG_PASSWORD env vars to run login test")
		return
	}
	c, err := NewClient()
	assert.NoError(t, err)
	err = c.Login(username, password)
	assert.NoError(t, err)
	res, err := c.AdvSearch("https://boardgamegeek.com/browse/boardgame/page/21?1&sort=rank&sortdir=asc")
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}
