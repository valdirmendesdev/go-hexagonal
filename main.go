package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/valdirmendesdev/go-hexagonal/adapters/db"
	"github.com/valdirmendesdev/go-hexagonal/application"
)

func main() {
	dbConn, _ := sql.Open("sqlite3", "db.sqlite")
	productDbAdapter := db.NewProductDb(dbConn)
	service := application.NewProductService(productDbAdapter)
	product, _ := service.Create("Meu produto", 25.90)

	service.Enable(product)

	fmt.Println("Funcionou!")

}
