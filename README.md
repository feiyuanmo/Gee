# Gee

# 1.0 雏形

利用
package http

type Handler interface {
    ServeHTTP(w ResponseWriter, r *Request)
}

func ListenAndServe(address string, h Handler) error


handler := (http.Handler)(engine) // 手动转换为借口类型
log.Fatal(http.ListenAndServe(":9999", handler))

======================================================

main

gee.New()

r.GET("/",func())

r.run("127.0.0.1:8080")

gee
type HandlerFunc func(http.ResponseWriter, *http.Request)

type Engine struct{
	router map[string]HandlerFunc
}

func New() *Engine{
	return &Engine{}
}


func (engine *Engine) addRoute{
	
}

func (engine *Engine) GET(){
	engine.addroute()
}


func (engine *Engine) Run(){
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(){
	if handler, ok := engine.router[key]; ok {
		handler()	
	}
}