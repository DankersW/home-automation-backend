package db

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	DB = "iot_db"
)

type mongoDb struct {
	dbi *mongo.Database
	ctx context.Context
}
type MongoDb interface {
	Get(string, primitive.D, *options.FindOptions) (*mongo.Cursor, error)
	ListCollectionNames() ([]string, error)
	TimestampBetween(int, int) primitive.D
}

func newMongoDb(ctx context.Context, usr string, pwd string, addr string, port int) (MongoDb, error) {
	mongoUri := fmt.Sprintf("mongodb://%s:%s@%s:%d/", usr, pwd, addr, port)
	dbi, err := connectToDb(ctx, mongoUri)
	if err != nil {
		return nil, err
	}

	m := &mongoDb{
		dbi: dbi,
		ctx: ctx,
	}
	return m, nil
}

func connectToDb(ctx context.Context, uri string) (*mongo.Database, error) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("could not make contact with mongoDB, %s", err.Error())
	}
	return client.Database(DB), nil
}

func (m *mongoDb) Get(collectionName string, filter primitive.D, options *options.FindOptions) (*mongo.Cursor, error) {
	collection := m.dbi.Collection(collectionName)
	data, err := collection.Find(m.ctx, filter, options)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (m *mongoDb) ListCollectionNames() ([]string, error) {
	filter := bson.D{{}}
	names, err := m.dbi.ListCollectionNames(m.ctx, filter)
	if err != nil {
		return nil, err
	}
	return names, nil
}

// Generates a timestamp filter between a start day and nr of days in the past
func (m *mongoDb) TimestampBetween(nrDays int, startDay int) primitive.D {
	filter := bson.D{
		primitive.E{
			Key: "timestamp", Value: bson.D{primitive.E{
				Key: "$gte", Value: primitive.NewDateTimeFromTime(time.Now().AddDate(0, 0, -nrDays))}}},
		primitive.E{
			Key: "timestamp", Value: bson.D{primitive.E{
				Key: "$lte", Value: primitive.NewDateTimeFromTime(time.Now().AddDate(0, 0, -startDay))}}},
	}
	return filter
}
