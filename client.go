package openbd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// Client for openbd.
type Client struct {
	URL        *url.URL
	HTTPClient *http.Client
}

// NewClientV1 generates a client for openbd.
func NewClientV1(host string, cli *http.Client) (Fetcher, error) {
	u, err := url.ParseRequestURI(host)
	if err != nil {
		return nil, err
	}
	u.Path = "/v1"
	if cli == nil {
		cli = http.DefaultClient
	}
	return Fetcher(&Client{
		URL:        u,
		HTTPClient: cli,
	}), nil
}

// Get book information by isbn.
func (c *Client) Get(isbn ...string) (map[string]BookInformer, error) {
	if len(isbn) > 1e3 {
		return nil, fmt.Errorf("opendb: args length should be under 1,000 - length = %d", len(isbn))
	}
	resp, err := c.HTTPClient.Get(c.URL.String() + "/get?isbn=" + strings.Join(isbn, ","))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("opendb: faild to request - status code = %d", resp.StatusCode)
	}
	rs := make([]*Root, 0, len(isbn))
	if err := json.NewDecoder(resp.Body).Decode(&rs); err != nil {
		return nil, err
	}
	m := make(map[string]BookInformer, len(rs))
	for i := range rs {
		if rs[i] != nil {
			m[rs[i].Summary.Isbn] = BookInformer(rs[i])
		}
	}
	return m, nil
}
