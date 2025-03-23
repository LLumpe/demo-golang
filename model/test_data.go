package model

// TestData 测试数据结构体.
type TestData struct {
	ID   int64  `json:"ID"`
	Name string `json:"Name"`
	Num  int64  `json:"Num"`
}

func (r *TestData) TableName() string {
	return "test_data"
}

// TestDataPage 分页查询结构.
type TestDataPage struct {
	Total        int64       `json:"Total`
	TotalPage    int         `json:"TotalPage"`
	Current      int         `json:"Current"`
	Size         int         `json:"Size"`
	TestDataList []*TestData `json:"TestDataList"`
}
