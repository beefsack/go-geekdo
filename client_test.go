package bgg

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
)

func TestClient_request(t *testing.T) {
	c := NewClient()
	v := BoardGames{}
	assert.NoError(t, c.decode("/thing?type=boardgame&id=154203&marketplace=1", &v))
	spew.Dump(v.BoardGames[0].MarketplaceListings)
}
