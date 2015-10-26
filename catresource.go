package main

import (
	"net/http"
	"strconv"

	"github.com/emicklei/go-restful"
)

type CatFact struct {
	ID      string `json:"ID"`
	Species string `json:"species"`
	Fact    string `json:"fact"`
}

type CatFacts []CatFact

type CatResource struct {
	facts map[string]CatFact
}

type Message struct {
	Text string
}

func (u CatResource) Register(container *restful.Container) {
	ws := new(restful.WebService)

	ws.
		Path("/cats").
		Doc("Get some facts about cats :D").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	ws.Route(ws.GET("/").To(u.home).
		Doc("homepage of API").
		Operation("home"))

	ws.Route(ws.GET("/{fact-id}").To(u.serveFact).
		Doc("get a cat fact!").
		Operation("serveFact").
		Param(ws.PathParameter("fact-id", "identifier of the catfact").DataType("string")).
		Writes(CatFact{})) // on the response

	ws.Route(ws.POST("").To(u.createFact).
		// docs
		Doc("create a cat fact!").
		Operation("createFact").
		Reads(CatFact{})) // from the request

	container.Add(ws)
}

// GET http://localhost:8080/cats/1
//
func (u CatResource) home(reqest *restful.Request, response *restful.Response) {
	p := &Message{"Welcome to catfacts API!"}
	response.WriteEntity(p)
}

func (u CatResource) serveFact(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("fact-id")
	cat := u.facts[id]
	if len(cat.ID) == 0 {
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusNotFound, "404: Fact not found!.")
		return
	}
	response.WriteEntity(cat)
}

func (u CatResource) createFact(request *restful.Request, response *restful.Response) {
	cat := new(CatFact)
	err := request.ReadEntity(cat)
	if err != nil {
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}
	cat.ID = strconv.Itoa(len(u.facts) + 1) // simple id generation
	u.facts[cat.ID] = *cat
	response.WriteHeaderAndEntity(http.StatusCreated, cat)
}
