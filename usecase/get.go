package usecase

import (
	"fmt"
	"net/http"
)

func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetTaskHandler")
}
