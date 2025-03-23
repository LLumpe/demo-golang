package dal

import (
	"fmt"
	"math"

	"demo/model"
	"demo/util/mysql"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// GetTestDataPage 分页查询.
func GetTestDataPage(c *gin.Context, current int, size int) (*model.TestDataPage, error) {
	// 计算偏移量.
	offset := (current - 1) * size

	// 查询参数，为空则不查询.
	name := c.Query("name")

	// 查询总记录数.
	var total int64
	db := mysql.GetDB()

	query := db.Model(model.TestData{})
	if name != "" {
		query = query.Where("name LIKE ?", fmt.Sprintf("%%%s%%", name))
	}
	err := query.Count(&total).Error
	if err != nil {
		logrus.Errorf("[DB] count test_data failed, err：%v", err)
		return nil, err
	}

	// 计算总页数.
	totalPage := int(math.Ceil(float64(total) / float64(size)))

	// 查询当前页数据.
	var testDataList []*model.TestData
	err = query.Offset(offset).Limit(size).Find(&testDataList).Error
	if err != nil {
		logrus.Errorf("[DB] get test_data list failed, err：%v", err)
		return nil, err
	}

	return &model.TestDataPage{
		Total:        total,
		TotalPage:    totalPage,
		Size:         size,
		Current:      current,
		TestDataList: testDataList,
	}, nil
}
