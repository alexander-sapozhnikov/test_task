package dao

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"test_task/internal/logger"
)

var Entity1DB *Entity1Dao
var Entity2DB *Entity2Dao

func InitDao() {
	db := getBD()
	Entity1DB = GetEntity1Dao(db)
	Entity2DB = GetEntity2Dao(db)
}

func getBD() *sqlx.DB {
	var (
		host     = "pg_db"
		port     = 5432
		user     = "postgres"
		password = "1"
		dbname   = "entity"
	)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)


	logger.Logger.Printf("%v", psqlInfo)
	db, err := sqlx.Connect("postgres", psqlInfo)
	if err != nil {
		logger.Logger.Fatalf("%v", err)
	}
	return db
}
