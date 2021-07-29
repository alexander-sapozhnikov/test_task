package handlers

import (
	"encoding/json"
	"net/http"
	"test_task/internal/dao"
	"test_task/internal/entity"
	"test_task/internal/logger"
)

func HandlerAddFirst(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	entity1 := &entity.Entity1{}

	err := decoder.Decode(entity1)
	if err != nil {
		MakeResponse(w, nil, err)
		logger.Logger.Error(err)
		return
	}

	err = dao.Entity1DB.Save(entity1)
	if err != nil {
		MakeResponse(w, nil, err)
		logger.Logger.Error(err)
		return
	}

	MakeResponse(w, "ok", nil)
}

func HandlerAddSecond(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	entity2 := &entity.Entity2{}

	err := decoder.Decode(entity2)
	if err != nil {
		MakeResponse(w, nil, err)
		logger.Logger.Error(err)
		return
	}

	err = dao.Entity2DB.Save(entity2)
	if err != nil {
		MakeResponse(w, nil, err)
		logger.Logger.Error(err)
		return
	}

	MakeResponse(w, "ok", nil)
}
