package instagram

import (
	"errors"
	"fmt"
	"net/url"
)

// Get basic information about a user. To get information
// about the owner of the access token, you can use self instead of the user-id.
//
// Instagram API docs: https://instagram.com/developer/endpoints/users/#get_users
func (c *Client) GetUserInfo(userID string) (*UserResponse, error) {
	if userID == "" {
		return nil, errors.New("set user ID or \"self\"")
	}

	res := &UserResponse{}
	err := c.doRequest("GET", fmt.Sprintf("users/%v", userID), nil, res)
	return res, err
}

// See the authenticated user's feed.
//
// COUNT	Count of media to return.
// MIN_ID	Return media later than this min_id.
// MAX_ID	Return media earlier than this max_id.
//
// Instagram API docs: https://instagram.com/developer/endpoints/users/#get_users_feed
func (c *Client) GetSelfFeed(params url.Values) (*MediaResponse, error) {
	res := &MediaResponse{}
	err := c.doRequest("GET", "users/self/feed", params, res)
	return res, err
}

// Get the most recent media published by a user. To get the most recent media
// published by the owner of the access token, you can use self instead of the user-id.
//
// COUNT			Count of media to return.
// MAX_TIMESTAMP	Return media before this UNIX timestamp.
// MIN_TIMESTAMP	Return media after this UNIX timestamp.
// MIN_ID			Return media later than this min_id.
// MAX_ID			Return media earlier than this max_id.
//
// Instagram API docs: https://instagram.com/developer/endpoints/users/#get_users_media_recent
func (c *Client) GetRecentMedia(userID string, params url.Values) (*MediaResponse, error) {
	if userID == "" {
		return nil, errors.New("set user ID or \"self\"")
	}

	res := &MediaResponse{}
	err := c.doRequest("GET", fmt.Sprintf("users/%v/media/recent", userID), params, res)
	return res, err
}
