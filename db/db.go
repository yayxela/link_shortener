package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"link_shortener/config"
)

type DB struct {
	client   *mongo.Client
	database string
}

func (db *DB) GetDatabase() *mongo.Database {
	return db.client.Database(db.database)
}

func (db *DB) Find(ctx context.Context, collection string, model interface{}, filter primitive.M) error {
	return db.GetDatabase().Collection(collection).FindOne(ctx, filter).Decode(model)
}

func (db *DB) Insert(ctx context.Context, name string, model interface{}) (err error) {
	_, err = db.GetDatabase().Collection(name).InsertOne(ctx, model)
	return
}

func (db *DB) DeleteById(ctx context.Context, collection string, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = db.GetDatabase().Collection(collection).DeleteOne(ctx, bson.M{"_id": oid})
	return err
}

func New(ctx context.Context, config *config.Config) (*DB, error) {
	connection := options.Client().ApplyURI(config.GetConnectionString())
	client, err := mongo.Connect(ctx, connection)
	if err != nil {
		return nil, err
	}
	if err = client.Ping(ctx, nil); err != nil {
		return nil, err
	}
	return &DB{
		client,
		config.DBName,
	}, nil
}
