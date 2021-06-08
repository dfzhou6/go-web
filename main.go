package main

import (
	"fmt"
	"github.com/dfzhou6/go-web/zdf"
	"log"
	"net/http"
)

func main() {
	r := zdf.New()
	r.Get("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "URL.path = %s\n", req.URL.Path)
	})
	r.Get("/hello", func(w http.ResponseWriter, req *http.Request) {
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%s] = %s\n", k, v)
		}
	})
	log.Fatal(r.Run(":9999"))
}
