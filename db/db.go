package db

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client     *mongo.Client
	col        *mongo.Collection
	database   string
	collection string
	username   string
	password   string
	host       string
	port       string
	authSource string
	ctx, _     = context.WithTimeout(context.Background(), 10*time.Second)
)

// CrawlabItem 插入mongo的struct需要继承它
type CrawlabItem struct {
	TaskID string `json:"task_id"`
}

// Init 获取环境变量
func Init() {
	database = os.Getenv("CRAWLAB_MONGO_DB")
	collection = os.Getenv("CRAWLAB_COLLECTION")
	username = os.Getenv("CRAWLAB_MONGO_USERNAME")
	password = os.Getenv("CRAWLAB_MONGO_PASSWORD")
	host = os.Getenv("CRAWLAB_MONGO_HOST")
	if len(host) == 0 {
		host = "localhost"
	}
	port = os.Getenv("CRAWLAB_MONGO_PORT")
	if len(port) == 0 {
		port = "27107"
	}
	authSource = os.Getenv("CRAWLAB_MONGO_AUTHSOURCE")
	if len(authSource) == 0 && len(username) > 0 {
		authSource = "admin"
	}
}

// NewCollection 新客户端
func NewCollection() (*mongo.Collection, context.Context, error) {
	var err error

	if client == nil {
		Init()
		applyURI := fmt.Sprintf(`mongodb://%s:%s@%s:%s/%s?authSource=%s`,
			username, password, host, port, database, authSource)
		_ = applyURI
		client, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	}

	if col == nil {
		col = client.Database(database).Collection(collection)
	}

	return col, ctx, err
}

// Close 关闭数据库
// TODO: 好像没有关闭
func Close() {
	if client == nil {
		return
	}
}
