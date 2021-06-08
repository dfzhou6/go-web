package main

import (
	"github.com/dfzhou6/go-web/zdf"
	"log"
	"net/http"
)

func main() {
	r := zdf.New()
	r.Get("/", func(c *zdf.Context) {
		c.Html(http.StatusOK, "<h1>hello, zdf</h1>")
	})
	r.Get("/hello", func(c *zdf.Context) {
		c.String(http.StatusOK, "hello, u are %s at [%s]%s", c.Query("name"), c.Method, c.Path)
	})
	r.Post("/login", func(c *zdf.Context) {
		c.Json(http.StatusOK, zdf.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})
	log.Fatal(r.Run(":9999"))
}
