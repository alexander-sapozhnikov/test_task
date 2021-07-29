package main

import (
	"fmt"
	"test_task/internal"
	"test_task/internal/dao"
	"test_task/internal/logger"
)

func main() {
	fmt.Println("Start App")
}

func init(){
	logger.InitLog()
	dao.InitDao()
	internal.InitHandlers()
}