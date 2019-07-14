package main

import (
	"fmt"
	"net/http"
)

//to Front
func APISet(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HELLO")
}
