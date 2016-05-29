package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"TodoIndex",
		"GET",
		"/vms",
		VMList,
	},
	Route{
		"VmStart",
		"POST",
		"/vms/start/{nameOrID}",
		VMStart,
	},
	Route{
		"VmStop",
		"POST",
		"/vms/stop/{nameOrID}",
		VMStop,
	},
}
