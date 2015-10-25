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

	cmap := make(map[string]CatFact)
	cmap["Lion"] = CatFact{"1", "Lion", "Lions have sharp teef :D"}
	cat := CatResource{cmap}
	wsContainer := restful.NewContainer()
	cat.Register(wsContainer)
	server := &http.Server{Addr: "localhost:8080", Handler: wsContainer}
	log.Printf("start listening on localhost:8080")

	//ws.Route(ws.GET("/{cat-id}").To(u.)) // u is a UserResource

	log.Fatal(server.ListenAndServe())
}

// GET http://localhost:8080/cats/1
// func (u CatResource) findCat(request *restful.Request, response *restful.Response) {
// 	id := request.PathParameter("cat-id")
// 	u.facts[id].
// }
