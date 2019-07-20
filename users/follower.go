package users

import (
	"encoding/json"
	"io/ioutil"

	"github.com/k-ueki/TwitterManager/config"
)

type Client config.Client

func (u *Client) GetFollowersList(path, pathToGetIds string) ([]byte, FollowersIds) {
	var followers []Followers
	var ids FollowersIds
	//	var db = &DBHandler{
	//		DB: config.SetDB(),
	//	}

	respFollowers, _ := u.HttpClient.Get(path)
	defer respFollowers.Body.Close()

	//並列化できる？
	respIds, _ := u.HttpClient.Get(pathToGetIds)
	defer respIds.Body.Close()

	bodyFollowers, _ := ioutil.ReadAll(respFollowers.Body)
	bodyIds, _ := ioutil.ReadAll(respIds.Body)
	_ = json.Unmarshal(bodyFollowers, &followers)
	_ = json.Unmarshal(bodyIds, &ids)

	//	if err := db.RegisterIds(ids); err != nil {
	//		fmt.Println(err)
	//	}

	return bodyFollowers, ids
}
