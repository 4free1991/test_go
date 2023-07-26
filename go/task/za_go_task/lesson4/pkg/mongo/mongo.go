package mongo

import (
	"app/lesson4/config"
	"app/lesson4/pkg/shutdown"
	"context"
	"github.com/qiniu/qmgo"
)

var (
	MongoClient *qmgo.Client
	MongoDB     *qmgo.Database
)

func InitMongoClient() {
	var url = config.GetConfig().Mongodb.URL
	var dbName = config.GetConfig().Mongodb.Dbname
	ctx := context.Background()
	client, err := qmgo.NewClient(ctx, &qmgo.Config{Uri: url})
	if err != nil {
		panic(err)
	}

	db := client.Database(dbName)
	MongoClient = client
	MongoDB = db
	shutdown.Add(func() {
		client.Close(context.Background())
	})
}
