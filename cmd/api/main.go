package main

import (
	"log"

	"github.com/justmamadou/gopher-social/internal/db"
	"github.com/justmamadou/gopher-social/internal/env"
	"github.com/justmamadou/gopher-social/internal/store"
)

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":8000"),
		db: dbConfig{
			addr:        env.GetString("DB_ADDR", "postgres://admin:adminpassword@localhost/social?sslmode=disable"),
			maxOpenConn: env.GetInt("DB_MAX_OPEN_CONN", 25),
			maxIdleConn: env.GetInt("DB_MAX_IDLE_CONN", 25),
			maxIdleTime: env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
	}

	db, err := db.NewDB(cfg.db.addr, cfg.db.maxOpenConn, cfg.db.maxIdleConn, cfg.db.maxIdleTime)
	if err != nil {
		log.Panic(err)
	}

	defer db.Close()
	log.Println("Connected to database")

	store := store.NewStorage(db)

	app := &application{
		config: cfg,
		store:  store,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
