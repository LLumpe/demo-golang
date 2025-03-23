package main

import (
	"demo/service"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// routeInit 路由初始化
func routeInit() {
	r := gin.Default()
	r.GET("/GetTestDataPage", service.GetTestDataPage)
	err := r.Run(":8000")
	if err != nil {
		logrus.Errorf("listening port failed, err: %v", err)
		panic(err)
	}
}
