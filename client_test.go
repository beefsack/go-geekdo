package geekdo

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_AdvSearch(t *testing.T) {
	c := NewClient()
	res, err := c.AdvSearch("https://boardgamegeek.com/browse/boardgame/page/1?sort=rank&sortdir=asc")
	assert.NoError(t, err)
	fmt.Printf("%#v", res)
	assert.NotEmpty(t, res)
	assert.True(t, false)
}
