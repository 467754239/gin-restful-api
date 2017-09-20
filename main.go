package main

import (
	db "github.com/467754239/db-api/database"
)

func main() {
	defer db.SqlDB.Close()
	router := initRouter()
	router.Run(":8000")
}
