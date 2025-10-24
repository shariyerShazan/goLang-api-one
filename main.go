package main

import (
    "log"

    "github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()

	// server main path get
    app.Get("/", func (c *fiber.Ctx) error {
        return c.SendString("Hello, World!")
    })
    
	// params get
    app.Get("/:userId?" , func (c *fiber.Ctx) error {
          if c.Params("userId") != "" {
            return c.SendString("User is is = " + c.Params("userId"))
          } else {
			return c.Status(fiber.StatusBadRequest).SendString("User ID is missing")
		  }
	  })     

	//  api path get
	app.Get("/api/*", func (c *fiber.Ctx) error {
		return c.SendString("API path: " + c.Params("*"))
	} )

	
    log.Fatal(app.Listen(":3000"))
}