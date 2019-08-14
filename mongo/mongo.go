package mongo

import (
	"context"
	"log"
	"time"

	"github.com/Khamliuk/testsCI/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type db struct {
	*mongo.Collection
}

// New creates new connection to db.
func New() (*db, error) {
	opts := options.Client().ApplyURI("mongodb://localhost:27017/test?ssl=false")
	err := opts.Validate()
	if err != nil {
		return nil, err
	}
	c, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err = c.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}
	database := c.Database("test")
	return &db{
		Collection: database.Collection("persons"),
	}, nil
}

// Create creates new person in db.
func (d db) Create(ctx context.Context, req model.Person) (*model.Person, error) {
	ins, err := d.Collection.InsertOne(ctx, req)
	if err != nil {
		return nil, err
	}
	log.Println(ins.InsertedID)
	req.ID = ins.InsertedID.(primitive.ObjectID)
	return &req, err
}

// List returns list of person from db.
func (d db) List(ctx context.Context) ([]model.Person, error) {
	cur, err := d.Collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	persons := make([]model.Person, 0)
	for cur.Next(ctx) {
		var person model.Person
		err = cur.Decode(&person)
		if err != nil {
			return nil, err
		}
		persons = append(persons, person)
	}
	return persons, nil
}

// Delete creates new person in db.
func (d db) Update(ctx context.Context, req model.Person) error {
	_, err := d.Collection.ReplaceOne(ctx, bson.M{"_id": req.ID}, req)
	if err != nil {
		return err
	}
	return nil
}

// Delete creates new person in db.
func (d db) Delete(ctx context.Context, id string) error {
	_id, err := primitive.ObjectIDFromHex(id)
	_, err = d.Collection.DeleteOne(ctx, bson.M{"_id": _id})
	if err != nil {
		return err
	}
	return nil
}
