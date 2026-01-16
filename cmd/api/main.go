package main

import (
	_ "github.com/lib/pq"
	"github.com/ultrabor/warehouse-api/internal/app"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	app.RunApp()
}
