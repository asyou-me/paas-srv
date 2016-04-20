package models

import (
	"github.com/jackc/pgx"
)

var connConfig = pgx.ConnConfig{
	Host:     "127.0.0.1",
	User:     "pgx_md5",
	Password: "secret",
	Database: "pgx_test",
}

func init() {
	conn, err := pgx.NewConnPool(*connConfig)
	if err != nil {
		fmt.Printf("Unable to establish connection: %v", err)
		return
	}
}
