package server

import (
	"strconv"

	"github.com/decipher07/Go-Getting/29UrlShortener/model"
	randomforurl "github.com/decipher07/Go-Getting/29UrlShortener/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	log "github.com/sirupsen/logrus"
)

func getAllGolies(ctx *fiber.Ctx) error {
	golies, err := model.GetAllGolies()

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error getting all goly links " + err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(golies)
}

func getGoly(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error getting a param id " + err.Error(),
		})
	}

	goly, err := model.GetGoly(id)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error getting document from db " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(goly)
}

func createGoly(c *fiber.Ctx) error {
	c.Accepts("application/json")

	var goly model.Goly
	err := c.BodyParser(&goly)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error parsing JSON " + err.Error(),
		})
	}

	if goly.Random {
		goly.Goly = randomforurl.RandomURL(8)
	}

	err = model.CreateGoly(goly)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "could not create goly in db " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(goly)
}

func updateGoly(c *fiber.Ctx) error {
	c.Accepts("application/json")

	var goly model.Goly

	err := c.BodyParser(&goly)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error in body parser " + err.Error(),
		})
	}

	err = model.UpdateGoly(goly)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error getting document from db " + err.Error(),
		})
	}

	return c.Status(fiber.StatusAccepted).JSON(goly)
}

func deleteGoly(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error getting a param id " + err.Error(),
		})
	}

	err = model.DeleteGoly(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error getting document from db " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
	})
}

func redirectGoly(c *fiber.Ctx) error {
	golyUrl := c.Params("redirect")
	goly, err := model.FindByGolyUrl(golyUrl)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error getting a param id " + err.Error(),
		})
	}

	goly.Clicked += 1
	err = model.UpdateGoly(goly)

	if err != nil {
		log.Println("Sorry, An Error occured!")
	}

	return c.Redirect(goly.Redirect, fiber.StatusTemporaryRedirect)
}

func SetupAndListen() {
	router := fiber.New()

	router.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	router.Use(logger.New())

	router.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendStatus(200)
	})
	router.Get("/goly", getAllGolies)
	router.Get("/goly/:id", getGoly)
	router.Post("/goly", createGoly)
	router.Patch("/goly", updateGoly)
	router.Delete("/goly/:id", deleteGoly)
	router.Get("/r/:redirect", redirectGoly)

	log.Fatal(router.Listen(":3000"))
}
