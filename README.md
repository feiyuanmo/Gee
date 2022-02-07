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

# 2.0 上下文封装

type router struct {
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{handlers: make(map[string]HandlerFunc)}
}

type Engine struct {
	router *router
}

func New() *Engine {
	return &Engine{router: newRouter()}
}
======================================================



func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	engine.router.handle(c)
}


type H map[string]interface{}

type Context struct {
	// origin objects
	Writer http.ResponseWriter
	Req    *http.Request
	// request info
	Path   string
	Method string
	// response info
	StatusCode int
}

func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
	}
}

func (c *Context) PostForm(key string) string {
}

========================================================
func (r *router) handle(c *Context) {
	handler(c)
}

#3.0路由树
========================================================
type router struct {
	roots    map[string]*node
	handlers map[string]HandlerFunc
}

type node struct {
	pattern  string 
	part     string 
	children []*node 
	isWild   bool 
}
handlers map[路径] + 回调函数
roots  存储 + 解析 合格的路径   用于handlers使用


func (n *node) insert(pattern string, parts []string, height int) {
	//用于插入 递归函数
	if len(parts) == height {
		n.pattern = pattern
		return
	} //中断条件

	child := n.matchChild(part) // 第一个匹配成功的节点

	....//为空时插入 

	child.insert(pattern, parts, height+1) //递归 进行下一层插入
}

func (n *node) search(parts []string, height int) *node {
	//用于查找
	if len(parts) == height || strings.HasPrefix(n.part, "*") {

	}//中断条件

	part := parts[height]
	children := n.matchChildren(part)  // 所有匹配成功的节点

	for _, child := range children { //对这一层匹配的节点 进行下一层匹配 
		result := child.search(parts, height+1) 
		if result != nil {
			return result
		}
	}

	return nil //都没有就返回
}

func parsePattern(pattern string) []string {
	//解析路径 分成 []string
}

func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	parsePattern(pattern)
	...
	r.roots[method].insert(pattern, parts, 0)
	...
	key := method + "-" + pattern
	r.handlers[key] = handler
}

func (r *router) getRoute(method string, path string) (*node, map[string]string) {
	parsePattern(path)
	...
	root.search(searchParts, 0)
	...
}

func (r *router) handle(c *Context) {
	n, params := r.getRoute(c.Method, c.Path)
	key := c.Method + "-" + n.pattern
	r.handlers[key](c)
}
=================================================================================
 测试
func TestNamexxxx(t *testing.T) {

}

reflect.DeepEqual() 
func DeepEqual(x, y interface{}) bool
用于检查是否x和y是“deeply equal”与否
=================================================================================
addRouter -> getRouter -> handle 
用前缀树结构存 前缀树结构取, 用GET /a/asd/c || GET a/s/c 匹配到路由(GET-/a/:param/c)对应的HandlerFunc 并把asd || s 存在context的Params里.
=================================================================================
#4.0分组路由
=================================================================================
go 中的嵌套类型 (类似 Java/Python 等语言的继承)

RouterGroup struct {
	.....
	engine      *Engine      
}

Engine struct {
	*RouterGroup
	...
}

func New() *Engine {
	engine := &Engine{router: newRouter()}
    engine.RouterGroup = &RouterGroup{engine: engine}
	return engine
}

 (*Engine).engine 指向自己  套娃？
=================================================================================
pattern := group.prefix + comp

group.engine.router.addRoute(method, pattern, handler)

engine.router.addRoute(method, pattern, handler)

=================================================================================

#5.0中间件
思路 Context handlers []HandlerFunc 按顺序放好要执行的handlers 然后一个一个的执行

func (c *Context) Next() {
	c.index++
	s := len(c.handlers)
	for ; c.index < s; c.index++ {
		c.handlers[c.index](c)
	}
}

func A(c *Context) {
    part1
    c.Next()
    part2
}
func B(c *Context) {
    part3
    c.Next()
    part4
}

part1 -> part3 -> Handler -> part 4 -> part2