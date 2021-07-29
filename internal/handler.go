package internal

import (
	"log"
	"net/http"
	"test_task/internal/handlers"
)

func InitHandlers() {
	http.HandleFunc("/first/entity/", firstEntity)
	http.HandleFunc("/second/entity/", secondEntity)
	http.HandleFunc("/entity/all", entityAll)
	http.HandleFunc("/first/entity/new", firstEntityNew)
	http.HandleFunc("/second/entity/new", secondEntityNew)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

const (
	GET = "GET"
	POST = "POST"
)

func firstEntity(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case GET:
		handlers.HandlerShowFirst(w, r)
	case POST:
		handlers.HandlerUpdateFirst(w, r)
	}
}


func secondEntity(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case GET:
		handlers.HandlerShowSecond(w, r)
	case POST:
		handlers.HandlerUpdateSecond(w, r)
	}
}

func entityAll(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case GET:
		handlers.HandlerShowAll(w, r)
	}
}

func firstEntityNew(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case POST:
		handlers.HandlerAddFirst(w, r)
	}
}

func secondEntityNew(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case POST:
		handlers.HandlerAddSecond(w, r)
	}
}