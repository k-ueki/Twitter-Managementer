package db

import (
	"github.com/gocraft/dbr"
)

func SetSession() *dbr.Session {
	var dsn string = "root:@/TwitterManager"
	conn, _ := dbr.Open("mysql", dsn, nil)
	return conn.NewSession(nil)
}
func main() {

}
