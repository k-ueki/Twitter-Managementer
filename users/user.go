package users

import (
	"github.com/k-ueki/TwitterManager/entities"
)

const (
	baseURL = "https://api.twitter.com/1.1/"
)

type (
	User struct {
		Id                             int64             `json:"id"`
		IdStr                          string            `json:"id_str"`
		Name                           string            `json:"name"`
		ScreenName                     string            `json:"screen_name"`
		Location                       string            `json:"location"`
		Description                    string            `json:"description"`
		Url                            string            `json:"url"`
		Entities                       entities.Entities `json:"entities"`
		Protected                      bool              `json:"protected"`
		FollowersCount                 int64             `json:"followers_count"`
		FriendsCount                   int64             `json:"friends_count"`
		ListedCount                    int64             `json:"listed_count"`
		Createdat                      string            `json:"created_at"`
		FavouritesCount                int64             `json:"fabourites_count"`
		UtcOffset                      int64             `json:"utc_offset"`
		TimeZone                       string            `json:"time_zone"`
		GeoEnabled                     bool              `json:"geo_enabled"`
		Verified                       bool              `json:"verified"`
		StatusesCount                  int64             `json:"statuses_count"`
		Lang                           string            `json:"lang"`
		ContributorsEnabled            bool              `json:"contiributors_enabled"`
		IsTranslator                   bool              `json:"is_translator"`
		IsTranslationEnabled           bool              `json:"is_translation_enabled"`
		ProfileBackgroundColor         string            `json:"profile_background_color"`
		ProfileBackgroundImageUrl      string            `json:"profile_background_image_url"`
		ProfileBackgroundImageUrlHttps string            `json:"profile_background_image_url_https"`
		ProfileBackgroundTile          bool              `json:"profile_background_tile"`
		ProfileImageUrl                string            `json:"profile_image_url"`
		ProfileImageUrlHttps           string            `json:"profile_image_url_https"`
		ProfileLinkColor               string            `json:"profile_link_color"`
		ProfileSidebarBorderColor      string            `json:"profile_sidebar_border_color"`
		ProfileSidebarFillColor        string            `json:"profile_sidebar_fill_color"`
		ProfileTextColor               string            `json:"profile_text_color"`
		ProfileUseBackgroundImage      bool              `json:"profile_use_background_image"`
		HasExtendedProfile             bool              `json:"has_extended_profile"`
		DefaultProfile                 bool              `json:"default_profile"`
		DefaultProfileImage            bool              `json:"default_profile_image"`
		Following                      bool              `json:"following"`
		FollowRequestSent              bool              `json:"follow_request_sent"`
		Notifications                  bool              `json:"notifications"`
		TranslatorType                 string            `json:"translator_typa"`
	}

	Followers struct {
		Users             []User `json:"users"`
		NextCursor        int64  `json:"next_cursor"`
		NextCursorStr     string `json:"next_cursor_str"`
		PreviousCursor    int64  `json:"previous_cursor"`
		PreviousCursorStr string `json:"previous_cursor_str"`
	}

	FollowersIds struct {
		Ids               []int64 `json:"ids"`
		NextCursor        int64   `json:"next_cursor"`
		NextCursorStr     string  `json:"next_cursor_str"`
		PreviousCursor    int64   `json:"previous_cursor"`
		PreviousCursorStr string  `json:"previous_cursor_str"`
	}

	UserInfo struct {
		ScreenName string
		UserId     int64
		Follow     bool
		Client     *Client
	}

	ResponseStruct struct {
		Mode  string `json:"mode"`
		Users []User `json:"users"`
	}
)

func main() {}
