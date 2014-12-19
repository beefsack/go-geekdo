package geekdo

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// Client is a client of a geekdo site.
type Client struct {
	httpClient *http.Client
}

// NewClient creates a Client.
func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{},
	}
}

func (c *Client) request(path string) (*http.Response, error) {
	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf("http://www.boardgamegeek.com/xmlapi2%s", path),
		nil,
	)
	if err != nil {
		return nil, err
	}
	return c.httpClient.Do(req)
}

func (c *Client) getXML(path string) (io.ReadCloser, error) {
	resp, err := c.request(path)
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}

func (c *Client) decode(path string, v interface{}) error {
	body, err := c.getXML(path)
	if err != nil {
		return err
	}
	defer body.Close()
	return xml.NewDecoder(body).Decode(v)
}

// ThingOptions specify options used when querying things.
type ThingOptions struct {
	Versions       bool
	Videos         bool
	Stats          bool
	Historical     bool
	Marketplace    bool
	Comments       bool
	RatingComments bool
	Page           int
	PageSize       int
	From           time.Time
	To             time.Time
}

// Thing queries a thing of a kind and an id.
func (c *Client) Thing(kind string, id int, options ThingOptions) (Things, error) {
	return c.Things([]string{kind}, []int{id}, options)
}

// Things query things given kinds and ids.
func (c *Client) Things(kinds []string, ids []int, options ThingOptions) (Things, error) {
	things := Things{}
	if kinds == nil || len(kinds) == 0 {
		return things, errors.New("kinds must include at least one item")
	}
	if ids == nil || len(ids) == 0 {
		return things, errors.New("ids must include at least one item")
	}
	query := url.Values{}
	query.Set("type", strings.Join(kinds, ","))
	idStrs := make([]string, len(ids))
	for i, id := range ids {
		idStrs[i] = strconv.Itoa(id)
	}
	query.Set("id", strings.Join(idStrs, ","))
	if options.Versions {
		query.Set("versions", "1")
	}
	if options.Videos {
		query.Set("videos", "1")
	}
	if options.Stats {
		query.Set("stats", "1")
	}
	if options.Historical {
		query.Set("historical", "1")
	}
	if options.Marketplace {
		query.Set("marketplace", "1")
	}
	if options.Comments {
		query.Set("comments", "1")
	}
	if options.RatingComments {
		query.Set("ratingcomments", "1")
	}
	if options.Page > 0 {
		query.Set("page", strconv.Itoa(options.Page))
	}
	if options.PageSize > 0 {
		query.Set("pagesize", strconv.Itoa(options.PageSize))
	}
	if !options.From.IsZero() {
		query.Set("from", options.From.Format("2006-01-02"))
	}
	if !options.To.IsZero() {
		query.Set("to", options.To.Format("2006-01-02"))
	}
	err := c.decode(fmt.Sprintf("/thing?%s", query.Encode()), &things)
	return things, err
}

// SearchOptions are the options used when searching.
type SearchOptions struct {
	Kinds []string
	Exact bool
}

// Search searches Geekdo given a query and options.
func (c *Client) Search(query string, options SearchOptions) (Things, error) {
	things := Things{}
	queryParams := url.Values{}
	queryParams.Set("query", query)
	if options.Kinds != nil && len(options.Kinds) > 0 {
		queryParams.Set("type", strings.Join(options.Kinds, ","))
	}
	if options.Exact {
		queryParams.Set("exact", "1")
	}
	err := c.decode(fmt.Sprintf("/search?%s", queryParams.Encode()), &things)
	return things, err
}

// AdvSearch scrapes an advanced search page for more powerful searching and
// ordering.
func (c *Client) AdvSearch(url string) ([]CollectionItem, error) {
	resp, err := c.httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return ParseAdvSearch(body)
}

// Ratings gets the rating counts (ratings from 1 to 10) for a thing.
func (c *Client) Ratings(id int) (map[int]int, error) {
	resp, err := c.httpClient.Get(fmt.Sprintf(
		"http://boardgamegeek.com/graphstats/thing/%d?ajax=1",
		id,
	))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return ParseGraph(body)
}
