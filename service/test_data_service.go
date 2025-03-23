package service

import (
	"errors"
	"net/http"
	"strconv"

	"demo/dal"
	"demo/util/open_api"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// GetTestDataPage 分页查询测试数据.
func GetTestDataPage(c *gin.Context) {

	// 参数校验.
	current, size, err := paramCheck(c)
	if err != nil {
		logrus.Errorf("param check failed, err: %v", err)
		open_api.OpenApiErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// 查询数据
	testDataPage, err := dal.GetTestDataPage(c, current, size)
	if err != nil {
		logrus.Errorf("GetTestDataPage failed, err: %v", err)
		open_api.OpenApiErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	open_api.OpenApiSuccessResponse(c, testDataPage)

}

// paramCheck 参数校验.
func paramCheck(c *gin.Context) (current int, size int, err error) {
	currentStr := c.Query("current")
	sizeStr := c.Query("size")

	// 分页默认值
	current = 1
	size = 10

	// 参数校验
	if currentStr != "" {
		current, err = strconv.Atoi(currentStr)
		if err != nil || current < 1 {
			logrus.Errorf("invalid current parameter, current = %v", current)
			err = errors.New("invalid current parameter")
			return
		}
	}

	if sizeStr != "" {
		size, err = strconv.Atoi(sizeStr)
		// 分页最大数量限制.
		if err != nil || size < 1 || size > 100 {
			logrus.Errorf("invalid size parameter, size = %v", size)
			err = errors.New("invalid size parameter")
			return
		}
	}

	return
}
