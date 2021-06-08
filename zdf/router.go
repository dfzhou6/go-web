package zdf

import (
	"fmt"
	"log"
	"net/http"
)

type router struct {
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{handlers: map[string]HandlerFunc{}}
}

func (r *router) addRoute(method string, path string, handler HandlerFunc) {
	log.Printf("add Route [%s]%s", method, path)
	key := r.formatRouteKey(method, path)
	r.handlers[key] = handler
}

func (r *router) formatRouteKey(method string, path string) string {
	return fmt.Sprintf("%v-%v", method, path)
}

func (r *router) handle(c *Context) {
	key := r.formatRouteKey(c.Method, c.Path)
	if handler, ok := r.handlers[key]; ok {
		handler(c)
	} else {
		c.String(http.StatusNotFound, "404 not found: [%s]%s", c.Method, c.Path)
	}
}
