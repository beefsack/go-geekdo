package geekdo

import (
	"bytes"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
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

func singleContent(n *html.Node, s string) (string, error) {
	nodes, err := htmlquery.QueryAll(n, s)
	if err != nil {
		return "", err
	}
	if len(nodes) != 1 {
		return "", fmt.Errorf(
			"could not find single node with search '%s' in %s",
			s,
			htmlquery.InnerText(n),
		)
	}
	return htmlquery.InnerText(nodes[0]), nil
}

// ParseAdvSearch parses the items out of an advanced search.
func ParseAdvSearch(page []byte) ([]SearchCollectionItem, error) {
	doc, err := htmlquery.Parse(bytes.NewReader(page))
	if err != nil {
		return nil, err
	}
	rows, err := htmlquery.QueryAll(doc, "//tr[@id='row_']")
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
		// Thumbnail, may be missing for a game
		if img := htmlquery.FindOne(r, "td[@class='collection_thumbnail']//img"); img != nil {
			item.Thumbnail = htmlquery.SelectAttr(img, "src")
		}
		// Name and URL from the same link element
		aNode, err := htmlquery.Query(r,
			"td[starts-with(@id, 'CEcell_objectname')]/div[starts-with(@id, 'results_objectname')]/a")
		if err != nil {
			return nil, err
		}
		if aNode == nil {
			return nil, fmt.Errorf("could not find name link in row")
		}
		item.Name = htmlquery.InnerText(aNode)
		item.URL = htmlquery.SelectAttr(aNode, "href")
		// ID and type (from URL)
		submatches := urlPartsRegexp.FindStringSubmatch(item.URL)
		if submatches == nil {
			return nil, fmt.Errorf("could not find type and ID in %s", item.URL)
		}
		item.Type = submatches[1]
		item.ID, err = strconv.Atoi(submatches[2])
		if err != nil {
			return nil, err
		}
		// Year
		years, err := htmlquery.QueryAll(r,
			"td[starts-with(@id, 'CEcell_objectname')]//span[@class='smallerfont dull']")
		if err != nil {
			return nil, err
		}
		if len(years) > 0 {
			item.Year, err = strconv.Atoi(strings.TrimFunc(
				htmlquery.InnerText(years[0]),
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
		ratings, err := htmlquery.QueryAll(r, "td[@class='collection_bggrating']")
		if err != nil {
			return nil, err
		}
		if len(ratings) != 3 {
			return nil, fmt.Errorf("expected to find 3 ratings in %s", htmlquery.InnerText(r))
		}
		if content := strings.TrimSpace(htmlquery.InnerText(ratings[0])); content != "N/A" {
			item.BayesAverage, err = strconv.ParseFloat(content, 64)
			if err != nil {
				return nil, err
			}
		}
		if content := strings.TrimSpace(htmlquery.InnerText(ratings[1])); content != "N/A" {
			item.Average, err = strconv.ParseFloat(content, 64)
			if err != nil {
				return nil, err
			}
		}
		if content := strings.TrimSpace(htmlquery.InnerText(ratings[2])); content != "N/A" {
			item.UsersRated, err = strconv.Atoi(content)
			if err != nil {
				return nil, err
			}
		}
		items = append(items, item)
	}
	return items, nil
}
