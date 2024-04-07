package usecase

import (
	"fmt"
	"net/http"
)

func UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("UpdateTaskHandler")
}
