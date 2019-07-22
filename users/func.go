package users

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
)

//func ConvertIdsToUsers(ids []int64) users.User {
func (u *Client) ConvertIdsToUsers(ids []int64) []byte {
	var users []User
	var nums string

	url := baseURL + "users/lookup.json?user_id="

	for i, v := range ids {
		if i != 0 {
			nums += "," + strconv.FormatInt(v, 10)
		}
		if i == 0 {
			nums += strconv.FormatInt(v, 10)
		}
		fmt.Println(i)
	}
	url += nums
	fmt.Println(url)

	resp, _ := u.HttpClient.Get(url)
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	_ = json.Unmarshal(body, &users)

	fmt.Println(users)

	return body
}
