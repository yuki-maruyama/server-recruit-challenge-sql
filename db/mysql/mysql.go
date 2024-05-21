package mysql

import (
	"database/sql"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/pulse227/server-recruit-challenge-sample/util"
)

func Init() (*sql.DB, error) {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return nil, err
	}

	c := mysql.Config{
		DBName:    util.GetEnv("DB_NAME", "app"),
		User:      util.GetEnv("DB_USER", "root"),
		Passwd:    util.GetEnv("DB_PASSWORD", "mysql"),
		Addr:      util.GetEnv("DB_ADDRESS", "localhost:3306"),
		Net:       "tcp",
		ParseTime: true,
		Collation: "utf8mb4_general_ci",
		Loc:       jst,
	}

	db, err := sql.Open("mysql", c.FormatDSN())

	if err != nil {
		return nil, err
	}

	return db, nil
}
