package geekdo

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/moovweb/gokogiri"
	"github.com/moovweb/gokogiri/xml"
)

var urlPartsRegexp = regexp.MustCompile(`^/(.+)/(\d+)/`)

// SearchCollectionItem is a row in an advanced search result.
type SearchCollectionItem struct {
	ID           int
	Type         string
	Rank         int
	Thumbnail    string
	Name         string
	URL          string
	Year         int
	BayesAverage float64
	Average      float64
	UsersRated   int
}

func singleContent(n xml.Node, s string) (string, error) {
	nodes, err := n.Search(s)
	if err != nil {
		return "", err
	}
	if len(nodes) != 1 {
		return "", fmt.Errorf(
			"could not find single node with search '%s' in %s",
			s,
			n.Content(),
		)
	}
	return nodes[0].Content(), nil
}

// ParseAdvSearch parses the items out of an advanced search.
func ParseAdvSearch(page []byte) ([]SearchCollectionItem, error) {
	doc, err := gokogiri.ParseHtml(page)
	if err != nil {
		return nil, err
	}
	rows, err := doc.Search("//table/tr[@id='row_']")
	if err != nil {
		return nil, err
	}
	items := []SearchCollectionItem{}
	for _, r := range rows {
		item := SearchCollectionItem{}
		// Rank
		rank, err := singleContent(r, "td[@class='collection_rank']")
		if err != nil {
			return nil, err
		}
		if content := strings.TrimSpace(rank); content != "N/A" {
			item.Rank, err = strconv.Atoi(content)
			if err != nil {
				return nil, err
			}
		}
		// Thumbnail
		item.Thumbnail, err = singleContent(r,
			"td[@class='collection_thumbnail']//img/@src")
		if err != nil {
			return nil, err
		}
		// Name
		item.Name, err = singleContent(r,
			"td[starts-with(@id, 'CEcell_objectname')]/div[starts-with(@id, 'results_objectname')]/a")
		if err != nil {
			return nil, err
		}
		// URL
		item.URL, err = singleContent(r,
			"td[starts-with(@id, 'CEcell_objectname')]/div[starts-with(@id, 'results_objectname')]/a/@href")
		if err != nil {
			return nil, err
		}
		// ID and type (from URL)
		submatches := urlPartsRegexp.FindStringSubmatch(item.URL)
		if submatches == nil {
			return nil, fmt.Errorf("could not find type and ID in %s", item.URL)
		}
		item.Type = submatches[1]
		item.ID, err = strconv.Atoi(submatches[2])
		// Year
		years, err := r.Search(
			"td[starts-with(@id, 'CEcell_objectname')]//span[@class='smallerfont dull']")
		if err != nil {
			return nil, err
		}
		if len(years) > 0 {
			item.Year, err = strconv.Atoi(strings.TrimFunc(
				years[0].Content(),
				func(r rune) bool {
					// Trim brackets and spaces
					return r < '0' || r > '9'
				},
			))
			if err != nil {
				return nil, err
			}
		}
		// Ratings (can be N/A)
		ratings, err := r.Search("td[@class='collection_bggrating']")
		if err != nil {
			return nil, err
		}
		if len(ratings) != 3 {
			return nil, fmt.Errorf("expected to find 3 ratings in %s", r.Content())
		}
		if content := strings.TrimSpace(ratings[0].Content()); content != "N/A" {
			item.BayesAverage, err = strconv.ParseFloat(content, 64)
			if err != nil {
				return nil, err
			}
		}
		if content := strings.TrimSpace(ratings[1].Content()); content != "N/A" {
			item.Average, err = strconv.ParseFloat(content, 64)
			if err != nil {
				return nil, err
			}
		}
		if content := strings.TrimSpace(ratings[2].Content()); content != "N/A" {
			item.UsersRated, err = strconv.Atoi(content)
			if err != nil {
				return nil, err
			}
		}
		items = append(items, item)
	}
	return items, nil
}
