package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	// Add logger middleware
	app.Use(logger.New())

	// Menyajikan file statis dari folder public
	app.Static("/", "./public")

	// Route utama yang menampilkan HTML dengan link ke CSS
	app.Get("/", func(c *fiber.Ctx) error {
		html := `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>Hallo Dunia dengan gofiber</title>
			<link rel="stylesheet" href="/css/style.css">
		</head>
		<body>
			<div class="container">
				<div class="binary-background" id="binary-bg"></div>
				<h1 class="glowing-text">Hallo Duniaaaaa!</h1>
				<h2 class="glowing-text"> halo guys  </h2>
			</div>

			<script src="/js/script.js"></script>
		</body>
		</html>
		`
		return c.Type("html").SendString(html)
	})

	// Menjalankan server di port 3000
	app.Listen(":3000")
}
