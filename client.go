package geekdo

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type Client struct {
	httpClient *http.Client
}

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

func (c *Client) getXml(path string) (io.ReadCloser, error) {
	resp, err := c.request(path)
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}

func (c *Client) decode(path string, v interface{}) error {
	body, err := c.getXml(path)
	if err != nil {
		return err
	}
	defer body.Close()
	return xml.NewDecoder(body).Decode(v)
}

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

func (c *Client) Thing(kind string, id int, options ThingOptions) (Things, error) {
	return c.Things([]string{kind}, []int{id}, options)
}

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

type SearchOptions struct {
	Kinds []string
	Exact bool
}

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
