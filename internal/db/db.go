package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/LikhithMar14/social/internal/env"
)


func New(addr string,maxOpenConns,maxIdleConns int,maxIdleTime string)(*sql.DB,error){
	db,err := sql.Open("postgres",addr)
	if err != nil{
		return nil,err
	}
	db.SetMaxOpenConns(maxOpenConns)
	db.SetMaxIdleConns(maxIdleConns)
	db.SetConnMaxIdleTime(env.GetDuration("DB_MAX_IDLE_TIME",15*time.Minute))
	ctx,cancel := context.WithTimeout(context.Background(),5*time.Second)
	defer cancel()
	if err := db.PingContext(ctx);err != nil{
		return nil,err
	}
	return db,nil
}