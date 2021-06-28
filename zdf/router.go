package zdf

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

type router struct {
	roots    map[string]*node
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{
		roots:    map[string]*node{},
		handlers: map[string]HandlerFunc{},
	}
}

func parsePattern(pattern string) []string {
	sl := strings.Split("/", pattern)

	var parts []string
	for _, v := range sl {
		if v == "" {
			continue
		}
		sl = append(sl, v)
		if v[0] == '*' {
			break
		}
	}

	return parts
}

func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	log.Printf("add Route [%s]%s", method, pattern)
	parts := parsePattern(pattern)
	if _, ok := r.roots[method]; !ok {
		r.roots[method] = &node{}
	}
	r.roots[method].insert(pattern, parts, 0)
	key := r.formatRouteKey(method, pattern)
	r.handlers[key] = handler
}

func (r *router) getRoute(method string, path string) (*node, map[string]string) {
	if _, ok := r.roots[method]; !ok {
		return nil, nil
	}

	searchParts := parsePattern(path)
	node := r.roots[method].search(searchParts, 0)
	if node == nil {
		return nil, nil
	}

	params := map[string]string{}
	parts := parsePattern(node.pattern)
	for i, part := range parts {
		if part[0] == ':' {
			params[part[1:]] = searchParts[i]
		} else if part[0] == '*' && len(part) > 1 {
			params[part[1:]] = strings.Join(searchParts[i:], "/")
			break
		}
	}

	return node, params
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
