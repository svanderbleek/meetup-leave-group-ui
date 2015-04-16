package main

import (
	"./meetup"
	"html/template"
	"net/http"
)

var groupTemplate = template.Must(template.ParseFiles("meetup/groups.html"))

func listGroups(response http.ResponseWriter) {
	groups := meetup.Groups()
	groupTemplate.Execute(response, groups)
}

func deleteGroup(response http.ResponseWriter, request *http.Request) {
	id := request.FormValue("id")
	if meetup.LeaveGroup(id) {
		response.WriteHeader(http.StatusOK)
	} else {
		http.Error(response, "unsuccessful", http.StatusBadRequest)
	}
}

func groupsHandler(response http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "GET":
		listGroups(response)
	case "DELETE":
		deleteGroup(response, request)
	}
}

func main() {
	routes := http.NewServeMux()
	routes.HandleFunc("/groups", groupsHandler)
	server := &http.Server{
		Addr:    ":8080",
		Handler: routes,
	}
	server.ListenAndServe()
}
