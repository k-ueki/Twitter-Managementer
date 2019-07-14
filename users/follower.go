package users

import (
	"fmt"
	"io/ioutil"
	"net/url"

	"github.com/ChimeraCoder/anaconda"
	"github.com/k-ueki/TwitterManager/config"
)

//type Client struct {
//	Config     *oauth1.Config
//	Token      *oauth1.Token
//	HttpClient *http.Client
//}
type Client config.Client

func (u *Client) GetFollowersList(path string) []byte {
	var twitter anaconda.TwitterApi
	v := url.Values{}
	res := twitter.GetFollowersListAll(v)
	fmt.Println(res)

	resp, _ := u.HttpClient.Get(path)
	if resp.StatusCode != 200 {
		return nil
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	//err := json.Unmarshal(body)
	//fmt.Println(err, "ERR")

	return body
}
