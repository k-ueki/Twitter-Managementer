package db

import "github.com/k-ueki/TwitterManager/users"

func (d *DBHandler) DropOutByes(ids users.FollowersIds) {
	d.Drop(ids.Ids)
}
func (d *DBHandler) Drop(ids []int64) {
	sess := SetSession()
	for _, v := range ids {
		sess.DeleteFrom("followers").Where("personal_id=?", v).Exec()
	}
}
