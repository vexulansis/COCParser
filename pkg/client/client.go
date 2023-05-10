package client

import (
	"database/sql"
	"fmt"

	"github.com/go-resty/resty/v2"
	. "github.com/vexulansis/COCParser/pkg/task"
)

type Client interface {
	Exec(task Task)
}
type TestClient struct {
	DB   *sql.DB
	HTTP *resty.Client
}

func (c *TestClient) Exec(task Task) {
	fmt.Printf("%v %v IN EXECUTION\n", task.Type, task.Data)
}
