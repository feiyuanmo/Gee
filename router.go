package gee

import (
	"net/http"
	"strings"

	"github.com/feiyuanmo/gee/log"
)

type router struct {
	roots    map[string]*node
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{
		roots:    make(map[string]*node),
		handlers: make(map[string]HandlerFunc)}
}

func parsePath(path string) []string {
	vs := strings.Split(path, "/")

	parts := make([]string, 0)

	for _, item := range vs {
		if item != "" {
			parts = append(parts, item)
			if item[0] == '*' {
				break
			}
		}
	}

	return parts
}

func (r *router) addRoute(method string, path string, handler HandlerFunc) {
	parts := parsePath(path)

	//log.InfofB("Route %4s - %s", method, path)
	key := method + "-" + path
	_, ok := r.roots[method]
	if !ok {
		r.roots[method] = &node{}
	}

	r.roots[method].insert(path, parts, 0)
	r.handlers[key] = handler
}

func (r *router) getRoute(method string, path string) (*node, map[string]string) {
	searchParts := parsePath(path)

	params := make(map[string]string)

	root, ok := r.roots[method]

	if !ok {
		return nil, nil
	}

	n := root.search(searchParts, 0)

	if n != nil {
		parts := parsePath(n.path)
		for index, part := range parts {
			if part[0] == ':' {
				params[part[1:]] = searchParts[index]
			}
			if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(searchParts[index:], "/")
				break
			}
		}
		return n, params
	}

	return nil, nil
}

func (r *router) handle(c *Context) {
	if c.Path == "/favicon.ico" {

	} else {
		// key := c.Method + "-" + c.Path
		// log.InfofW("IP:%s Method:%s Path:%s", c.Host, c.Method, c.Path)
		// if hadler, ok := r.handlers[key]; ok {
		// 	hadler(c)
		// } else {
		// 	c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
		// }
		n, params := r.getRoute(c.Method, c.Path)
		if n != nil {
			c.Params = params
			key := c.Method + "-" + n.path
			log.InfofW("IP:%s Method:%s Path:%s", c.Host, c.Method, c.Path)
			r.handlers[key](c)
		} else {
			c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
		}
	}
}
