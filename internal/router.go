package internal

import "net/http"

type HandlerFunc func(w http.ResponseWriter, r *http.Request)

type router struct {
	groups []*routerGroup
}

func newRouter() *router {
	return &router{}
}

func (r *router) Group(name string) *routerGroup {
	g := &routerGroup{
		groupName:  name,
		handlerMap: make(map[string]HandlerFunc),
	}
	r.groups = append(r.groups, g)
	return g
}

type routerGroup struct {
	groupName  string
	handlerMap map[string]HandlerFunc
}

func (r *routerGroup) Add(name string, handlerFunc HandlerFunc) {
	r.handlerMap[name] = handlerFunc
}
