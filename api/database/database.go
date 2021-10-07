package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	cc "github.com/onurozerdal/hepsiburada-case/api/config"
)

var config cc.Config

func NewRecommendedDbConnection() *sql.DB {
	config.Read()
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.RecommendedDb.Server, config.RecommendedDb.Port, config.RecommendedDb.User, config.RecommendedDb.Password, config.RecommendedDb.Database)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected recommended db!")
	return db
}
