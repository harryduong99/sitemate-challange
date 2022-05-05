package issue_repository

import (
	"context"
	"github/harryduong99/sitemate/config"
	"github/harryduong99/sitemate/databasedriver"
	"github/harryduong99/sitemate/models"

	"go.mongodb.org/mongo-driver/bson"
)

type IssueRepo struct{}

var IssueRepository = &IssueRepo{}

// func (issueRepo *IssueRepo) Get(id primitive.ObjectID) models.Issue {

// }

func (issueRepo *IssueRepo) Store(issue models.Issue) error {
	collection := databasedriver.Mongo.ConnectCollection(config.DB_NAME, config.COL_ISSUE)

	bbytes, _ := bson.Marshal(issue)

	_, err := collection.InsertOne(context.Background(), bbytes)
	if err != nil {
		return err
	}

	return nil
}

// func (issueRepo *IssueRepo) Delete(id primitive.ObjectID) bool {

// }

// func (issueRepo *IssueRepo) Update(issue models.Issue) bool {

// }
