package main

import (
	"log"
	"time"

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
	log.Printf("[SOYSOS (logger)] %s,%s\n", req.Request.Method, req.Request.URL)
	chain.ProcessFilter(req, resp)
}

// Global Filter
func GlobalLogging(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	log.Printf("[SOYSOS (log)] %s,%s\n", req.Request.Method, req.Request.URL)
	chain.ProcessFilter(req, resp)
}

// Route Filter (defines FilterFunction)
func RouteLogging(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	log.Printf("[SOYSOS(logger)] %s,%s\n", req.Request.Method, req.Request.URL)
	chain.ProcessFilter(req, resp)
}
