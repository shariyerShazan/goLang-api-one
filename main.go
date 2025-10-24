package main

import (
	// "errors"
	"log"

	"github.com/gofiber/fiber/v2"
	// "github.com/valyala/fasthttp/reuseport"
)

func main() {
    app := fiber.New()

	// server main path get
    app.Get("/", func (c *fiber.Ctx) error {
        return c.SendString("Hello, World!")
    })


// version controll with api v1,v2 group
	apiV1 := app.Group("/v1")
	apiV1.Get("/" , func (c *fiber.Ctx) error {
		return c.SendString("Version 1 is running")
	})

	apiv2 := app.Group("/v2")
	apiv2.Get("/" , func (c *fiber.Ctx) error {
		return c.SendString("Version 2 is runnign")
	})

    
	// params get
    app.Get("user/:userId?" , func (c *fiber.Ctx) error {
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


	// statuc
	app.Static("/html" , "./public" ) // app.Static("api path" , "folder path" )
	// and it's return the html file of that folder path



    log.Fatal(app.Listen(":3000"))
	// app.Listen(":3000")
}