package instagram

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

const (
	BaseURL = "https://api.instagram.com/v1/"
)

type Client struct {
	// Base URL for API requests.
	BaseURL string

	// Authenticated user's access_token.
	AccessToken string

	// Total number of possible Instagram API calls per hour.
	RatelimitLimit int

	// How many Instagram API calls are left for this particular access token.
	RatelimitRemaining int
}

func NewClient(accessToken string) *Client {
	client := &Client{
		BaseURL:     BaseURL,
		AccessToken: accessToken,
	}

	return client
}

// Handle request to Instagram API.
func (c *Client) doRequest(method, path string, params url.Values, r interface{}) error {
	apiURL, err := c.getCompletedApiUrl(path, params)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(method, apiURL, nil)
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(r)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return errors.New("request failed: check meta info")
	}

	err = c.getRatelimit(resp)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) getCompletedApiUrl(path string, p url.Values) (string, error) {
	if p == nil {
		p = url.Values{}
	}
	p.Set("access_token", c.AccessToken)

	apiURL := fmt.Sprintf("%v%v", c.BaseURL, path)
	u, err := url.Parse(apiURL)
	if err != nil {
		return "", err
	}

	u.RawQuery = p.Encode()
	return u.String(), nil
}

func (c *Client) getRatelimit(resp *http.Response) error {
	var err error

	c.RatelimitLimit, err = strconv.Atoi(resp.Header.Get("x-ratelimit-limit"))
	if err != nil {
		return err
	}

	c.RatelimitRemaining, err = strconv.Atoi(resp.Header.Get("x-ratelimit-remaining"))
	if err != nil {
		return err
	}

	return nil
}
