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

	// Carga las variables de entorno desde el archivo .env.local
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatalf("Error al cargar el archivo .env.local: %v", err)
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
