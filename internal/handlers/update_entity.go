package handlers

import (
	"encoding/json"
	"net/http"
	"test_task/internal/dao"
	"test_task/internal/entity"
	"test_task/internal/logger"
)

func HandlerUpdateFirst(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	entity1 := &entity.Entity1{}

	err := decoder.Decode(entity1)
	if err != nil {
		MakeResponse(w, nil, err)
		logger.Logger.Error(err)
		return
	}

	entity1.Id, err = GetLasParam(r.URL.Path)
	if err != nil {
		MakeResponse(w, nil, err)
		logger.Logger.Error(err)
		return
	}

	err = dao.Entity1DB.Update(entity1)
	if err != nil {
		MakeResponse(w, nil, err)
		logger.Logger.Error(err)
		return
	}

	MakeResponse(w, "update", nil)
}

func HandlerUpdateSecond(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	entity2 := &entity.Entity2{}

	err := decoder.Decode(entity2)
	if err != nil {
		MakeResponse(w, nil, err)
		logger.Logger.Error(err)
		return
	}

	entity2.Id, err = GetLasParam(r.URL.Path)
	if err != nil {
		MakeResponse(w, nil, err)
		logger.Logger.Error(err)
		return
	}

	err = dao.Entity2DB.Update(entity2)
	if err != nil {
		MakeResponse(w, nil, err)
		logger.Logger.Error(err)
		return
	}

	MakeResponse(w, "update", nil)
}
