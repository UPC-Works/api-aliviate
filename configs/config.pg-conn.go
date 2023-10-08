package config

import (
	"context"
	"log"
	"os"
	"sync"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
)

var (
	once_pg sync.Once
	p_pg    *pgxpool.Pool
)

func Conn_Pg_DB() *pgxpool.Pool {

	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//Create the context
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()

	//Handle conection with the master
	once_pg.Do(func() {
		urlString := os.Getenv("URL_PG_DATABASE_MASTER")
		config, _ := pgxpool.ParseConfig(urlString)
		p_pg, _ = pgxpool.ConnectConfig(ctx, config)
	})

	return p_pg
}
