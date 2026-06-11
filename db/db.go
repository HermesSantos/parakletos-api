package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/go-sql-driver/mysql"
)

func NewMySQL() (*sql.DB, error) {
	cfg := mysql.NewConfig()
	cfg.User = os.Getenv("DBUSER")
	cfg.Passwd = os.Getenv("DBPASS")
	cfg.Net = "tcp"
	cfg.Addr = "127.0.0.1:3306"
	cfg.DBName = os.Getenv("DBNAME")
	cfg.ParseTime = true

	conn, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, err
	}

	if err := conn.Ping(); err != nil {
		if err := conn.Close(); err != nil {
			return nil, err
		}

		return nil, err
	}

	fmt.Println("Connected!")

	return conn, nil
}
