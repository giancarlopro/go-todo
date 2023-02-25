package main

import (
	"encoding/json"
	"fmt"

	"github.com/giancarlopro/go-todo/pkg/models"
)

func main() {
	task := &models.Task{
		Title:  "Write presentation",
		Status: models.Pending,
	}

	data, _ := json.Marshal(task)

	fmt.Println(string(data))
}
