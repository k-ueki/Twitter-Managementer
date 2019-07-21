package db

import (
	"database/sql"
	"fmt"

	"github.com/k-ueki/TwitterManager/users"
)

type DBHandler struct {
	DB *sql.DB
}
type follower struct {
	PersonalID int64 `db:personal_id,json:personal_id`
	Deleted    int   `db:deleted,json:deleted`
}

func (d *DBHandler) RegisterIds(ids users.FollowersIds) error {
	if err := d.BulkInsert(ids.Ids); err != nil {
		fmt.Println("Err Bulk Insert", err)
	}
	return nil
}
func FindNewBye(ids *users.FollowersIds, followers []follower) (newfollowers, byefollowers users.FollowersIds) {

NewFollowers:
	//Idsの番号がdbにあるかチェック
	//速度改善の余地あり(今はとりあえず次へ行きます)
	for _, v := range ids.Ids {
		for i2, v2 := range followers {
			if v == v2.PersonalID {
				continue NewFollowers
			}
			if i2+1 == len(followers) {
				newfollowers.Ids = append(newfollowers.Ids, v)
			}
		}
	}

ByeFollowers:
	//DBの中の番号がIdsにあるかチェック
	//速度改善の余地あり
	for _, v := range followers {
		//fmt.Println(v.PersonalID)
		for i2, v2 := range ids.Ids {
			if v.PersonalID == v2 {
				continue ByeFollowers
			}
			if i2+1 == len(ids.Ids) {
				byefollowers.Ids = append(byefollowers.Ids, v.PersonalID)
			}
		}
	}

	return newfollowers, byefollowers
}
