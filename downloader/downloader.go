package downloader

import (
	"github.com/dkostenko/instagram_crawler/instagram"
	"github.com/dkostenko/instagram_crawler/models"
)

type Downloader struct {
	InstaClient *instagram.Client
}

func (d *Downloader) GetWhoFollowsUser(userID string) {
	resp, _ := d.InstaClient.GetWhoFollowsUser(userID)

	for _, instaUser := range resp.Data {
		u := models.NewUser(&instaUser)
		u.Save()
	}
}
