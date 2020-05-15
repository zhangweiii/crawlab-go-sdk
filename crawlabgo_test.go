package crawlabgo

import (
	"os"
	"testing"
)

func setCrawlabEnv() {
	os.Setenv("CRAWLAB_TASK_ID", "377e3d13-d650-4098-be3e-37c852393b54")
	os.Setenv("CRAWLAB_IS_DEDUP", "false")
	os.Setenv("CRAWLAB_DEDUP_METHOD", "ignore")
	os.Setenv("CRAWLAB_MONGO_DB", "crawlab_test")
	os.Setenv("CRAWLAB_COLLECTION", "results_gotest")
	os.Setenv("CRAWLAB_DEDUP_METHOD", "ignore")
}

type Item struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	TaskID string `json:"task_id"`
}

func TestSaveItem(t *testing.T) {
	setCrawlabEnv()

	t.Run("save item", func(t *testing.T) {
		SaveItem(&Item{Name: "zhangweiii", Age: 1})
	})

	t.Run("save item with overwrite", func(t *testing.T) {
		SaveItem(&Item{Name: "overwrite", Age: 1})
		os.Setenv("CRAWLAB_DEDUP_METHOD", "overwrite")
		os.Setenv("CRAWLAB_DEDUP_FIELD", "Name")
		os.Setenv("CRAWLAB_IS_DEDUP", "true")
		SaveItem(&Item{Name: "overwrite", Age: 2})
	})
}
