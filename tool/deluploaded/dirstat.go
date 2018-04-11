package main

import (
	"errors"
	"io/ioutil"
	"os"
	"time"
)

type UserStat struct {
	Name              string
	Count             int
	Size              int64
	ExpiredCount      int
	ExpiredSize       int64
	LastFileModTime   time.Time
	OldestFileModTime time.Time
}

func (u *UserStat) getDirStat(typ string) (err error) {
	if !(typ == "stories" || typ == "posts" || typ == "postlives") {
		return errors.New("bad type")
	}
	err = os.Chdir(typ)
	if err != nil {
		return
	}

	files, err := ioutil.ReadDir(".")
	if err != nil {
		return
	}
	for _, file := range files {
		u.Count++
		u.Size += file.Size()
		if file.ModTime().Before(u.OldestFileModTime) {
			u.OldestFileModTime = file.ModTime()
		}
		if file.ModTime().After(u.OldestFileModTime) {
			u.LastFileModTime = file.ModTime()
		}
		if isExpired(file.ModTime()) {
			u.ExpiredCount++
			u.ExpiredSize += file.Size()
		}
	}

	err = os.Chdir("..")
	return
}

func isExpired(t time.Time) bool {
	return time.Now().Sub(t) > 24*time.Hour
}

func getUserDirStat(username string) (user UserStat, err error) {
	err = os.Chdir(username)
	if err != nil {
		return
	}

	user.Name = username
	user.OldestFileModTime = time.Now()
	user.LastFileModTime, err = time.Parse("2006-01-02", "1900-01-01")
	if err != nil {
		return
	}
	dirs, err := ioutil.ReadDir(".")
	if err != nil {
		return
	}
	for _, dir := range dirs {
		err := user.getDirStat(dir.Name())
		if err != nil {
			return user, err
		}
	}

	err = os.Chdir("..")
	return
}

func getAllUserStat() (users []UserStat, err error) {
	err = os.Chdir(os.Getenv("IG_INSTAGRAM_DIR"))
	if err != nil {
		return
	}

	dirs, err := ioutil.ReadDir(".")
	if err != nil {
		return
	}

	for _, dir := range dirs {
		if dir.IsDir() {
			user, err := getUserDirStat(dir.Name())
			if err != nil {
				return users, err
			}

			users = append(users, user)
		}
	}

	return
}
