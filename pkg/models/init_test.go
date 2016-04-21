package models

import (
	"testing"
	"time"

	"github.com/jackc/pgx"
)

var connConfig = pgx.ConnPoolConfig{
	ConnConfig: pgx.ConnConfig{
		Host:     "jxspy.com",
		User:     "postgres",
		Password: "Jx201501",
		Database: "gmc_dev",
	},
	MaxConnections: 100,
	AcquireTimeout: time.Second * 10,
}

func TestInit(t *testing.T) {
	err := Init(&connConfig)
	if err != nil {
		t.Error(err)
	}
}
