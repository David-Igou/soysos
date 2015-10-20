package main

import "github.com/emicklei/go-restful"

func main() {
	ws := new(restful.WebService)
	ws.
		Path("/cats").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	ws.Route(ws.GET("/{cat-id}").To(u.findCat)) // u is a UserResource
}

// GET http://localhost:8080/users/1
func (u CatResource) findCat(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("cat-id")
}
