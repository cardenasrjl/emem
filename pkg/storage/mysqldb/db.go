package mysqldb

import (
	"context"
	"database/sql"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type DB struct {
	mysqldb *sql.DB
}

//NewDB ..
func NewDB(db *sql.DB) *DB {
	return &DB{mysqldb:db}
}

// connect returns SQL database connection from the pool
func (db *DB) Connect(ctx context.Context) (*sql.Conn, error) {
	c, err := db.mysqldb.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
	}
	return c, nil
}


