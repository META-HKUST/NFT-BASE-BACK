package base

import "encoding/json"

type Response struct {
	Code int         `json:"code"` // 错误码
	Msg  string      `json:"msg"`  // 错误描述
	Data interface{} `json:"data"` // 返回数据
}

func (res *Response) SetData(data interface{}) Response {
	return Response{
		Code: res.Code,
		Msg:  res.Msg,
		Data: data,
	}
}

func (res *Response) SetCode(Err ErrCode) Response {
	return Response{
		Code: int(Err),
		Msg:  Err.String(),
		Data: res.Data,
	}
}

func (res *Response) ToString() string {
	err := &struct {
		Code int         `json:"code"`
		Msg  string      `json:"msg"`
		Data interface{} `json:"data"`
	}{
		Code: res.Code,
		Msg:  res.Msg,
		Data: res.Data,
	}
	raw, _ := json.Marshal(err)
	return string(raw)
}

type PageResponse struct {
	Code       int         `json:"code"`  // 错误码
	Msg        string      `json:"msg"`   // 错误描述
	TotalCount int         `json:"total"` //总数
	Data       interface{} `json:"data"`  // 返回数据
}

func (res *PageResponse) SetData(data interface{}) PageResponse {
	return PageResponse{
		Code:       res.Code,
		Msg:        res.Msg,
		Data:       data,
		TotalCount: res.TotalCount,
	}
}

func (res *PageResponse) SetCode(Err ErrCode) PageResponse {
	return PageResponse{
		Code:       int(Err),
		Msg:        Err.String(),
		Data:       res.Data,
		TotalCount: res.TotalCount,
	}
}

func (res *PageResponse) SetCount(count int) PageResponse {
	return PageResponse{
		Code:       res.Code,
		Msg:        res.Msg,
		Data:       res.Data,
		TotalCount: count,
	}
}

// ToString 返回 JSON 格式的错误详情
func (res *PageResponse) ToString() string {
	err := &struct {
		Code       int         `json:"code"`
		Msg        string      `json:"msg"`
		Data       interface{} `json:"data"`
		Totalcount int         `json:"totalcount"`
	}{
		Code:       res.Code,
		Msg:        res.Msg,
		Data:       res.Data,
		Totalcount: res.TotalCount,
	}
	raw, _ := json.Marshal(err)
	return string(raw)
}
