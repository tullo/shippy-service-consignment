package main

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// CreateClient ...
func CreateClient(ctx context.Context, uri string, retry int32) (*mongo.Client, error) {
	o := options.Client()
	o.ApplyURI(uri)
	conn, err := mongo.Connect(o)
	if err != nil {
		return nil, err
	}
	if err := conn.Ping(ctx, nil); err != nil {
		if retry >= 3 {
			return nil, err
		}
		retry = retry + 1
		time.Sleep(time.Second * 2)
		return CreateClient(ctx, uri, retry)
	}

	return conn, err
}
