package main

import (
    "database/sql"
    "fmt"
    "log"
    "os"

    _ "github.com/lib/pq"

    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/template/html/v2"
)

func indexHandler(c *fiber.Ctx, db *sql.DB) error {
    var res string
    var todos []string
    rows, err := db.Query("SELECT * FROM todos")
    defer rows.Close()
    if err != nil {
        log.Fatalln(err)
        return c.Status(fiber.StatusInternalServerError).SendString("An error occurred")
    }
    for rows.Next() {
        rows.Scan(&res)
        todos = append(todos, res)
    }
    return c.Render("index", fiber.Map{
        "Todos": todos,
    })
}

type todo struct {
    Item string
}

func postHandler(c *fiber.Ctx, db *sql.DB) error {
    newTodo := todo{}
    if err := c.BodyParser(&newTodo); err != nil {
        log.Printf("An error occurred: %v", err)
        return c.Status(fiber.StatusBadRequest).SendString(err.Error())
    }
    if newTodo.Item != "" {
        _, err := db.Exec("INSERT INTO todos (item) VALUES ($1)", newTodo.Item)
        if err != nil {
            log.Printf("An error occurred while executing query: %v", err)
            return c.Status(fiber.StatusInternalServerError).SendString("An error occurred")
        }
    }

    return c.Redirect("/")
}

func updateHandler(c *fiber.Ctx, db *sql.DB) error {
    oldItem := c.Query("olditem")
    newItem := c.Query("newitem")

    if oldItem == "" || newItem == "" {
        return c.Status(fiber.StatusBadRequest).SendString("Invalid input")
    }

    _, err := db.Exec("UPDATE todos SET item=$1 WHERE item=$2", newItem, oldItem)
    if err != nil {
        log.Printf("An error occurred while executing query: %v", err)
        return c.Status(fiber.StatusInternalServerError).SendString("An error occurred")
    }
    return c.Redirect("/")
}

func deleteHandler(c *fiber.Ctx, db *sql.DB) error {
    itemToDelete := c.Query("item") // Use query parameter for the item to delete
    _, err := db.Exec("DELETE FROM todos WHERE item=$1", itemToDelete)
    if err != nil {
        log.Printf("An error occurred while executing query: %v", err)
        return c.Status(fiber.StatusInternalServerError).SendString("An error occurred")
    }
    return c.SendString("Item deleted successfully")
}

func main() {
    connStr := "postgresql://postgres:gopher@localhost/todos?sslmode=disable"
    // Connect to database
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal(err)
    }
    engine := html.New("./views", ".html")
    app := fiber.New(fiber.Config{
        Views: engine,
    })

    app.Get("/", func(c *fiber.Ctx) error {
        return indexHandler(c, db)
    })

    app.Post("/", func(c *fiber.Ctx) error {
        return postHandler(c, db)
    })

    // Route for updating a todo
    app.Put("/update", func(c *fiber.Ctx) error {
        return updateHandler(c, db)
    })

    // Route for deleting a todo
    app.Delete("/delete", func(c *fiber.Ctx) error {
        return deleteHandler(c, db)
    })

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    app.Static("/", "./public")
    log.Fatalln(app.Listen(fmt.Sprintf(":%v", port)))
}
