package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	// Menyajikan file statis dari folder public
	app.Static("/", "./public")

	// Route utama yang menampilkan HTML dengan link ke CSS
	app.Get("/", func(c *fiber.Ctx) error {
		html := `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<title>Hello Fiber</title>
			<link rel="stylesheet" href="/css/style.css">
		</head>
		<body>
			<h1>Hello World!</h1>
		</body>
		</html>
		`
		return c.Type("html").SendString(html)
	})

	// Menjalankan server di port 3000
	app.Listen(":3000")
}
