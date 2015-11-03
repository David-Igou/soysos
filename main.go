package main

import (
	"log"
	"net/http"

	"github.com/emicklei/go-restful"
)

func main() {
	cmap := make(map[string]CatFact)
	cmap["Lion"] = CatFact{"1", "Lion", "Lions have sharp teef :D"}

	cat := CatResource{cmap}
	root := RootResource{"Hello"}

	wsContainer := restful.NewContainer()
	wsContainer.Filter(GlobalLogging)

	cat.Register(wsContainer)
	root.Register(wsContainer)

	server := &http.Server{Addr: "localhost:8080", Handler: wsContainer}

	log.Printf("start listening on localhost:8080")

	log.Fatal(server.ListenAndServe())
}
