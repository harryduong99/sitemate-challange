package issue_repository

import (
	"context"
	"github/harryduong99/sitemate/config"
	"github/harryduong99/sitemate/databasedriver"
	"github/harryduong99/sitemate/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IssueRepo struct{}

var IssueRepository = &IssueRepo{}

func (issueRepo *IssueRepo) Get(id primitive.ObjectID) models.Issue {
	collection := databasedriver.Mongo.ConnectCollection(config.DB_NAME, config.COL_ISSUE)
	var issue models.Issue
	data := collection.FindOne(context.Background(), bson.M{"_id": id})
	data.Decode(&issue)

	return issue
}

func (issueRepo *IssueRepo) Store(issue models.Issue) error {
	collection := databasedriver.Mongo.ConnectCollection(config.DB_NAME, config.COL_ISSUE)

	bbytes, _ := bson.Marshal(issue)

	_, err := collection.InsertOne(context.Background(), bbytes)
	if err != nil {
		return err
	}

	return nil
}

func (issueRepo *IssueRepo) Delete(id primitive.ObjectID) bool {
	collection := databasedriver.Mongo.ConnectCollection(config.DB_NAME, config.COL_ISSUE)

	filter := bson.M{"_id": id}

	_, err := collection.DeleteOne(context.TODO(), filter)
	return err == nil
}

func (issueRepo *IssueRepo) Update(issue models.Issue, id primitive.ObjectID) bool {
	collection := databasedriver.Mongo.ConnectCollection(config.DB_NAME, config.COL_ISSUE)
	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", bson.D{{"title", issue.Title}}}}
	_, err := collection.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		return false
	}

	return true
}
