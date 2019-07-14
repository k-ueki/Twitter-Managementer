package main

import (
	"fmt"
	"net/http"
)

//to Front
// test to get timeline
func APISet(w http.ResponseWriter, r *http.Request) {
	var tcl = NewClient()

	path := baseURL + "/statuses/home_timeline.json"
	body := tcl.GetTimeline(path)
	fmt.Println("HELLO")
	w.Write(body)
}
