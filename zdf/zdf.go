package zdf

import (
	"fmt"
	"log"
	"net/http"
)

type HandlerFunc func(w http.ResponseWriter, req *http.Request)

type Engine struct {
	router map[string]HandlerFunc
}

func New() *Engine {
	return &Engine{router: map[string]HandlerFunc{}}
}

func (engine *Engine) addRoute(method string, path string, handler HandlerFunc) {
	key := fmt.Sprintf("%s-%s", method, path)
	log.Printf("addRoute: %s", key)
	engine.router[key] = handler
}

func (engine *Engine) Get(path string, handler HandlerFunc) {
	engine.addRoute("GET", path, handler)
}

func (engine *Engine) Post(path string, handler HandlerFunc) {
	engine.addRoute("POST", path, handler)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	method := req.Method
	key := fmt.Sprintf("%s-%s", method, path)
	if handler, ok := engine.router[key]; ok {
		handler(w, req)
	} else {
		fmt.Fprintf(w, "404 not found: %s\n", key)
	}
}

func (engine *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, engine)
}
