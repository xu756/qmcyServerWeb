package result

import "github.com/xu756/qmcy/common/tool"

type ResponseSuccessBean struct {
	Success   bool        `json:"success"`
	TimeStamp int64       `json:"timeStamp"`
	Data      interface{} `json:"data"`
}
type NullJson struct{}

func Success(data interface{}) *ResponseSuccessBean {
	return &ResponseSuccessBean{true, tool.TimeNowInTimeZoneUnix(), data}
}

type ResponseErrorBean struct {
	Success      bool   `json:"success"`
	TimeStamp    int64  `json:"timeStamp"`
	ErrorCode    uint32 `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
}

func Error(errCode uint32, errMsg string) *ResponseErrorBean {
	return &ResponseErrorBean{false, tool.TimeNowInTimeZoneUnix(), errCode, errMsg}
}
