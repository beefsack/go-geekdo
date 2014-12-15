package geekdo

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var graphRatingsRegexp = regexp.MustCompile("chd=t:([0-9.,]+)")

// ParseGraph parses a BGG ratings graph to get rating percentages.
func ParseGraph(page []byte) (map[int]float64, error) {
	submatches := graphRatingsRegexp.FindStringSubmatch(string(page))
	if submatches == nil {
		return nil, fmt.Errorf("could not find graphs in %s", page)
	}
	ratings := map[int]float64{}
	rating := 1
	for _, r := range strings.Split(submatches[1], ",") {
		perc, err := strconv.ParseFloat(r, 64)
		if err != nil {
			return nil, err
		}
		ratings[rating] = perc
		rating++
	}
	return ratings, nil
}
