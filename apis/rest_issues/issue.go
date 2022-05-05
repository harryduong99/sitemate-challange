package rest_issues

import (
	"encoding/json"
	"github/harryduong99/sitemate/apis"
	"github/harryduong99/sitemate/models"
	issue_repository "github/harryduong99/sitemate/repository/issue"

	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Find(response http.ResponseWriter, request *http.Request) {
	var params = mux.Vars(request)

	id, err := primitive.ObjectIDFromHex(params["id"])
	result := issue_repository.IssueRepository.Get(id)
	if err != nil {
		apis.ResponseWithError(response, http.StatusBadRequest, "Something went wrong!")
		return
	}

	apis.ResponseWithJSON(response, http.StatusOK, result)
}

func Create(response http.ResponseWriter, request *http.Request) {
	var issue models.Issue
	_ = json.NewDecoder(request.Body).Decode(&issue)

	err := issue_repository.IssueRepository.Store(issue)
	if err != nil {
		apis.ResponseWithError(response, http.StatusBadRequest, "Something went wrong!")
		return
	}

	apis.ResponseWithJSON(response, http.StatusOK, "Create successfully")
}

func Update(response http.ResponseWriter, request *http.Request) {
	var params = mux.Vars(request)

	id, err := primitive.ObjectIDFromHex(params["id"])
	result := issue_repository.IssueRepository.Get(id)
	if err != nil {
		apis.ResponseWithError(response, http.StatusBadRequest, "Something went wrong!")
		return
	}

	apis.ResponseWithJSON(response, http.StatusOK, result)
}

func Delete(response http.ResponseWriter, request *http.Request) {
	// get params
	var params = mux.Vars(request)

	// string to primitve.ObjectID
	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		return
	}

	result := issue_repository.IssueRepository.Delete(id)

	apis.ResponseWithJSON(response, http.StatusOK, result)
}
