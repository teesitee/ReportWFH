package repository

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoRepository struct {
	Collection *mongo.Collection
}

type InfoCreate struct {
	ID        string `bson:"id"`
	FirstName string `bson:"firstname"`
	LastName  string `bson:"lastname"`
	Detail    string `bson:"detail"`
}

func (r MongoRepository) GetAll() ([]InfoCreate, error) {
	cursor, err := r.Collection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var Info []InfoCreate
	if err = cursor.All(context.TODO(), &Info); err != nil {
		log.Fatal(err)
	}
	return Info, err
}

func (r MongoRepository) InsertInfo(id, fname, lname, detail string) error {
	_, err := r.Collection.InsertOne(context.TODO(), InfoCreate{
		ID:        id,
		FirstName: fname,
		LastName:  lname,
		Detail:    detail,
	})

	return err
}

func (r MongoRepository) Deleteinfo(id string) (int64, error) {
	delcount, err := r.Collection.DeleteMany(context.TODO(), bson.M{"id": id})
	return delcount.DeletedCount, err
}

func (r MongoRepository) FindInfo(id string) ([]InfoCreate, error) {
	cur, err := r.Collection.Find(context.TODO(), bson.M{"id": id})
	var Info []InfoCreate
	if err2 := cur.All(context.TODO(), &Info); err != nil {
		log.Fatal(err2)
	}
	return Info, nil
}
