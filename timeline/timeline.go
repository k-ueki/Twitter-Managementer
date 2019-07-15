package timeline

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/k-ueki/TwitterManager/config"
	"github.com/k-ueki/TwitterManager/tweets"
)

type Client config.Client

func (cli *Client) GetTimeline(path string) []byte {
	var tweets []tweets.Tweet
	resp, _ := cli.HttpClient.Get(path)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	_ = json.Unmarshal(body, &tweets)
	fmt.Println("TWEETS", tweets)
	//fmt.Println(string(body))
	return body
}
