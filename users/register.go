package users

import (
	"database/sql"
	"fmt"

	"github.com/k-ueki/TwitterManager/db"
)

type DBHandler struct {
	DB *sql.DB
}

func (d *DBHandler) RegisterIds(ids FollowersIds) error {
	if err := d.BulkInsert(ids.Ids); err != nil {
		fmt.Println("Err Bulk Insert", err)
	}
	return nil
}
func (d *DBHandler) BulkInsert(ids []int64) error {
	type followers struct {
		PersonalID int64
		Deleted    int
	}
	flwers := []followers{}
	for _, v := range ids {
		flwers = append(flwers, followers{PersonalID: v})
	}

	sess := db.SetSession()
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
