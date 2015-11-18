package main

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/emicklei/go-restful"
)

// WebService (post-process) Filter (as a struct that defines a FilterFunction)
func MeasureTime(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	now := time.Now()
	chain.ProcessFilter(req, resp)
	log.Printf("[SOYSOS (Timer)] %v\n", time.Now().Sub(now))
}

// WebService Filter
func WebserviceLogging(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	log.Printf("[SOYSOS (WebserviceLogging)] {%s}, %s - \"%s\" - %s, %d %d\n",
		strings.Split(req.Request.RemoteAddr, ":")[0],
		req.Request.Method,
		req.Request.URL,
		req.Request.Proto,
		resp.StatusCode(),
		resp.ContentLength(),
	)

	chain.ProcessFilter(req, resp)
}

// Global Filter
func GlobalLogging(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	log.Printf("[SOYSOS (GlobalLogging)] %s,%s\n", req.Request.Method, req.Request.URL)
	chain.ProcessFilter(req, resp)
}

// Route Filter (defines FilterFunction)
func RouteLogging(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	log.Printf("[SOYSOS(RouteLogging)] %s,%s\n", req.Request.Method, req.Request.URL)
	chain.ProcessFilter(req, resp)
}

func Sessions(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	token := req.HeaderParameter("token")
	if token == "" {
		log.Print("token was empty")
		resp.WriteErrorString(http.StatusUnauthorized, "You need to establish a session first!")
		return
	}
	x := DB{Database()}
	c, err := x.FindToken(token)
	if err != nil || c == false {
		log.Print("token was incorrect")
		resp.WriteErrorString(http.StatusUnauthorized, "Session does not exist!")
		return
	}
	chain.ProcessFilter(req, resp)
}
