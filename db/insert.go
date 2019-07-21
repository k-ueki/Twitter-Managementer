package db

import "fmt"

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
