package geekdo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_AdvSearch(t *testing.T) {
	c, err := NewClient()
	assert.NoError(t, err)
	res, err := c.AdvSearch("https://boardgamegeek.com/browse/boardgame/page/1?sort=rank&sortdir=asc")
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
	// Check that all the games have ranks
	for _, g := range res {
		assert.NotZero(t, g.Rank)
	}
}
