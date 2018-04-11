package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/siongui/instago"
	"github.com/siongui/instago/download"
)

var mgr *instago.IGApiManager

func handler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) != 5 {
		fmt.Fprintf(w, "invalid request")
		return
	}
	username := parts[2]
	fmt.Fprintf(w, "username: %s\n", username)
	action1 := parts[1]
	action2 := parts[3]
	fmt.Fprintf(w, "action: %s %s\n", action1, action2)

	if action1 == "download" && action2 == "all_posts" {
		igdl.DownloadAllPosts(username, mgr)
	}
}

func main() {
	mgr = instago.NewInstagramApiManager(
		os.Getenv("IG_DS_USER_ID"),
		os.Getenv("IG_SESSIONID"),
		os.Getenv("IG_CSRFTOKEN"))

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8999", nil))
}
