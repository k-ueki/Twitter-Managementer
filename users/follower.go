package users

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/k-ueki/TwitterManager/config"
)

//type Client struct {
//	Config     *oauth1.Config
//	Token      *oauth1.Token
//	HttpClient *http.Client
//}
type Client config.Client

func (u *Client) GetFollowersList(path string) []byte {
	var followers []Followers

	resp, _ := u.HttpClient.Get(path)
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	_ = json.Unmarshal(body, &followers)
	fmt.Println("Followers", followers, "\n")

	return body
}
