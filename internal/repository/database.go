package repository

import (
	"database/sql"
	"time"
	"user-management-api/config"
	db "user-management-api/db/sqlc" //
	"user-management-api/internal/logger"

	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	DB      *sql.DB
	Queries *db.Queries // Fixed: use db.Queries
}

func NewDatabase(config *config.Config) (*Database, error) {
	dbConn, err := sql.Open("mysql", config.GetMySQLDSN())
	if err != nil {
		return nil, err
	}

	if err := dbConn.Ping(); err != nil {
		return nil, err
	}

	dbConn.SetMaxOpenConns(25)
	dbConn.SetMaxIdleConns(25)
	dbConn.SetConnMaxLifetime(5 * time.Minute)

	queries := db.New(dbConn) // Fixed: use db.New

	logger.Log.Info("Connected to MySQL database")

	return &Database{
		DB:      dbConn,
		Queries: queries,
	}, nil
}

func (db *Database) Close() {
	db.DB.Close()
}
