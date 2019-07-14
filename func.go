package main

import (
	"fmt"
	"net/http"
)

//to Front
func APISet(w http.ResponseWriter, r *http.Request) {
	var tcl = NewClient()

	path := "https://api.twitter.com/1.1/statuses/home_timeline.json"
	body := tcl.GetTimeline(path)
	fmt.Println("HELLO")
	w.Write(body)
}
