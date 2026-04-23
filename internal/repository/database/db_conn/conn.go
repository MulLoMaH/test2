package db_conn

import (
	"context"
	"errors"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnDB(ctx context.Context) (*pgxpool.Pool, error) {
	dbConn := os.Getenv("CONN_BD_DOCKER")
	if dbConn == "" {
		log.Panicln("no connect string for .env")
		return nil, errors.New("no connect string for .env")
	}

	// dbConn := os.Getenv("CONN_DB")
	// if dbConn == "" {
	// 	log.Panicln("no connect string for .env")
	// 	return nil, errors.New("no connect string for .env")
	// }

	pool, err := pgxpool.New(ctx, dbConn)
	if err != nil {
		log.Println("error: ", err)
		return nil, err
	}

	if err := pool.Ping(ctx); err != nil {
		log.Println("error: ", err)
		return nil, err
	}

	return pool, nil
}
