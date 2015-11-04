package instagram

import "fmt"

// Get the list of users this user follows.
//
// Instagram API docs: https://instagram.com/developer/endpoints/relationships/#get_users_follows
func (c *Client) GetWhoUserFollows(userID string) (*RelationshipResponse, error) {
	res := &RelationshipResponse{}
	err := c.doRequest("GET", fmt.Sprintf("users/%v/follows", userID), nil, res)
	return res, err
}

// Get the list of users this user is followed by.
//
// Instagram API docs: https://instagram.com/developer/endpoints/relationships/#get_users_followed_by
func (c *Client) GetWhoFollowsUser(userID string) (*RelationshipResponse, error) {
	res := &RelationshipResponse{}
	err := c.doRequest("GET", fmt.Sprintf("users/%v/followed-by", userID), nil, res)
	return res, err
}
