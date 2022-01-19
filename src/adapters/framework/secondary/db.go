package secondary

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	arith "github.com/daibertdiego/golang_hex_architecture/src/adapters/core/arithmetic"
)

type Adapter struct {
	client  *mongo.Client
	db      *mongo.Database
	timeout time.Duration
}

func NewAdapter(timeout time.Duration) (*Adapter, error) {
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable. See\n\t https://docs.mongodb.com/drivers/go/current/usage-examples/#environment-variable")
	}
	database := os.Getenv("MONGODB_DB")
	if database == "" {
		log.Fatal("You must set your 'MONGODB_DB' environmental variable. See\n\t https://docs.mongodb.com/drivers/go/current/usage-examples/#environment-variable")
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalf("MongoDB connection fail: %v", err)
	}
	db := client.Database(database)
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatalf("MongoDB ping fail: %v", err)
	}
	return &Adapter{
		client:  client,
		db:      db,
		timeout: timeout,
	}, nil
}

func (mongoAdapter *Adapter) CloseDbConnection() {
	if err := mongoAdapter.client.Disconnect(context.TODO()); err != nil {
		log.Fatalf("MongoDB disconnect fail: %v", err)
	}
}

func (mongoAdapter *Adapter) AddToHistory(answer int32, operation string) (arith.Adapter, error) {
	ctx, cancel := context.WithTimeout(context.Background(), mongoAdapter.timeout)
	defer cancel()
	collection := mongoAdapter.db.Collection("arith_history")
	result, err := collection.InsertOne(
		ctx,
		arith.Adapter{
			ID:        primitive.NewObjectID(),
			Answer:    answer,
			Operation: operation,
			CreatedAt: time.Now(),
		},
	)
	if err != nil {
		return arith.Adapter{}, err
	}

	idFilter := bson.M{
		"_id": result.InsertedID,
	}

	formatedResult := arith.Adapter{}

	singleResult := collection.FindOne(ctx, idFilter)

	err = singleResult.Decode(&formatedResult)
	if err != nil {
		return formatedResult, err
	}

	if err != nil {
		return formatedResult, err
	}

	json, _ := json.Marshal(formatedResult)
	fmt.Println(string(json))

	return formatedResult, nil
}
