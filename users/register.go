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
