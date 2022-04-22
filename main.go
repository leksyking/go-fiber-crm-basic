package main

import (
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	"github.com/leksyking/go-fiber-crm-basic/database"
	"github.com/leksyking/go-fiber-crm-basic/lead"
)

func setUpRoutes(app *fiber.App) {
	app.Get(GetLeads)
	app.GetO(GetLead)
	app.Post(NewLead)
	app.Delete(DeleteLead)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "leads.db")
	if err != nil {
		panic("failed to connect to database")
	}
	fmt.Println("Database connection opened")
	database.DBConn.AutoMigrate(&lead.Lead{})
	fmt.Println("Database migrated")
}

func main() {
	app := fiber.New()
	initDatabase()
	setUpRoutes(app)
	app.Listen(3000)
	defer database.DBConn.Close()
}
