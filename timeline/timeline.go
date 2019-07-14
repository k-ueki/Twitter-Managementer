package timeline

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/ChimeraCoder/anaconda"
	"github.com/k-ueki/TwitterManager/config"
)

type Client config.Client

type Api struct {
	TwitterApi *anaconda.TwitterApi
}

func main() {

}

func (a *Api) Timeline(w http.ResponseWriter, r *http.Request) {
}

func (cli *Client) GetTimeline(path string) []byte {
	resp, _ := cli.HttpClient.Get(path)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(body))
	return body
}
