package usecase

import (
	"fmt"
	"net/http"
)

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DeleteTaskHandler")
}
