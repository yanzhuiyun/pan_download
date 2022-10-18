package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"pandownload/settings"
)

var (
	db  *sql.DB
	err error
)

func Init() error {
	mysqlConfig := settings.Mysql()
	fmt.Println(mysqlConfig)
	db, err = sql.Open("mysql",
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
			mysqlConfig.User,
			mysqlConfig.Password,
			mysqlConfig.Host,
			mysqlConfig.Port,
			mysqlConfig.Dbname))
	err = db.Ping()
	return err
}

func Close() {
	db.Close()
}
