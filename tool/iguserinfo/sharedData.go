package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
)

type sharedData struct {
	EntryData struct {
		ProfilePage []struct {
			GraphQL struct {
				User IGUser `json:"user"`
			} `json:"graphql"`
		} `json:"ProfilePage"`
	} `json:"entry_data"`
}

type IGUser struct {
	Biography     string `json:"biography"`
	Id            string `json:"id"`
	Username      string `json:"username"`
	ProfilePicUrl string `json:"profile_pic_url"`
}

func getSource(username string) (b []byte, err error) {
	resp, err := http.Get("https://www.instagram.com/" + username + "/")
	if err != nil {
		return
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func getJsonStr(b []byte) string {
	pattern := regexp.MustCompile(`<script type="text\/javascript">window\._sharedData = (.*?);<\/script>`)
	m := string(pattern.Find(b))
	m1 := strings.TrimPrefix(m, `<script type="text/javascript">window._sharedData = `)
	return strings.TrimSuffix(m1, `;</script>`)
}

func decodeJsonString(s string) (user IGUser, err error) {
	d := sharedData{}
	err = json.Unmarshal([]byte(s), &d)
	user = d.EntryData.ProfilePage[0].GraphQL.User
	return
}

func main() {
	b, err := getSource(os.Getenv("IG_TEST_USERNAME"))
	if err != nil {
		panic(err)
	}

	jsonStr := getJsonStr(b)
	user, err := decodeJsonString(jsonStr)
	if err != nil {
		panic(err)
	}
	fmt.Println(user)
}
