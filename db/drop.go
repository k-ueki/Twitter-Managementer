package db

import (
	"fmt"

	"github.com/k-ueki/TwitterManager/users"
)

func (d *DBHandler) DropOutByes(ids users.FollowersIds) {
	d.Drop(ids.Ids)
}
func (d *DBHandler) Drop(ids []int64) {
	sess := SetSession()
	for _, v := range ids {
		sess.DeleteFrom("followers").Where("personal_id=?", v).Exec()
	}
}

func (d *DBHandler) TruncateTable(str string) error {
	fmt.Println(d)
	_, err := d.DB.Exec(fmt.Sprintf(`truncate table %s`, str))
	if err != nil {
		fmt.Println("err:", err)
		return err
	}

	return nil
}
