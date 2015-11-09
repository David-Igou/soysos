package main

import (
	"log"
	"net/http"

	"github.com/emicklei/go-restful"
)

type User struct {
	ID           int
	Username     string
	Password     string
	SessionToken string
}

type UserResource struct{}

func (u UserResource) Register(container *restful.Container) {
	ws := new(restful.WebService)

	ws.
		Path("/login").
		Doc("Establishes session with the API").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	ws.Route(ws.POST("").To(u.login).
		Doc("login route").
		Operation("login").
		Reads(User{}))

	ws.Route(ws.GET("").To(u.findUser).
		Doc("find user").
		Operation("findUser"))

	ws.Filter(WebserviceLogging).Filter(MeasureTime)

	container.Add(ws)
}

//login inserts the user into the database
func (u UserResource) login(request *restful.Request, response *restful.Response) {
	user := new(User)
	err := request.ReadEntity(user)
	if err != nil {
		log.Printf(err.Error())
	}

	x := DB{Database()}
	sesh, id, err := x.InsertUser(user)
	if err != nil {
		log.Printf(err.Error())
	}
	user.SessionToken = sesh
	user.ID = id
	response.WriteHeaderAndEntity(http.StatusCreated, user)
}

//findUser will query the database for all users
func (u UserResource) findUser(request *restful.Request, response *restful.Response) {
	user := new(User)
	err := request.ReadEntity(user)
	if err != nil {
		log.Printf(err.Error())
	}
	x := DB{Database()}
	resp, err := x.FindUser(user)
	if resp {
		response.WriteHeaderAndEntity(http.StatusFound, user)
	}
	if !resp {
		response.WriteError(http.StatusBadRequest, err)
	}
}
