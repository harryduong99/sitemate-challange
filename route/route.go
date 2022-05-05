package route

import (
	"github/harryduong99/sitemate/apis/rest_issues"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func InitRoutes() {
	r := mux.NewRouter()

	apiRoute := r.PathPrefix("/api/").Subrouter()
	apiRoute.HandleFunc("/v1/issue", rest_issues.Create).Methods("POST")
	apiRoute.HandleFunc("/v1/issue", rest_issues.Find).Methods("GET")
	apiRoute.HandleFunc("/v1/issue", rest_issues.Delete).Methods("DELETE")
	apiRoute.HandleFunc("/v1/issue", rest_issues.Update).Methods("PUT")

	log.Fatal(http.ListenAndServe(":6868", r))
}
