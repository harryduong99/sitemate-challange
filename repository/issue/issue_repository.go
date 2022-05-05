package issue_repository

import (
	"github/harryduong99/sitemate/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IssueRepo struct{}

var issueRepo = &IssueRepo{}

func (issueRepo *IssueRepo) Get(id primitive.ObjectID) models.Issue {

}

func (issueRepo *IssueRepo) Store(issue models.Issue) error {

}

func (issueRepo *IssueRepo) Delete(id primitive.ObjectID) bool {

}

func (issueRepo *IssueRepo) Update(issue models.Issue) bool {

}
