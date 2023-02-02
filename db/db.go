package db

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var Client *mongo.Client

const After = options.After

var C map[string]*mongo.Collection

var database string

type M = primitive.M

type Options = options.FindOneAndUpdateOptions

var Mongo IMongoDB

type MongoService struct{}

func Connect() error {
	Mongo = &MongoService{}
	database = "Students"
	C = make(map[string]*mongo.Collection)
	Ctx := context.Background()
	defer Ctx.Done()
	var err error
	Client, err = mongo.Connect(Ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return err
	}
	err = Client.Ping(Ctx, readpref.Primary())
	if err != nil {
		return err
	}
	fmt.Println("Db Connected...")
	return nil
}

func (m *MongoService) Find(i interface{}, q bson.M) error {
	ctx := context.Background()
	defer ctx.Done()
	coll := Col(i)
	err := coll.FindOne(ctx, q).Decode(i)
	return err
}

func (m *MongoService) InsertOne(e interface{}) error {
	ctx := context.Background()
	defer ctx.Done()
	cl := Col(e)
	_, err := cl.InsertOne(ctx, e)
	if err != nil {
		return err
	}
	return nil
}
func (m *MongoService) DeleteOne(i interface{}, q bson.M) error {
	ctx := context.Background()
	defer ctx.Done()
	coll := Col(i)
	_, err := coll.DeleteOne(ctx, q)
	return err
}

func (m *MongoService) FindAll(res interface{}, q bson.M, sort bson.M) error {
	ctx := context.Background()
	defer ctx.Done()
	if !IsPtr(res) {
		return errors.New("you must pass in a pointer")
	}
	opts := options.Find()
	opts.SetSort(sort)
	coll := Col(res)
	cursor, err := coll.Find(ctx, q, opts)
	if err != nil {
		return err
	}
	cursor.All(ctx, res)
	return err
}

func (m *MongoService) FindAndUpdate(i interface{}, q bson.M, update bson.M, opts *options.FindOneAndUpdateOptions) error {
	ctx := context.Background()
	defer ctx.Done()
	coll := Col(i)
	errr := coll.FindOneAndUpdate(ctx, q, update, opts).Decode(i)
	if errr != nil {
		log.Println("Mongo Update Error : ", errr)
	}
	return errr
}
func Col(e interface{}) *mongo.Collection {
	cname := typeName(e)
	res, ok := C[cname]
	if !ok {
		db := Client.Database(database)
		log.Printf("Type:%s", cname)
		r2 := db.Collection(cname)
		res = r2
		C[cname] = r2
	}
	return res
}

func typeName(i interface{}) string {
	t := reflect.TypeOf(i)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		if t.Kind() == reflect.Ptr {
			t = t.Elem()
		}
	}
	if isSlice(t) {
		t = t.Elem()
		if t.Kind() == reflect.Ptr {
			t = t.Elem()
		}
	}
	return t.Name()
}

func isSlice(t reflect.Type) bool {
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t.Kind() == reflect.Slice
}

func IsPtr(i interface{}) bool {
	return reflect.ValueOf(i).Kind() == reflect.Ptr
}
