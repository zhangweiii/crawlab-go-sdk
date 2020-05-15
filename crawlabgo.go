package crawlabgo

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/zhangweiii/crawlab-go-sdk/db"
	"gopkg.in/mgo.v2/bson"
)

// SaveItem 保存文档到数据库
func SaveItem(item interface{}) error {
	col, ctx, err := db.NewCollection()
	if err != nil {
		panic(err.Error())
	}
	if col == nil {
		panic("mongo collection is nil")
	}

	elem := reflect.ValueOf(item).Elem()
	elem.FieldByName("TaskID").
		SetString(os.Getenv("CRAWLAB_TASK_ID"))
	isDedup, err := strconv.ParseBool(os.Getenv("CRAWLAB_IS_DEDUP"))
	if err != nil {
		isDedup = false
	}
	dedupField := os.Getenv("CRAWLAB_DEDUP_FIELD")
	dedupMethod := os.Getenv("CRAWLAB_DEDUP_METHOD")

	if isDedup {
		if dedupMethod == "overwrite" {
			key := getTagName(dedupField, "bson", item)
			value := elem.FieldByName(dedupField).String()
			delete := bson.M{
				key: value,
			}
			col.DeleteMany(
				ctx,
				delete,
			)
			_, err := col.InsertOne(ctx, item)
			return err
		} else if dedupMethod == "ignore" {
			_, err := col.InsertOne(ctx, item)
			return err
		} else {
			_, err := col.InsertOne(ctx, item)
			return err
		}
	} else {
		_, err := col.InsertOne(ctx, item)
		return err
	}
}
func getTagName(field, key string, s interface{}) (fieldname string) {
	rt := reflect.TypeOf(s)
	if rt.Kind() == reflect.Ptr {
		rt = rt.Elem()
	}

	if rt.Kind() != reflect.Struct {
		panic(fmt.Sprintf("bad type: %v", rt.Kind()))
	}
	for i := 0; i < rt.NumField(); i++ {
		f := rt.Field(i)
		v := strings.Split(f.Tag.Get(key), ",")[0] // use split to ignore tag "options" like omitempty, etc.
		if f.Name == field {
			return v
		}
	}
	return ""
}

// Close 关闭连接
func Close() {
	_, _, err := db.NewCollection()
	if err != nil {
		return
	}
}
