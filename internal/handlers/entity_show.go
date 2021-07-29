package handlers

import (
	"net/http"
	"strconv"
	"strings"
	"test_task/internal/dao"
)

func HandlerShowAll(w http.ResponseWriter, r *http.Request) {
	entity1, err := dao.Entity1DB.GetAll()
	if err != nil {
		MakeResponse(w, nil, err)
	}
	entity2, err := dao.Entity2DB.GetAll()
	if err != nil {
		MakeResponse(w, nil, err)
	}

	var entities []interface{}

	for _, el := range entity1{
		entities = append(entities, el)
	}
	for _, el := range entity2{
		entities = append(entities, el)
	}

	MakeResponse(w, entities, err)
}



func HandlerShowFirst(w http.ResponseWriter, r *http.Request) {
	id,_ := GetLasParam(r.URL.Path)
	entity1, err := dao.Entity1DB.GetByID(id)
	MakeResponse(w, entity1, err)
}

func HandlerShowSecond(w http.ResponseWriter, r *http.Request) {
	id, _ := GetLasParam(r.URL.Path)
	entity2, err := dao.Entity2DB.GetByID(id)
	MakeResponse(w, entity2, err)
}

func GetLasParam(path string) (int, error){
	paths := strings.Split(path, "/")
	i, err := strconv.ParseInt(paths[len(paths)-1], 10, 64)
	if err != nil {
		return 0, err
	}
	return int(i), nil
}