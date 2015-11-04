package instagram

type MediaResponse struct {
	Pagination *Pagination `json:"pagination,omitempty"`
	Meta       *Meta       `json:"meta,omitempty"`
	Data       []Media     `json:"data,omitempty"`
}

type UserResponse struct {
	Meta *Meta `json:"meta,omitempty"`
	Data User  `json:"data,omitempty"`
}

type RelationshipResponse struct {
	Pagination *Pagination `json:"pagination,omitempty"`
	Meta       *Meta       `json:"meta,omitempty"`
	Data       []User      `json:"data,omitempty"`
}

type Pagination struct {
	NextURL   string `json:"next_url,omitempty"`
	NextMaxID string `json:"next_max_id,omitempty"`
}

type Meta struct {
	Code         int    `json:"code,omitempty"`
	ErrorType    string `json:"error_type,omitempty"`
	ErrorMessage string `json:"error_message,omitempty"`
}

type Media struct {
	ID           string      `json:"id,omitempty"`
	Filter       string      `json:"filter,omitempty"`
	Attribution  interface{} `json:"attribution,omitempty"`
	Caption      Caption     `json:"caption,omitempty"`
	Comments     Comments    `json:"comments,omitempty"`
	CreatedTime  string      `json:"created_time,omitempty"`
	Images       Images      `json:"images,omitempty"`
	Likes        Likes       `json:"likes,omitempty"`
	Link         string      `json:"link,omitempty"`
	Location     Location    `json:"location,omitempty"`
	Tags         []string    `json:"tags,omitempty"`
	Type         string      `json:"type,omitempty"`
	User         User        `json:"user,omitempty"`
	UserHasLiked bool        `json:"user_has_liked,omitempty"`
	// UsersInPhoto []interface{} `json:"users_in_photo,omitempty"`
}

type Likes struct {
	Count int    `json:"count,omitempty"`
	Data  []User `json:"data,omitempty"`
}

type Comments struct {
	Count int       `json:"count,omitempty"`
	Data  []Comment `json:"data,omitempty"`
}

type Images struct {
	LowResolution      Image `json:"low_resolution,omitempty"`
	StandardResolution Image `json:"standard_resolution,omitempty"`
	Thumbnail          Image `json:"thumbnail,omitempty"`
}

type Image struct {
	Height int    `json:"height,omitempty"`
	URL    string `json:"url,omitempty"`
	Width  int    `json:"width,omitempty"`
}

type Location struct {
	ID        int     `json:"id,omitempty"`
	Latitude  float64 `json:"latitude,omitempty"`
	Longitude float64 `json:"longitude,omitempty"`
	Name      string  `json:"name,omitempty"`
}

type Caption struct {
	CreatedTime string `json:"created_time,omitempty"`
	From        User   `json:"from,omitempty"`
	ID          string `json:"id,omitempty"`
	Text        string `json:"text,omitempty"`
}

type Comment struct {
	CreatedTime string `json:"created_time,omitempty"`
	From        User   `json:"from,omitempty"`
	ID          string `json:"id,omitempty"`
	Text        string `json:"text,omitempty"`
}

type Counts struct {
	Media      int `json:"media,omitempty"`
	Follows    int `json:"follows,omitempty"`
	FollowedBy int `json:"followed_by,omitempty"`
}

type User struct {
	FullName       string  `json:"full_name,omitempty"`
	ID             string  `json:"id,omitempty"`
	ProfilePicture string  `json:"profile_picture,omitempty"`
	Username       string  `json:"username,omitempty"`
	Bio            string  `json:"bio,omitempty"`
	Website        string  `json:"website,omitempty"`
	Counts         *Counts `json:"counts,omitempty"`
}
