package tweets

import (
	"io/ioutil"

	"github.com/k-ueki/TwitterManager/config"
)

type Client config.Client

func main() {

}
func (cli *Client) GetTimeline(path string) []byte {
	resp, _ := cli.HttpClient.Get(path)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(body))
	return body
}
