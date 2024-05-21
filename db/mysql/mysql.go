package mysql

import (
	"database/sql"
	"time"

	"github.com/go-sql-driver/mysql"
)

func Init() (*sql.DB, error) {
	// TODO: コンフィグ突っ込めるようにする（環境変数？）
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return nil, err
	}

	c := mysql.Config{
		DBName:    "app",
		User:      "root",
		Passwd:    "mysql",
		Addr:      "db:3306",
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
