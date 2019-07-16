package db

import (
	"database/sql"
	"fmt"
)

type DBHandler struct {
	DB *sql.DB
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
	fmt.Println(flwers)
}

func main() {

}
