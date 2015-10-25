package main

import (
	"log"
	"net/http"

	"github.com/emicklei/go-restful"
)

func main() {
	// ws := new(restful.WebService)
	// ws.
	// 	Path("/cats").
	// 	Consumes(restful.MIME_JSON).
	// 	Produces(restful.MIME_JSON)
	cat := new(CatResource)
	wsContainer := restful.NewContainer()
	cat.Register(wsContainer)
	server := &http.Server{Addr: ":8080", Handler: wsContainer}
	//ws.Route(ws.GET("/{cat-id}").To(u.)) // u is a UserResource

	log.Fatal(server.ListenAndServe())
}

// GET http://localhost:8080/cats/1
// func (u CatResource) findCat(request *restful.Request, response *restful.Response) {
// 	id := request.PathParameter("cat-id")
// 	u.facts[id].
// }
