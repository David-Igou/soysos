package main

import (
	"flag"
	"log"
	"net/http"
	"os"

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
		log.Fatalf("[soysos][error] Unable to read properties:%v\n", err)
	}

	addr := props.MustGet("http.server.host") + ":" + props.MustGet("http.server.port")
	basePath := "http://" + addr
	SwaggerPath = props.GetString("swagger.path", "")

	cmap := make(map[string]CatFact)
	cmap["Lion"] = CatFact{"", "Lion", "Lions have sharp teef :D"}

	cat := CatResource{cmap}
	root := RootResource{SwaggerPath}
	user := UserResource{}

	wsContainer := restful.NewContainer()
	wsContainer.Filter(GlobalLogging)
	wsContainer.Filter(enableCORS)

	restful.TraceLogger(log.New(os.Stdout, "[restful] ", log.LstdFlags|log.Lshortfile))
	cat.Register(wsContainer)
	root.Register(wsContainer)
	user.Register(wsContainer)

	config := swagger.Config{
		WebServices:     wsContainer.RegisteredWebServices(),
		WebServicesUrl:  basePath,
		ApiPath:         "/apidocs.json",
		SwaggerPath:     SwaggerPath,
		SwaggerFilePath: props.GetString("swagger.file.path", "")}

	swagger.RegisterSwaggerService(config, wsContainer)

	cors := restful.CrossOriginResourceSharing{
		AllowedHeaders: []string{"Accept", "Authorization"},
		AllowedMethods: []string{"GET"},
		CookiesAllowed: false,
		Container:      wsContainer}

	//wsContainer.Filter(wsContainer.OPTIONSFilter)
	wsContainer.Filter(cors.Filter)

	server := &http.Server{Addr: addr, Handler: wsContainer}

	log.Printf("start listening on %s", basePath)

	log.Fatal(server.ListenAndServe())
}
