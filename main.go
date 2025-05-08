package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	// Static folder for CSS
	app.Static("/", "./public")

	// Route dengan HTML dan CSS
	app.Get("/", func(c *fiber.Ctx) error {
		html := `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<title>Hello Fiber</title>
			<link rel="stylesheet" href="/css/styles.css">
		</head>
		<body>
			<h1>Halo ini Hello, World! dengan GoFiber</h1>
		</body>
		</html>
		`
		return c.Type("html").SendString(html)
	})

	app.Listen(":3000")
}
