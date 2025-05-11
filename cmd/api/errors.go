package main

import (
	"log"
	"net/http"
)

func (* application) internalServerError(w http.ResponseWriter, r *http.Request, err error){
	log.Printf("internal server error: %s path: %s error: %s", r.URL.Path, r.Method, err.Error())
	writeJSONError(w, http.StatusInternalServerError, "the server encountered a problem")
}

func (* application)badRequestResposne(w http.ResponseWriter, r *http.Request, err error){
	log.Printf("bad request: %s path: %s error: %s", r.URL.Path, r.Method, err.Error())
	writeJSONError(w, http.StatusBadRequest, err.Error())
}

func (* application) notFound(w http.ResponseWriter, r *http.Request){
	log.Printf("not found: %s path: %s", r.URL.Path, r.Method)
	writeJSONError(w, http.StatusNotFound, "the requested resource could not be found")
}