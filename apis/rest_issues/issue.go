package rest_issues

import (
	"encoding/json"
	"github/harryduong99/sitemate/apis"
	"github/harryduong99/sitemate/issue_repository"
	"github/harryduong99/sitemate/models"
	"net/http"
)

func Find(response http.ResponseWriter, request *http.Request) {

}

func GetAll(response http.ResponseWriter, request *http.Request) {

}

func Create(response http.ResponseWriter, request *http.Request) {
	var issue models.Issue
	_ = json.NewDecoder(request.Body).Decode(&issue)

	err := issue_repository.Store(issue)

	if err != nil {
		apis.ResponseWithError(response, http.StatusBadRequest, "Something went wrong!")
		return
	}

	apis.ResponseWithJSON(response, http.StatusOK, "Create successfully")
}

func Update(response http.ResponseWriter, request *http.Request) {

}

func Delete(response http.ResponseWriter, request *http.Request) {

}
