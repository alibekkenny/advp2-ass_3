package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

func OpenDb(host string, port int, user string, password string, dbname string) (*pgxpool.Pool, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s", host, port, user, password, dbname)
	dbConn, err := pgxpool.Connect(context.Background(), dsn)
	if err != nil {
		return nil, err
	}

	return dbConn, nil
}
