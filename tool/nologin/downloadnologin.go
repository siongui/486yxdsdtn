package main

import (
	"flag"
	"fmt"

	"github.com/siongui/instago/download"
)

func main() {
	typ := flag.String("downloadtype", "post", "Download 1) post 2) profilepic 3) recent")
	arg := flag.String("argument", "", "code or username")
	flag.Parse()

	switch *typ {
	case "post":
		fmt.Println("Download single post")
		igdl.DownloadPostNoLogin(*arg)
	case "profilepic":
		fmt.Println("Download profile pic")
		igdl.DownloadUserProfilePicUrlHd(*arg)
	case "recent":
		fmt.Println("Download recent posts")
		igdl.DownloadRecentPostsNoLogin(*arg)
	default:
		fmt.Println("You have to choose a download type")
	}
}
