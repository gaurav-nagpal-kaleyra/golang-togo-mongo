package usecase

import (
	"fmt"
	"net/http"
)

func AllPathHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Invalid path")

}
