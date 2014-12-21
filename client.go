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

// BoolOpt is an option boolean type used for querying the API.
type BoolOpt int

func (bo BoolOpt) encode() string {
	switch bo {
	case OptTrue:
		return "1"
	case OptFalse:
		return "0"
	}
	return ""
}

var (
	// OptTrue is equivalent to true.
	OptTrue = BoolOpt(1)
	// OptFalse is equivalent to false.
	OptFalse = BoolOpt(-1)
	// OptIgnore is equivalent to nil.
	OptIgnore = BoolOpt(0)
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

func (c *Client) apiGet(path string) (io.ReadCloser, error) {
	resp, err := c.get(fmt.Sprintf("http://www.boardgamegeek.com/xmlapi2%s", path))
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}

func (c *Client) apiGetDecode(path string, v interface{}) error {
	body, err := c.apiGet(path)
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
	err := c.apiGetDecode(fmt.Sprintf("/thing?%s", query.Encode()), &things)
	return things, err
}

// SearchOptions are the options used when searching.
type SearchOptions struct {
	Kinds []string
	Exact bool
}

func (c *Client) get(url string) (resp *http.Response, err error) {
	if resp, err = c.httpClient.Get(url); err != nil {
		return
	}
	if resp.StatusCode == http.StatusAccepted {
		// Request added to a queue, wait a bit and try again.
		// see http://boardgamegeek.com/thread/1188687/export-collections-has-been-updated-xmlapi-develop
		time.Sleep(time.Second)
		return c.get(url)
	}
	return
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
	err := c.apiGetDecode(fmt.Sprintf("/search?%s", queryParams.Encode()), &things)
	return things, err
}

// AdvSearch scrapes an advanced search page for more powerful searching and
// ordering.
func (c *Client) AdvSearch(url string) ([]SearchCollectionItem, error) {
	resp, err := c.get(url)
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
	resp, err := c.get(fmt.Sprintf(
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

// CollectionOptions are the options to send with a collection request.
type CollectionOptions struct {
	Version          bool
	SubType          string
	ExcludeSubType   string
	IDs              []int
	Brief            bool
	Stats            bool
	Own              BoolOpt
	Rated            BoolOpt
	Played           BoolOpt
	Comment          BoolOpt
	Trade            BoolOpt
	Want             BoolOpt
	Wishlist         BoolOpt
	WishlistPriority int
	Preordered       BoolOpt
	WantToPlay       BoolOpt
	WantToBuy        BoolOpt
	PrevOwned        BoolOpt
	HasParts         BoolOpt
	WantParts        BoolOpt
	MinBGGRating     int
	MaxBGGRating     int
	MinPlays         int
	MaxPlays         int
	ShowPrivate      bool
	CollID           int
	ModifiedSince    time.Time
}

// Collection queries a user's collection.
func (c *Client) Collection(username string, options CollectionOptions) (
	CollectionItems,
	error,
) {
	collection := CollectionItems{}
	query := url.Values{}
	query.Set("username", username)
	if options.Version {
		query.Set("version", "1")
	}
	if options.SubType != "" {
		query.Set("subtype", options.SubType)
	}
	if options.ExcludeSubType != "" {
		query.Set("excludesubtype", options.ExcludeSubType)
	}
	if options.IDs != nil && len(options.IDs) > 0 {
		idStrs := make([]string, len(options.IDs))
		for i, id := range options.IDs {
			idStrs[i] = strconv.Itoa(id)
		}
		query.Set("id", strings.Join(idStrs, ","))
	}
	if options.Brief {
		query.Set("brief", "1")
	}
	if options.Stats {
		query.Set("stats", "1")
	}
	if options.Own != OptIgnore {
		query.Set("own", options.Own.encode())
	}
	if options.Rated != OptIgnore {
		query.Set("rated", options.Rated.encode())
	}
	if options.Played != OptIgnore {
		query.Set("played", options.Played.encode())
	}
	if options.Comment != OptIgnore {
		query.Set("comment", options.Comment.encode())
	}
	if options.Trade != OptIgnore {
		query.Set("trade", options.Trade.encode())
	}
	if options.Trade != OptIgnore {
		query.Set("trade", options.Trade.encode())
	}
	if options.Wishlist != OptIgnore {
		query.Set("wishlist", options.Wishlist.encode())
	}
	if options.WishlistPriority != 0 {
		query.Set("wishlistpriority", strconv.Itoa(options.WishlistPriority))
	}
	if options.Preordered != OptIgnore {
		query.Set("preordered", options.Preordered.encode())
	}
	if options.WantToPlay != OptIgnore {
		query.Set("wanttoplay", options.WantToPlay.encode())
	}
	if options.WantToBuy != OptIgnore {
		query.Set("wanttobuy", options.WantToBuy.encode())
	}
	if options.PrevOwned != OptIgnore {
		query.Set("prevowned", options.PrevOwned.encode())
	}
	if options.HasParts != OptIgnore {
		query.Set("hasparts", options.HasParts.encode())
	}
	if options.WantParts != OptIgnore {
		query.Set("wantparts", options.WantParts.encode())
	}
	if options.MinBGGRating != 0 {
		query.Set("minbggrating", strconv.Itoa(options.MinBGGRating))
	}
	if options.MaxBGGRating != 0 {
		query.Set("bggrating", strconv.Itoa(options.MaxBGGRating))
	}
	if options.MinPlays != 0 {
		query.Set("minplays", strconv.Itoa(options.MinPlays))
	}
	if options.MaxPlays != 0 {
		query.Set("maxplays", strconv.Itoa(options.MaxPlays))
	}
	if options.ShowPrivate {
		query.Set("showprivate", "1")
	}
	if options.CollID != 0 {
		query.Set("collid", strconv.Itoa(options.CollID))
	}
	if !options.ModifiedSince.IsZero() {
		query.Set("modifiedsince", options.ModifiedSince.Format(
			"06-01-02 15:04:05"))
	}
	if err := c.apiGetDecode(
		fmt.Sprintf("/collection?%s", query.Encode()),
		&collection,
	); err != nil {
		return collection, err
	}
	return collection, nil
}
