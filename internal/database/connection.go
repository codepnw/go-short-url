package database

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/codepnw/go-short-url/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type manager struct {
	connection *mongo.Client
	ctx        context.Context
	cancel     context.CancelFunc
}

var Mgr Manager

type Manager interface {
	Insert(any, string) (any, error)
	GetUrlFromCode(string, string) (models.UrlDB, error)
}

func ConnectDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	uri := fmt.Sprintf("mongodb://%s", os.Getenv("DB_HOST"))
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}

	Mgr = &manager{connection: client, ctx: ctx, cancel: cancel}
}
