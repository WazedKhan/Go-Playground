package main

import (
	"database/sql"
	"flag"
	"log"

	"cli-tool/internal"
	"cli-tool/repository"
)

func main() {
	useDB := flag.Bool("db", false, "use sqlite backend")
	flag.Parse()

	if *useDB {
		db, err := sql.Open("sqlite3", repository.TodoSqlitePath())
		if err != nil {
			log.Fatal("failed to open sqlite", err)
		}
		defer db.Close()

		store, err := repository.InitSqliteStore(db)
		if err != nil {
			log.Fatal("failed to initialize sqlite store", err)
		}
		repository.SetTodoStore(store)
	}
	internal.AppLoop()
}
