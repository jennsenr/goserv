package goserv

type Group struct {
	Server *Server
	path   string
}

func NewGroup(path string, server *Server) *Group {
	return &Group{path: path, Server: server}
}

func (g *Group) GET(path string, handlerFunc HandlerFunc, middlewares ...MiddlewareFunc) {
	g.Server.GET(g.path+path, handlerFunc, middlewares...)
}

func (g *Group) POST(path string, handlerFunc HandlerFunc, middlewares ...MiddlewareFunc) {
	g.Server.POST(g.path+path, handlerFunc, middlewares...)
}

func (g *Group) PUT(path string, handlerFunc HandlerFunc, middlewares ...MiddlewareFunc) {
	g.Server.PUT(g.path+path, handlerFunc, middlewares...)
}

func (g *Group) PATCH(path string, handlerFunc HandlerFunc, middlewares ...MiddlewareFunc) {
	g.Server.PATCH(g.path+path, handlerFunc, middlewares...)
}

func (g *Group) DELETE(path string, handlerFunc HandlerFunc, middlewares ...MiddlewareFunc) {
	g.Server.DELETE(g.path+path, handlerFunc, middlewares...)
}
