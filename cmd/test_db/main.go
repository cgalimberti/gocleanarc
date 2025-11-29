package main

import (
	"database/sql"
	"fmt"

	"github.com/cgalimberti/gocleanarc/20-CleanArch/internal/entity"
	"github.com/cgalimberti/gocleanarc/20-CleanArch/internal/infra/database"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE TABLE orders (id text, price float, tax float, final_price float)")
	if err != nil {
		panic(err)
	}

	repo := database.NewOrderRepository(db)

	order, _ := entity.NewOrder("1", 10.0, 1.0)
	err = repo.Save(order)
	if err != nil {
		panic(err)
	}

	order2, _ := entity.NewOrder("2", 20.0, 2.0)
	err = repo.Save(order2)
	if err != nil {
		panic(err)
	}

	orders, err := repo.List()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Found %d orders\n", len(orders))
	for _, o := range orders {
		fmt.Printf("Order: %s, Price: %f\n", o.ID, o.Price)
	}
}
