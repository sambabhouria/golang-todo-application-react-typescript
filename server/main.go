package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
	Body  string `json:"body"`
}

func main() {
	fmt.Print("Hello world")
	/*
				Fiber is an Express inspired web framework built on top of Fasthttp,
				the fastest HTTP engine for Go. Designed to ease things up for fast
				development with zero memory allocation and performance in mind.
				==========Quickstart==========
				package main

				import "github.com/gofiber/fiber/v2"

				func main() {
					app := fiber.New()

					app.Get("/", func(c *fiber.Ctx) error {
						return c.SendString("Hello, World 👋!")
					})
					// GET /api/register
					app.Get("/api/*", func(c *fiber.Ctx) error {
						msg := fmt.Sprintf("✋ %s", c.Params("*"))
						return c.SendString(msg) // => ✋ register
					})

					// GET /flights/LAX-SFO
					app.Get("/flights/:from-:to", func(c *fiber.Ctx) error {
						msg := fmt.Sprintf("💸 From: %s, To: %s", c.Params("from"), c.Params("to"))
						return c.SendString(msg) // => 💸 From: LAX, To: SFO
					})

					// GET /dictionary.txt
					app.Get("/:file.:ext", func(c *fiber.Ctx) error {
						msg := fmt.Sprintf("📃 %s.%s", c.Params("file"), c.Params("ext"))
						return c.SendString(msg) // => 📃 dictionary.txt
					})

					// GET /john/75
					app.Get("/:name/:age/:gender?", func(c *fiber.Ctx) error {
						msg := fmt.Sprintf("👴 %s is %s years old", c.Params("name"), c.Params("age"))
						return c.SendString(msg) // => 👴 john is 75 years old
					})

					// GET /john
					app.Get("/:name", func(c *fiber.Ctx) error {
						msg := fmt.Sprintf("Hello, %s 👋!", c.Params("name"))
						return c.SendString(msg) // => Hello john 👋!
					})

					app.Listen(":3000")
				}

				====== 📖 Serving Static Filess======
				 app.Static("/", "./public")
				// => http://localhost:3000/js/script.js
				// => http://localhost:3000/css/style.css

				app.Static("/prefix", "./public")
				// => http://localhost:3000/prefix/js/script.js
				// => http://localhost:3000/prefix/css/style.css

				app.Static("*", "./public/index.html")
				// => http://localhost:3000/any/path/shows/index/html

				===== 📖 Middleware & Next=========
				// Match any route
				app.Use(func(c *fiber.Ctx) error {
					fmt.Println("🥇 First handler")
					return c.Next()
				})

				// Match all routes starting with /api
				app.Use("/api", func(c *fiber.Ctx) error {
					fmt.Println("🥈 Second handler")
					return c.Next()
				})

				// GET /api/list
				app.Get("/api/list", func(c *fiber.Ctx) error {
					fmt.Println("🥉 Last handler")
					return c.SendString("Hello, World 👋!")
				})

				=====Using Trusted Proxy==
				app := fiber.New(fiber.Config{
		        EnableTrustedProxyCheck: true,
		        TrustedProxies: []string{"0.0.0.0", "1.1.1.1/30"}, // IP address or IP address range
		        ProxyHeader: fiber.HeaderXForwardedFor},
		    })
	*/
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		// http://localhost:5173/
		// SET HERE THE PORT YYOU CLIENT APP
		AllowOrigins: "http://localhost:5173",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	todos := []Todo{} // here we create  a slice for memorise the todos

	app.Get("/healthcheck", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	// create Todo
	app.Post("/api/todos", func(c *fiber.Ctx) error {

		// variable is declared and have the type of the sliceTodo{}
		todo := &Todo{}

		if err := c.BodyParser(todo); err != nil {
			return err
		}

		todo.ID = len(todos) + 1

		todos = append(todos, *todo)

		return c.JSON(todos)

	})

	// update todo app
	app.Patch("/api/todos/:id/done", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")

		if err != nil {
			return c.Status(401).SendString("Invalid id")
		}

		for i, t := range todos {
			if t.ID == id {
				if todos[i].Done == true {
					todos[i].Done = false
					break
				}
				todos[i].Done = true
				break
			}
		}

		return c.JSON(todos)
	})

	app.Delete("/api/todos/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")

		if err != nil {
			return c.Status(401).SendString("Invalid id")
		}
		fmt.Print(id)

		// Filter go out off here

		return c.JSON(todos)
	})

	// get all todos application
	app.Get("/api/todos", func(c *fiber.Ctx) error {
		return c.JSON(todos)
	})

	//     http://127.0.0.1:4000
	log.Fatal(app.Listen(":4000"))
}
