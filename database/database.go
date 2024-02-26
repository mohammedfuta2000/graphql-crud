package database

import (
	"context"
	"log"
	"time"

	"github.com/mohammedfuta2000/graphql-crud/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var connectionString string = ""
var dbname string = "graphql-job-board"

type DB struct {
	client *mongo.Client
}

func Connect() *DB {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	if err != nil {
		log.Fatal(err)
	}

	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal(err)
	}

	return &DB{
		client: client,
	}
}

func (db *DB) GetJob(id string) *model.JobListing {
	jobCollection := db.client.Database(dbname).Collection("jobs")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_id, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": _id}
	var jobListing model.JobListing
	if err := jobCollection.FindOne(ctx, filter).Decode(&jobListing); err != nil {
		log.Fatal(err)
	}

	return &jobListing
}

func (db *DB) GetJobs() []*model.JobListing {
	jobCollection := db.client.Database(dbname).Collection("jobs")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	var jobListings []*model.JobListing

	cursor, err := jobCollection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	if err = cursor.All(context.TODO(), &jobListings); err != nil {
		panic(err)
	}
	return jobListings
}

func (db *DB) CreateJobListing(jobInfo model.CreateJobListingInput) *model.JobListing {
	jobCollection := db.client.Database(dbname).Collection("jobs")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	inserted, err := jobCollection.InsertOne(ctx, bson.M{
		"title":       jobInfo.Title,
		"description": jobInfo.Description,
		"company":     jobInfo.Company,
		"url":         jobInfo.URL,
	})
	if err != nil {
		log.Fatal(err)
	}
	insertedID := inserted.InsertedID.(primitive.ObjectID).Hex()
	returnJobListing := model.JobListing{ID: insertedID,
		Title:       jobInfo.Title,
		Description: jobInfo.Company,
		Company:     jobInfo.Company,
		URL:         jobInfo.URL}
	return &returnJobListing
}

func (db *DB) UpdateJobListing(jobId string, jobInfo model.UpdateJobListingInput) *model.JobListing {
	var jobListing model.JobListing
	return &jobListing
}

func (db *DB) DeleteJobListing(jobId string) *model.DeleteJobResponse {
	return &model.DeleteJobResponse{DeleteJobID: jobId}
}
