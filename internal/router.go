package internal

import "net/http"

type Context struct {
	W http.ResponseWriter
	R *http.Request
}

type HandlerFunc func(ctx *Context)

type router struct {
	groups []*routerGroup
}

func newRouter() *router {
	return &router{}
}

func (r *router) Group(name string) *routerGroup {
	g := &routerGroup{
		groupName:        name,
		handlerMap:       make(map[string]HandlerFunc),
		handlerMethodMap: make(map[string][]string),
	}
	r.groups = append(r.groups, g)
	return g
}

type routerGroup struct {
	groupName        string
	handlerMap       map[string]HandlerFunc
	handlerMethodMap map[string][]string
}

func (r *routerGroup) add(name string, handlerFunc HandlerFunc) {
	r.handlerMap[name] = handlerFunc
}

func (r *routerGroup) GET(name string, handlerFunc HandlerFunc) {
	r.add(name, handlerFunc)
	r.handlerMethodMap[GET] = append(r.handlerMethodMap[GET], name)
}

func (r *routerGroup) POST(name string, handlerFunc HandlerFunc) {
	r.add(name, handlerFunc)
	r.handlerMethodMap[POST] = append(r.handlerMethodMap[POST], name)
}

func (r *routerGroup) PUT(name string, handlerFunc HandlerFunc) {
	r.add(name, handlerFunc)
	r.handlerMethodMap[PUT] = append(r.handlerMethodMap[PUT], name)
}

func (r *routerGroup) DELETE(name string, handlerFunc HandlerFunc) {
	r.add(name, handlerFunc)
	r.handlerMethodMap[DELETE] = append(r.handlerMethodMap[DELETE], name)
}

func (r *routerGroup) ANY(name string, handlerFunc HandlerFunc) {
	r.add(name, handlerFunc)
	r.handlerMethodMap[ANY] = append(r.handlerMethodMap[ANY], name)
}
