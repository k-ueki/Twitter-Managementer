package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/k-ueki/TwitterManager/config"
	"github.com/k-ueki/TwitterManager/db"
	"github.com/k-ueki/TwitterManager/timeline"
	"github.com/k-ueki/TwitterManager/users"
)

type reqbody struct {
	mode             string
	typef            string
	user_id          string
	user_screen_name string
}

// -------Followers----------
func NewUsersClient() *users.Client {
	conf, token, client := config.Set()

	return &users.Client{
		Config:     conf,
		Token:      token,
		HttpClient: client,
	}
}

func Followers(w http.ResponseWriter, r *http.Request) {
	var ucl = NewUsersClient()
	req := reqbody{}

	if r.ContentLength != 0 {
		req = SepRequest(r)
	}
	fmt.Println(req)

	var dbh = &db.DBHandler{
		DB: config.SetDB(),
	}

	pathToGetFollowers := baseURL + "followers/list.json"
	pathToGetIds := baseURL + "followers/ids.json"

	_, Ids := ucl.GetFollowersList(pathToGetFollowers, pathToGetIds)
	//bodyF, Ids := ucl.GetFollowersList(pathToGetFollowers, pathToGetIds)

	if req.mode == "init" {

		err := dbh.TruncateTable("followers")
		if err != nil {
			fmt.Println("truncate err:", err)
			fmt.Fprintf(w, "Truncate Error", err)
			return
		}

		if err := dbh.RegisterIds(Ids); err != nil {
			fmt.Println("ERR", err)
			return
		}
		fmt.Println("Complete Init your followers")
		fmt.Fprintf(w, "Complete Init your followers")
		return
	}

	if req.mode == "status" {
		_, fromdb := dbh.Select("followers")

		//dbの情報とIdsを比較
		newf, byef := db.FindNewByeIds(&Ids, fromdb)
		//fmt.Println("NEW", newf, "\nBYE", byef) //Ids

		var resp = make([]users.ResponseStruct, 2)

		if len(byef.Ids) != 0 {
			resp[1].Mode = "bye"
			users, _ := ucl.ConvertIdsToUsers(byef.Ids)

			resp[1].Users = users
		}

		if len(newf.Ids) != 0 {
			resp[0].Mode = "new"
			users, _ := ucl.ConvertIdsToUsers(newf.Ids)
			resp[0].Users = users
		}

		bytes, _ := json.Marshal(&resp)
		fmt.Fprintf(w, string(bytes))
	}

	if req.mode == "follow" {
		_, fromdb := dbh.Select("followers")

		user := users.FollowersIds{}

		num, _ := strconv.Atoi(req.user_id)
		user.Ids = append(user.Ids, int64(num))

		params := ""
		url := ""
		if req.typef == "follow" {

			if !db.IsContain(int64(num), fromdb) {
				if err := dbh.RegisterIds(user); err != nil {
					fmt.Println("follow err:", err)
				}
			}
			params = "?screen_name=" + req.user_screen_name + "&follow=true"
			url = "https://api.twitter.com/1.1/friendships/create.json" + params

		} else if req.typef == "unfollow" {

			dbh.DropOutByes(user)

			params = "?screen_name=" + req.user_screen_name
			url = "https://api.twitter.com/1.1/friendships/destroy.json" + params

		}

		_, err := ucl.HttpClient.PostForm(url, nil)
		if err != nil {
			fmt.Println("POST error : ", err)
		}
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
func SepRequest(r *http.Request) reqbody {
	res := reqbody{}
	err := r.ParseForm()
	if err != nil {
		fmt.Println("ParseError:", err)
	}
	res.mode = r.Form.Get("mode")
	res.typef = r.Form.Get("typef")
	res.user_id = r.Form.Get("user_id")
	res.user_screen_name = r.Form.Get("screen_name")

	return res
}
func Separate(str string) string {
	tmp := strings.Split(str, "=")
	return tmp[1]
}
