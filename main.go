package main

import (
	"github.com/gofiber/fiber"
)

func setUpRoutes(app *fiber.App) {
	app.Get(GetLeads)
	app.GetO(GetLead)
	app.Post(NewLead)
	app.Delete(DeleteLead)
}
func main() {
	app := fiber.New()
	setUpRoutes(app)
	app.Listen(3000)
}
