package utils

type PageInfo struct {
	Page     int `json:"page" validate:"required" form:"page"`
	PageSize int `json:"pageSize" form:"pageSize"`
}

func defaultPageSize() int {
	return 15
}

// set default value
func (pageInfo *PageInfo) ParsePageInfo() {
	if pageInfo.PageSize == 0 {
		pageInfo.PageSize = defaultPageSize()
	}
}
