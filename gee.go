package gee

import (
	"fmt"
	"net/http"
)

type Engine struct{}

func New() *Engine {
	return &Engine{}
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
}

func (engine *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, engine)
}
