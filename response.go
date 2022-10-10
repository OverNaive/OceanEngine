package oceanengine

import (
	"fmt"
)

type BaseResponse struct {
	// Code 返回码
	Code int `json:"code,omitempty"`
	// Message 返回信息
	Message string `json:"message,omitempty"`
	// RequestID 请求的日志id，唯一标识一个请求
	RequestID string `json:"request_id,omitempty"`
}

func (r *BaseResponse) IsError() bool {
	return r.Code != 0
}

func (r *BaseResponse) ErrorMessage() string {
	return fmt.Sprintf("%d:%s", r.Code, r.Message)
}

type DataResponse struct {
	BaseResponse
	Data map[string]interface{} `json:"data,omitempty"`
}

type ListResponse struct {
	BaseResponse
	Data *struct {
		List         []map[string]interface{} `json:"list,omitempty"`
		CommentsList []map[string]interface{} `json:"comments_list,omitempty"`
		PageInfo     *struct {
			Page        int   `json:"page,omitempty"`
			PageSize    int   `json:"page_size,omitempty"`
			TotalNumber int64 `json:"total_number,omitempty"`
			TotalPage   int   `json:"total_page,omitempty"`
		} `json:"page_info,omitempty"`
	} `json:"data,omitempty"`
}
