package main

import (
	"log"
	"net/http"

	"github.com/emicklei/go-restful"
)

func main() {
	cmap := make(map[string]CatFact)
	cmap["Lion"] = CatFact{"", "Lion", "Lions have sharp teef :D"}

	cat := CatResource{cmap}
	root := RootResource{"Hello"}
	user := UserResource{}

	wsContainer := restful.NewContainer()
	wsContainer.Filter(GlobalLogging)

	cat.Register(wsContainer)
	root.Register(wsContainer)
	user.Register(wsContainer)

	server := &http.Server{Addr: "localhost:8008", Handler: wsContainer}

	log.Printf("start listening on localhost:8008")

	log.Fatal(server.ListenAndServe())
}
