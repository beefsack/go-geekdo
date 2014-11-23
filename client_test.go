package bgg

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_request(t *testing.T) {
	c := NewClient()
	v := struct {
		Items []string `xml:"items"`
	}{}
	assert.NoError(t, c.decode("/thing?type=boardgame&id=154203", &v))
	fmt.Printf("%#v\n", v)
}
