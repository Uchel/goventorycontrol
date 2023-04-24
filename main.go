package main

import (
	"go_inventory_ctrl/delivery"

	_ "github.com/lib/pq"
)

func main() {
	delivery.Exec()

}
