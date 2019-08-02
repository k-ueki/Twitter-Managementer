package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/k-ueki/TwitterManager/config"
	"github.com/k-ueki/TwitterManager/db"
	"github.com/k-ueki/TwitterManager/timeline"
	"github.com/k-ueki/TwitterManager/users"
)

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

	var mode string
	if r.ContentLength != 0 {
		mode = GetMode(r)
	}

	var dbh = &db.DBHandler{
		DB: config.SetDB(),
	}

	pathToGetFollowers := baseURL + "followers/list.json"
	pathToGetIds := baseURL + "followers/ids.json"

	_, Ids := ucl.GetFollowersList(pathToGetFollowers, pathToGetIds)
	//bodyF, Ids := ucl.GetFollowersList(pathToGetFollowers, pathToGetIds)

	if mode == "init" {

		err := dbh.TruncateTable("followers")
		if err != nil {
			fmt.Println("truncate err:", err)
			fmt.Fprintf(w, "Truncate Error", err)
			return
		}

		//init register
		if err := dbh.RegisterIds(Ids); err != nil {
			fmt.Println("ERR", err)
			return
		}
		fmt.Println("Complete Init your followers")
		fmt.Fprintf(w, "Complete Init your followers")
		return
	}
	if mode == "register" {
		_, fromdb := dbh.Select("followers")

		//dbの情報とIdsを比較
		newf, byef := db.FindNewByeIds(&Ids, fromdb)
		//fmt.Println("NEW", newf, "\nBYE", byef) //Ids

		var resp = make([]users.ResponseStruct, 2)
		if len(byef.Ids) != 0 {
			resp[1].Mode = "bye"
			users, _ := ucl.ConvertIdsToUsers(byef.Ids)

			resp[1].Users = users
			fmt.Println("RESP\n", resp)
		}

		if len(newf.Ids) != 0 {
			resp[0].Mode = "new"
			users, _ := ucl.ConvertIdsToUsers(newf.Ids)
			resp[0].Users = users
			fmt.Println("RESP", resp)
		}

		bytes, _ := json.Marshal(&resp)
		fmt.Fprintf(w, string(bytes))
		return

		//init register
		//if err := dbh.RegisterIds(Ids); err != nil {
		//	fmt.Println("ERR", err)
		//}
		//fmt.Println("OK")

		//-----------new register動作確認済み
		//		if err := dbh.RegisterIds(newf); err != nil {
		//			fmt.Println("ERR", err)
		//		}

		//------------bye dropout動作確認済み
		//if len(byef.Ids) >= 1 {
		//dbh.DropOutByes(byef)
		//}

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
func GetMode(r *http.Request) string {
	var body = make([]byte, r.ContentLength)
	r.Body.Read(body)
	return Separate(string(body))
}
func Separate(str string) string {
	tmp := strings.Split(str, "=")
	return tmp[1]
}
