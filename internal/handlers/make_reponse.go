package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"test_task/internal/entity"
)

func MakeResponse(w http.ResponseWriter, data interface{}, err error){
	w.Header().Set("Content-Type", "application/json")

	errStr := ""

	if err != nil {
		errStr = fmt.Sprintf("%v", err)
	}

	json.NewEncoder(w).Encode(entity.Response{
		Data:   data,
		Error: errStr,
	})
}
