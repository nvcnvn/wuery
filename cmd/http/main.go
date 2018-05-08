package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/nvcnvn/wuery"
)

func main() {
	db, err := sql.Open("postgres", "postgresql://root@solus:26257/wuery?sslmode=disable")
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

	http.ListenAndServe(":8080", &wuery.HTTPServer{})
}
