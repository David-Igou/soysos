package main

import (
	"log"
	"time"
	"strings"

	"github.com/emicklei/go-restful"
)

// WebService (post-process) Filter (as a struct that defines a FilterFunction)
func MeasureTime(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	now := time.Now()
	chain.ProcessFilter(req, resp)
	log.Printf("[SOYSOS (timer)] %v\n", time.Now().Sub(now))
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
