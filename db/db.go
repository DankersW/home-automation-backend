package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
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
	Get()
}

func New(ctx context.Context) (Db, error) {
	mongoDb, err := newMongoDb(ctx, USER, PASSWORD, ADDR, PORT)
	if err != nil {
		return nil, fmt.Errorf("failed to create MongoDb instance, %s", err.Error())
	}

	d := &db{
		mongoDb: mongoDb,
	}
	return d, nil
}

func (d *db) Get() {
	d.mongoDb.Get("test", bson.D{})
	fmt.Println("get")
}
