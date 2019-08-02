package users

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
)

func (u *Client) ConvertIdsToUsers(ids []int64) ([]User, []byte) {
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
	}
	url += nums

	resp, _ := u.HttpClient.Get(url)
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	if err := json.Unmarshal(body, &users); err != nil {
		fmt.Println("Err:", err)
		return nil, nil
	}

	return users, body
}
