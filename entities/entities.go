package entities

type Entities struct {
	Hashtags []struct {
		Text    string `json:text`
		indices []int  `json:indices`
	} `json:hashtags`
	UserMentions []struct {
		Id         int64  `json:id`
		IdStr      string `json:id_str`
		Name       string `json:name`
		ScreenName string `json:screen_name`
		indices    []int  `json:indices`
	} `json:user_mentions`
	Urls         []Urls      `json:urls`
	Url          UrlEntities `json:url`
	Deiscriprion []Urls      `json:description`
}

type Urls struct {
	Url         string `json:url`
	ExpandedUrl string `json:expanded_url`
	DisplayUrl  string `json:display_url`
	Indices     []int  `json:indices`
}

type UrlEntities struct {
	Urls []struct {
		Url         string `json:url`
		ExpandedUrl string `json:expanded_url`
		DisplayUrl  string `json:display_url`
		Indices     []int  `json:indices`
	} `json:urls`
}

func main() {

}
