package db

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IMongoDB interface {
	Find(i interface{}, q bson.M) error
	FindAll(res interface{}, q bson.M, sort bson.M) error
	InsertOne(e interface{}) error
	FindAndUpdate(i interface{}, q bson.M, update bson.M, opts *options.FindOneAndUpdateOptions) error
	DeleteOne(i interface{}, q bson.M) error
}
