package users

import (
	"database/sql"
	"fmt"
)

type DBHandler struct {
	DB *sql.DB
}

func (d *DBHandler) RegisterIds(ids FollowersIds) error {
	fmt.Println("KITAYO")
	fmt.Println(ids)
	fmt.Println("O", ids.Ids)
	d.BulkInsert()
	//	_, err := db.Exec("insert into Followers (PersonalID,name,screen_name,following) values (?,?,?,?)", ids.)
	//	if err != nil {
	//		return nil, err
	//	}
	//	return &User{
	//		Name:     u.Name,
	//		UserId:   u.UserId,
	//		Password: u.Password,
	//		Email:    u.Email,
	//		//File:u.File
	//	}, nil
	return nil
}
func (d *DBHandler) BulkInsert() {
	type followers struct {
		personalID int64
		deleted    int
	}

	flwers := []followers{}
	flwers = append(flwers, followers{personalID: 234})
	flwers = append(flwers, followers{personalID: 345})
	flwers = append(flwers, followers{personalID: 456})
	//follower = append(followers,)
	fmt.Println("FFFFFFFFLWERS\n", flwers)

	//stmt := sees.InsertInto("followers").Columns("PersonalID")
	//fmt.Println(stmt)
}
