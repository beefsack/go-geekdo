package geekdo

import (
	"fmt"
	"regexp"
	"strconv"
)

var graphRatingsLineRegexp = regexp.MustCompile(`chm=[^\n]+`)
var graphRatingRegexp = regexp.MustCompile(`t(\d+)`)

// ParseGraph parses a BGG ratings graph to get rating counts.
func ParseGraph(page []byte) (map[int]int, error) {
	line := graphRatingsLineRegexp.FindString(string(page))
	if line == "" {
		return nil, fmt.Errorf("could not find ratings line in %s", page)
	}
	ratings := map[int]int{}
	rating := 1
	matches := graphRatingRegexp.FindAllStringSubmatch(line, -1)
	if matches == nil {
		return nil, fmt.Errorf("could not find rating values in %s", line)
	}
	for _, r := range matches {
		num, err := strconv.Atoi(r[1])
		if err != nil {
			return nil, err
		}
		ratings[rating] = num
		rating++
	}
	return ratings, nil
}
