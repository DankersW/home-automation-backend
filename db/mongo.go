package db

import (
	"context"
	"fmt"

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
	Get(collectionName string, filter primitive.D) (*mongo.Cursor, error)
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

func (m *mongoDb) Get(collectionName string, filter primitive.D) (*mongo.Cursor, error) {
	collection := m.dbi.Collection(collectionName)
	data, err := collection.Find(m.ctx, filter)
	if err != nil {
		return nil, err
	}
	return data, nil
}
