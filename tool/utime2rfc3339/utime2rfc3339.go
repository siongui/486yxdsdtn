package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	utime := flag.Int64("utime", 1526720351, "utime of video")
	name := flag.String("name", "instagram", "name of user")
	flag.Parse()

	ts := time.Unix(*utime, 0).Format(time.RFC3339)
	filename := fmt.Sprintf("%s-fb-live-video-%s-%d.mp4", *name, ts, *utime)
	fmt.Println(filename)
}
