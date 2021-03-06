package main

import (
	"github.com/iris-framework/iris"
	"github.com/iris-framework/iris/adaptors/cors"
	"github.com/iris-framework/iris/adaptors/httprouter"
)

func main() {

	app := iris.New()
	app.Adapt(iris.DevLogger())
	app.Adapt(httprouter.New())

	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})

	app.Adapt(crs) // this line should be added
	// adaptor supports cors allowed methods, middleware does not.

	// if you want per-route-only cors
	// then you should check https://github.com/iris-contrib/middleware/tree/master/cors

	v1 := app.Party("/api/v1")
	{
		v1.Post("/home", func(c *iris.Context) {
			app.Log(iris.DevMode, "lalala")
			c.WriteString("Hello from /home")
		})
		v1.Get("/g", func(c *iris.Context) {
			app.Log(iris.DevMode, "lalala")
			c.WriteString("Hello from /home")
		})
		v1.Post("/h", func(c *iris.Context) {
			app.Log(iris.DevMode, "lalala")
			c.WriteString("Hello from /home")
		})
	}

	app.Listen(":8080")
}
