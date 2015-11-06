package main

import (
	"log"

	"github.com/emicklei/go-restful"
)

type User struct {
	//ID       bson.ObjectId `bson:"_id,omitempty"`
	Username string
	Password []byte
	//Posts    int
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

//Login validates and returns a user object if they exist in the database.
func LoginUser(user *User) bool {
	x := DB{Database()}
	x.InsertUser(user)
	return true
}

func (u UserResource) login(request *restful.Request, response *restful.Response) {
	user := new(User)
	err := request.ReadEntity(user)
	if err != nil {
		log.Printf(err.Error())
	}
	b := LoginUser(user)
	log.Printf("Login returned value %t", b)
}

func (u UserResource) findUser(request *restful.Request, response *restful.Response) {
	user := new(User)
	err := request.ReadEntity(user)
	if err != nil {
		log.Printf(err.Error())
	}
	x := DB{Database()}
	x.FindUser(user)

	//log.Printf("findUser returned value %t", b)
}
