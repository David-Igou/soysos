package main

import "github.com/emicklei/go-restful"

type RootResource struct {
	Message string
}

func (u RootResource) Register(container *restful.Container) {
	ws := new(restful.WebService)
	ws.
		Path("/").
		Doc("Home of catfact API")

	ws.Route(ws.GET("/").To(u.root).
		Doc("root").
		Operation("root"))

	container.Add(ws)
}

func (u RootResource) root(request *restful.Request, response *restful.Response) {
	p := &Message{"Welcome to catfacts API!"}
	response.WriteEntity(p)
}