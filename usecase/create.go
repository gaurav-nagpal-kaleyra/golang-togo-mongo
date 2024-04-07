package usecase

import (
	"fmt"
	"net/http"
)

func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CreateTaskHandler")

	// accept array of tasks and insert it in the mongodb collection
}
