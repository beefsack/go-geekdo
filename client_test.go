package geekdo

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
)

func TestClient_request(t *testing.T) {
	c := NewClient()
	things, err := c.Search("Imperial", SearchOptions{})
	assert.NoError(t, err)
	spew.Dump(things)
}
