package controllers

import (
	"fmt"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func test(w http.ResponseWriter, r *http.Request) {
	responseBody, err := json.Marshal("test")
	w.Header().Set("Content-Type", "application/json")
	w.Write(responseBody)
}

func startWebServer() error {
	router := mux.NewRouter().StrictSlash(true)
	router.Handle("/api/test", test).Method("GET")
	return http.ListenAndServe(fmt.Sprintf(":%d", 8080), router)
}
