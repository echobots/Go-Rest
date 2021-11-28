package main

import (
	"go-rest/db"
	"go-rest/routes"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db.Init()

	e := routes.Init()

	e.Logger.Fatal(e.Start(":1323"))
}
