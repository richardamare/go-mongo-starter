package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"go-mongo-starter/config"
	"go-mongo-starter/database"
	"go-mongo-starter/router"
	"os"
)

func SetupAndRunApp() error {
	// load env
	if err := config.LoadENV(); err != nil {
		return err
	}

	// connect to db
	if err := database.StartMongoDB(); err != nil {
		return err
	}

	// defer close db
	defer database.CloseMongoDB()

	// load stripe
	//if err := config.LoadStripe(); err != nil {
	//	return err
	//}

	// create fiber app
	app := fiber.New()

	// attach middleware
	app.Use(recover.New())
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path} ${latency}\n",
	}))
	app.Use(cors.New())

	// attach routes
	router.SetupRoutes(app)

	// attach swagger
	config.AddSwaggerRoutes(app)

	// attach monitoring
	app.Get("/metrics", monitor.New(monitor.Config{Title: "Digital Farm API Metrics"}))

	// start app
	port := os.Getenv("PORT")
	return app.Listen(":" + port)
}
