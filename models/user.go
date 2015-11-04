package models

import (
	"time"

	"github.com/dkostenko/instagram_crawler/instagram"
)

type User struct {
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time

	InstaID        string `sql:"unique"`
	FullName       string
	ProfilePicture string
	Username       string
	Bio            string
	Website        string
}

func NewUser(instaUser *instagram.User) *User {
	u := &User{}

	u.InstaID = instaUser.ID
	u.FullName = instaUser.FullName
	u.ProfilePicture = instaUser.ProfilePicture
	u.Username = instaUser.Username
	u.Bio = instaUser.Bio
	u.Website = instaUser.Website

	return u
}

func (u *User) Save() {
	DB.Save(u)
}
