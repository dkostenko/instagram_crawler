package main

import (
	"flag"
	"fmt"

	"github.com/dkostenko/instagram_crawler/downloader"
	"github.com/dkostenko/instagram_crawler/instagram"
)

var (
	accessToken       string
	interestingUserID string
)

func init() {
	flag.StringVar(&accessToken, "at", "", "Access token for Instagram API.")
	flag.StringVar(&interestingUserID, "id", "", "An Instagram user ID whose information should be downloaded.")
}

func main() {
	flag.Parse()

	if accessToken == "" {
		fmt.Println("You should set \"at\" flag")
		return
	}

	if interestingUserID == "" {
		fmt.Println("You should set \"id\" flag")
		return
	}

	instaClient := instagram.NewClient(accessToken)
	downloader := &downloader.Downloader{instaClient}
	downloader.GetWhoFollowsUser(interestingUserID)
}
