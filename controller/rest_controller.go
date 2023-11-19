package controller

import (
	"bearLinks/service"
	"github.com/gorilla/mux"
	"net/http"
)

func Init() {
	httpRouter := mux.NewRouter().StrictSlash(true)
	httpRouter.HandleFunc("/generateShortLink", service.GenerateShortLink).Methods("POST")
	httpRouter.HandleFunc("/deleteShortLink", service.DeleteShortLink).Methods("DELETE")
	httpRouter.HandleFunc("/{shortLink}", service.Redirect).Methods("GET")
	http.ListenAndServe(":80", httpRouter)
}
