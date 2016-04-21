package models

import (
	"github.com/jackc/pgx"
)

var (
	Pool *pgx.ConnPool
)

func Init(connConfig *pgx.ConnPoolConfig) error {
	var err error
	Pool, err = pgx.NewConnPool(*connConfig)
	return err
}
