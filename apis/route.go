package apis

import (
	"log"
	"net/http"

	"github/harryduong99/sitemate/apis/rest_issues"

	"github.com/gorilla/mux"
)

func InitRoutes() {
	r := mux.NewRouter()

	apiRoute := r.PathPrefix("/api/").Subrouter()
	apiRoute.HandleFunc("/v1/issue", rest_issues.Create).Methods("POST")

	log.Fatal(http.ListenAndServe(":6868", r))
}
