package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/k-ueki/TwitterManager/config"
	"github.com/k-ueki/TwitterManager/router"
	"github.com/k-ueki/TwitterManager/tweets"
	"github.com/k-ueki/TwitterManager/users"
)

const (
	baseURL    = "https://api.twitter/1.1"
	clientport = ":6060"
	port       = ":7777"
)

//func NewClient() (*config.Client, *tweets.Client) {
func NewClient() *tweets.Client {
	conf, token, client := config.Set()

	return &tweets.Client{
		Config:     conf,
		Token:      token,
		HttpClient: client,
	}
}

func main() {
	var con = router.Config{
		Port: clientport,
	}
	r, cors := con.NewRouter()

	r.HandleFunc("/", APISet)
	r.HandleFunc("/followers/", users.Followers)
	//r.HandleFunc("/get/followers", getFollowers)

	err := RunServe(port, r, cors)
	if err != nil {
		log.Println(err)
	}
}

func RunServe(port string, r *mux.Router, cors func(http.Handler) http.Handler) error {
	return http.ListenAndServe(port, cors(r))
}
