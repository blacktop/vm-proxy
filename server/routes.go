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
		"VBoxList",
		"GET",
		"/virtualbox/list",
		VBoxList,
	},
	Route{
		"VBoxStatus",
		"GET",
		"/virtualbox/status/{nameOrID}",
		VBoxStatus,
	},
	Route{
		"VBoxStart",
		"GET",
		"/virtualbox/start/{nameOrID}",
		VBoxStart,
	},
	Route{
		"VBoxStop",
		"GET",
		"/virtualbox/stop/{nameOrID}",
		VBoxStop,
	},
	Route{
		"VBoxSnapshotRestoreCurrent",
		"GET",
		"/virtualbox/snapshot/restorecurrent/{nameOrID}",
		VBoxSnapshotRestoreCurrent,
	},
	Route{
		"VBoxSnapshotRestore",
		"GET",
		"/virtualbox/snapshot/{nameOrID}/restore/{snapShot}",
		VBoxSnapshotRestore,
	},
}
