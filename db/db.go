package db

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	USER     = "admin"
	PASSWORD = "mongo_admin_iot"
	ADDR     = "localhost"
	PORT     = 27017
)

type db struct {
	mongoDb MongoDb
}
type Db interface {
	Get(string) (*mongo.Cursor, error)
	GetWithFilter(string, primitive.D) (*mongo.Cursor, error)
	FetchCollectionNames() ([]string, error)
	TimestampBetween(int, int) primitive.D
}

func New(ctx context.Context) (Db, error) {
	mongoDb, err := newMongoDb(ctx, USER, PASSWORD, ADDR, PORT)
	if err != nil {
		return nil, fmt.Errorf("failed to create MongoDb instance, %s", err.Error())
	}
	log.Info("Successfully connected to MongoDb")

	d := &db{
		mongoDb: mongoDb,
	}
	return d, nil
}

func (d *db) get(collectionName string, filter primitive.D, options *options.FindOptions) (*mongo.Cursor, error) {
	return d.mongoDb.Get(collectionName, filter, options)
}

func (d *db) Get(collectionName string) (*mongo.Cursor, error) {
	filter := bson.D{}
	options := options.Find()
	return d.get(collectionName, filter, options)
}

func (d *db) GetWithFilter(collectionName string, filter primitive.D) (*mongo.Cursor, error) {
	options := options.Find()
	return d.get(collectionName, filter, options)
}

func (d *db) FetchCollectionNames() ([]string, error) {
	return d.mongoDb.ListCollectionNames()
}

func (d *db) TimestampBetween(nrDays int, startDay int) primitive.D {
	return d.mongoDb.TimestampBetween(nrDays, startDay)
}
