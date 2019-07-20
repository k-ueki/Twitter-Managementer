package main

import (
	"fmt"
	"net/http"

	"github.com/k-ueki/TwitterManager/config"
	"github.com/k-ueki/TwitterManager/timeline"
	"github.com/k-ueki/TwitterManager/users"
)

// -------Followers----------
func NewFollowersClient() *users.Client {
	conf, token, client := config.Set()

	return &users.Client{
		Config:     conf,
		Token:      token,
		HttpClient: client,
	}
}

func Followers(w http.ResponseWriter, r *http.Request) {
	var ucl = NewFollowersClient()
	method := r.Method

	var db = &users.DBHandler{
		DB: config.SetDB(),
	}

	if method == "GET" {
		pathToGetFollowers := baseURL + "followers/list.json"
		pathToGetIds := baseURL + "followers/ids.json"
		bodyF, Ids := ucl.GetFollowersList(pathToGetFollowers, pathToGetIds)

		if err := db.RegisterIds(Ids); err != nil {
			fmt.Println("ERR", err)
		}

		fmt.Fprintf(w, string(bodyF))
	}
}

// ---------------------------

// ----------Tweets-----------

// ---------------------------

// ---------timeline----------
func NewTimelineClient() *timeline.Client {
	conf, token, client := config.Set()

	return &timeline.Client{
		Config:     conf,
		Token:      token,
		HttpClient: client,
	}
}

func Timeline(w http.ResponseWriter, r *http.Request) {
	var tcl = NewTimelineClient()

	path := baseURL + "/statuses/home_timeline.json"
	body := tcl.GetTimeline(path)
	w.Write(body)
}

// --------------------------
