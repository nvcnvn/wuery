package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/nvcnvn/wuery"
)

func main() {
	dataSource := os.Getenv("DATA_SOURCE_NAME")
	log.Println("Connection string:", dataSource)
	db, err := sql.Open("postgres", dataSource)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT * FROM wuery.all_types")
	if err != nil {
		log.Fatal(err)
	}

	colTypes, err := rows.ColumnTypes()
	for _, t := range colTypes {
		fmt.Printf("Name: %s, ScanType: %s, DatabaseTypeName: %s\n", t.Name(), t.ScanType(), t.DatabaseTypeName())
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	w := wuery.NewWuery(db)
	log.Panicln(http.ListenAndServe(":6969", wuery.NewHTTPServer(w)))
}
