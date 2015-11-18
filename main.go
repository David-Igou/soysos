package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/emicklei/go-restful"
	"github.com/emicklei/go-restful/swagger"
	"github.com/magiconair/properties"
)

var (
	props          *properties.Properties
	propertiesFile = flag.String("config", "soysos.properties", "the configuration file")
	SwaggerPath    string
)

func main() {
	flag.Parse()

	log.Printf("loading configuration from [%s]", *propertiesFile)
	var err error
	if props, err = properties.LoadFile(*propertiesFile, properties.UTF8); err != nil {
		log.Fatalf("[mora][error] Unable to read properties:%v\n", err)
	}

	cmap := make(map[string]CatFact)
	cmap["Lion"] = CatFact{"", "Lion", "Lions have sharp teef :D"}

	cat := CatResource{cmap}
	root := RootResource{}
	user := UserResource{}

	wsContainer := restful.NewContainer()
	wsContainer.Filter(GlobalLogging)

	cat.Register(wsContainer)
	root.Register(wsContainer)
	user.Register(wsContainer)
	SwaggerPath = props.GetString("swagger.path", "")

	config := swagger.Config{
		WebServices:     wsContainer.RegisteredWebServices(),
		WebServicesUrl:  "http://localhost:8008",
		ApiPath:         "/apidocs.json",
		SwaggerPath:     SwaggerPath,
		SwaggerFilePath: "./dist"}

	swagger.RegisterSwaggerService(config, wsContainer)

	server := &http.Server{Addr: "localhost:8008", Handler: wsContainer}

	log.Printf("start listening on localhost:8008")

	log.Fatal(server.ListenAndServe())
}
