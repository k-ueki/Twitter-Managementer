package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/k-ueki/TwitterManager/config"
	"github.com/k-ueki/TwitterManager/router"
)

const (
	baseURL    = "https://api.twitter.com/1.1/"
	clientport = ":6060"
	port       = ":7777"
)

func main() {
	api := config.GetTwitterApi()
	fmt.Println(api)

	var con = router.Config{
		Port: clientport,
	}
	r, cors := con.NewRouter()

	r.HandleFunc("/", APISet)
	r.HandleFunc("/followers/", Followers)

	err := RunServe(port, r, cors)
	if err != nil {
		log.Println(err)
	}
}

func RunServe(port string, r *mux.Router, cors func(http.Handler) http.Handler) error {
	return http.ListenAndServe(port, cors(r))
}
