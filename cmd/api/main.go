package main

import (
	"log"

	"github.com/LikhithMar14/social/internal/db"
	"github.com/LikhithMar14/social/internal/env"
	"github.com/LikhithMar14/social/internal/store"
)

const version = "0.0.1"
func main() {

	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
		db: dbConfig{
			addr:         env.GetString("DB_ADDR", "postgres://admin:adminpassword@localhost:5432/social?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
		env: env.GetString("ENV", "development"),
	}
	db, err := db.New(cfg.db.addr, cfg.db.maxOpenConns, cfg.db.maxIdleConns, cfg.db.maxIdleTime)
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()
	store := store.NewStorage(db)
	log.Println("database connection pool established")

	app := &application{
		config: cfg,
		store:  store,
	}
	httpHandler := app.mount()
	log.Fatal(app.run(httpHandler))

}
