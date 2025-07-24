package main

import (
	"github.com/jackc/pgx/v5"
	"time"
)

type Compound struct {
	ID        int64
	CreatedAt time.Time
	CompNo    string
	Amount    float64
	KodHasil  string
}

func getCompoundByCompoundNo(conn pgx.Conn) {

}
