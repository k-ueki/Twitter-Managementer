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
func (d *DBHandler) BulkInsert(ids []int64) error {
	flwers := []follower{}
	for _, v := range ids {
		flwers = append(flwers, follower{PersonalID: v})
	}

	sess := SetSession()
	stmt := sess.InsertInto("followers").Columns("personal_id")

	for _, v := range flwers {
		stmt.Record(v)
	}

	_, err := stmt.Exec()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
func (d *DBHandler) Select(tablename string) (int, []follower) {
	resp := []follower{}

	sess := SetSession()
	count, err := sess.Select("personal_id,deleted").From(tablename).Where("deleted=?", 0).Load(&resp)
	if err != nil {
		fmt.Println(err)
	}

	return count, resp
}
func Comp(ids *users.FollowersIds, followers []follower) {
	var newfollowers users.FollowersIds
	fmt.Println(newfollowers)

A:
	//Idsの番号がdbにあるかチェック
	//速度改善の余地あり(今はとりあえず次へ行きます)
	for i, v := range ids.Ids {
		for i2, v2 := range followers {
			if v == v2.PersonalID {
				continue A
			}
			if i2+1 == len(followers) {
				newfollowers.Ids = append(newfollowers.Ids, v)
			}
		}
	}

	fmt.Println(newfollowers)
}
