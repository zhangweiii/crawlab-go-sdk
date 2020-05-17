package main

import crawlabgo "github.com/zhangweiii/crawlab-go-sdk"

// Item 示例结构体
// bson 是mongo中列别名，主要排重时候用
type Item struct {
	TaskID string `bson:"task_id"` // 必须要有用于爬虫平台识别结果
	Name   string `bson:"name"`
	Age    int    `bson:"age"`
}

func main() {
	for i := 0; i < 1000; i++ {
		crawlabgo.SaveItem(&Item{
			Name: "crawlabgo",
			Age:  i,
		})
	}
}
