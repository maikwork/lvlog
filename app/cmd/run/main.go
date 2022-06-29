package main

import (
	"database/sql"

	"github.com/maikwork/lvlog/internal/model"
	"github.com/maikwork/lvlog/internal/server"
	"github.com/maikwork/lvlog/pkg/logger"

	_ "github.com/lib/pq"
)

func main() {
	log := logger.NewLog()

	cnfDB := model.Database{
		Type:     "postgres",
		Username: "rbliss",
		Password: "pass",
		Host:     "postgresdb",
		Port:     "5432",
		DBName:   "test",
	}

	db, err := sql.Open(cnfDB.Type, cnfDB.GetDSN())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	s := server.Server{
		DB:  db,
		Log: log,
	}

	s.Run()
}
