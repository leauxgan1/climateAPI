package server

import (
	"fmt"
	"climateAPI/types"
  "errors"
  "strconv"
	"github.com/gofiber/fiber/v2"
)


func getAll(c *fiber.Ctx) error {
  err := c.JSON(types.Store)
  return err
}


func getClimateYear(c *fiber.Ctx) error {
  id := c.Params("value")
  for i := 0; i < len(types.Store); i++ {
    currYear := types.Store[i].Year
    strID, err := strconv.Atoi(id)
    if err != nil {
      return errors.New(fmt.Sprintf("Entered ID could not be converted to a number."))
    }
    if currYear == strID {
      err := c.JSON(types.Store[i])
      return err 
    }
  }
  return errors.New(fmt.Sprintf("Location with id %s not found.",id))
  
}

func addMenuItem(c *fiber.Ctx) error {
  var err error
  return err
}

func LoadHandlers() {
	app := fiber.New(fiber.Config{
		ServerHeader:  "Fiber",
		CaseSensitive: true,
		StrictRouting: true,
		AppName:       "RestfulBoba",
	})

	app.Static("/", "./static")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hell World!")
	})

	if err := app.Get("/locations", getAll); err != nil {
	  fmt.Println("Store not found")
  }
  if err := app.Get("/locations/:value", getClimateYear); err != nil {
    fmt.Println(err)
  }
	// GET http://localhost:8080/hello%20world
	//app.Get("/:name?", func(c *fiber.Ctx) error {
	//	return c.SendString("Hello there, " + c.Params("name"))
	//})

	app.Listen(":3000")

}
