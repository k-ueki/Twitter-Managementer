package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/k-ueki/TwitterManager/config"
	"github.com/k-ueki/TwitterManager/db"
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

	var mode string
	if r.ContentLength != 0 {
		mode = GetMode(r)
	}

	var dbh = &db.DBHandler{
		DB: config.SetDB(),
	}

	pathToGetFollowers := baseURL + "followers/list.json"
	pathToGetIds := baseURL + "followers/ids.json"
	bodyF, Ids := ucl.GetFollowersList(pathToGetFollowers, pathToGetIds)

	if mode == "register" {
		//selectfromfollowers
		count, fromdb := dbh.Select("followers")
		fmt.Println(count)

		//dbの情報とIdsを比較
		fmt.Println("FROMDB", fromdb)
		newf, byef := db.FindNewBye(&Ids, fromdb)
		fmt.Println("NEW", newf, "\nBYE", byef) //Ids

		if err := dbh.RegisterIds(newf); err != nil {
			fmt.Println("ERR", err)
		}
		fmt.Println("OK")

		//fmt.Println(err)
		//fmt.Fprintf(w,string("success for inserting!"))
		//return
	}

	fmt.Println(dbh, Ids)
	fmt.Fprintf(w, string(bodyF))
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
func GetMode(r *http.Request) string {
	var body = make([]byte, r.ContentLength)
	r.Body.Read(body)
	return Separate(string(body))
}
func Separate(str string) string {
	tmp := strings.Split(str, "=")
	return tmp[1]
}
