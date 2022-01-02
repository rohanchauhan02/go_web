package main

import (
	"database/sql"
	"src/go-echo-vue/handlers"
	"fmt"
	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
	// "github.com/labstack/echo/v4/middleware"
	// "net/http"
)

func initDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)

	//Here we check for any db errors then exit
	if err != nil {
		panic(err)
	}
	//If we don't get any errors but somehow still don't get a db connection
	//we exit as well
	if db == nil {
		panic("db nil")
	}
	return db
}
func migrate(db *sql.DB) {
	sql := `
	CREATE TABLE IF NOT EXISTS tasks(
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		name VARCHAR NOT NULL
	);   
	`
	_, err := db.Exec(sql)
	//Exit if something goes wrong with our SQL statement above
	if err != nil {
		panic(err)
	}
}
func main() {
	//connect DB
	db := initDB("storage.db")
	migrate(db)
	//Create a new instance of Echo
	fmt.Println("Searver is running on port 1323")
	// e := echo.New()
	// e.GET("/task", func(c echo.Context) error {
	// 	return c.JSON(200, "GET Tasks")
	// })
	// e.PUT("/tasks", func(c echo.Context) error {
	// 	return c.JSON(200, "PUT Tasks")
	// })
	// e.POST("/tasks", func(c echo.Context) error {
	// 	name := c.FormValue("name")
	// 	email := c.FormValue("email")
	// 	return c.JSON(200, "name: "+name+", email: "+email)
	// })

	// e.DELETE("/tasks/:id", func(c echo.Context) error {
	// 	return c.JSON(200, "DELETE Tasks "+c.Param("id"))
	// })
		e:=echo.New()

		e.File("/","public/index.html")
		e.GET("tasks",handlers.GetTasks(db))
		e.PUT("/tasks",handlers.PutTask(db))
		e.DETELE("/tasks/:id",handlers.DeleteTask(db))


	//handlers


	//Start as a web server
	e.Logger.Fatal(e.Start(":1323"))
}
