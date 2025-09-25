package common

type successRes struct {
	Data   interface{} `json:"data"`
	Paging interface{} `json:"paging,omitempty"`
	Filter interface{} `json:"filter,omitempty"`
}

func NewSuccessPagingRes(data interface{}, paging interface{}, filter interface{}) *successRes {
	return &successRes{
		Data:   data,
		Paging: paging,
		Filter: filter,
	}
}
func NewSuccessRes(data interface{}) *successRes {
	return NewSuccessPagingRes(data, nil, nil)
}
