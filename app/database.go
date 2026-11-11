package app

import (
	"database/sql"
	"golang-restful-api/helper"
	"time"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3307)/golang_database_migration")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db

	// Membuat migrate
	// migrate create -ext sql -dir db/migrations create_table_first
	// migrate create -ext sql -dir db/migrations create_table_second
	// migrate create -ext sql -dir db/migrations create_table_third
	// migrate create -ext sql -dir db/migrations sample_dirty_state
	
	// Menjalankan migrate
	// migrate -database "mysql://root:root@tcp(127.0.0.1:3307)/golang_database_migration" -path db/migrations up
	// migrate -database "mysql://root:root@tcp(127.0.0.1:3307)/golang_database_migration" -path db/migrations down
	// migrate -database "mysql://root:root@tcp(127.0.0.1:3307)/golang_database_migration" -path db/migrations version
	// migrate -database "mysql://root:root@tcp(127.0.0.1:3307)/golang_database_migration" -path db/migrations force 20261111143737


}
