package main

import (
	//"embed"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"log"

	//"verify/cmd"
)


//var migrationFS embed.FS

func main() {
	engine := html.New("web/views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Static("/", "web/src")
	app.Get("/", func(c *fiber.Ctx) error {
		// Render index template
		return c.Render("index", fiber.Map{
			"Title": "Hello, World!",
		})
	})

	//cmd.Execute(migrationFS)
	log.Fatal(app.Listen(":3001"))

}
