package geekdo

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestParseAdvSearch(t *testing.T) {
	items, err := ParseAdvSearch(advSearchPage)
	assert.NoError(t, err)
	assert.Equal(t, advSearchExpected, items)
}
