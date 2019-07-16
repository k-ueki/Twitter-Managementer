package tweets

import "github.com/k-ueki/TwitterManager/entities"

type Tweet struct {
	CreatedAt            string            `json:created_at`
	Id                   int64             `json:id`
	IdStr                string            `json:id_str`
	Text                 string            `json:text`
	Truncated            bool              `json:truncated`
	Entities             entities.Entities `json:entities`
	Source               string            `json:source`
	InReplyToStatusId    int64             `json:in_reply_to_status_id`
	InReplyToStatusIdStr string            `json:in_reply_to_status_id_str`
	InReplyToUserId      int64             `json:in_reply_to_user_id`
	InReplyToUserIdStr   string            `json:in_reply_to_user_id_str`
	InReplyToScreenName  string            `json:in_reply_to_screen_name`
	Place                string            `json:place`
	Contributors         []int64           `json:contributors`
	RetweetedStatus      Retweet           `json:retweeted_status`
	IsQuoteStatus        bool              `json:is_quote_status`
	RetweetCount         int64             `json:retweet_count`
	FavoriteCount        int64             `json:favorite_count`
	Favorited            bool              `json:favorited`
	Retweeted            bool              `json:retweeted`
	Lang                 string            `json:lang`

	//Geo                  string            `json:geo`
	//Coordinates s
}

type Retweet struct {
	Createdat                   string            `json:created_at`
	Id                          int64             `json:id`
	IdStr                       string            `json:id_str`
	Text                        string            `json:text`
	Truncated                   bool              `json:truncated`
	Entities                    entities.Entities `json:entities`
	Place                       string            `json:place`
	Contributor                 []int64           `json:contributors`
	IsQuoteStatus               bool              `json:is_quote_status`
	RetweetCount                int64             `json:retweet_count`
	FavoriteCount               int64             `json:favorite_count`
	Favorited                   bool              `json:favorited`
	PossiblySensitive           bool              `json:possibly_sensitive`
	PossiblySensitiveAppealable bool              `json:possibly_sensitive_appealble`
	Lang                        string            `json:lang`

	//Geo string `json:geo`
}

func main() {

}
