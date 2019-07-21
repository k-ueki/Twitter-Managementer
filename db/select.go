package db

import "fmt"

func (d *DBHandler) Select(tablename string) (int, []follower) {
	resp := []follower{}

	sess := SetSession()
	count, err := sess.Select("personal_id,deleted").From(tablename).Where("deleted=?", 0).Load(&resp)
	if err != nil {
		fmt.Println(err)
	}

	return count, resp
}
