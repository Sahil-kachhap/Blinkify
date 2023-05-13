package main

import(
	"log"
	"os"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func setupRoutes(app *fiber.App){
	app.Get("/:url", routes.ResolveUrl)
	app.Post("/api/v1", routes.ShortenUrl)
}

func main(){
  err := godotenv.Load()
  app := fiber.New()
  app.Use(logger.New())
  setupRoutes(app)
  log.Fatal(app.listen(os.Getenv("APP_PORT")))
}