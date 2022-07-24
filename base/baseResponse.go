package base

import "encoding/json"

type Response struct {
	Code int         `json:"code"`  // 错误码
	Msg  string      `json:"msg" `  // 错误描述
	Data interface{} `json:"data" ` // 返回数据
}

type TokenUrlResponse struct {
	Name  string	`json:"name"`
	Image string	`json:"image"`
}


func (res *Response) SetData(data interface{}) Response {
	res.Data = data
	return Response{
		Code: res.Code,
		Msg:  res.Msg,
		Data: data,
	}
}

func (res *Response) SetCode(Err ErrCode) Response {
	res.Code = int(Err)
	res.Msg = Err.String()
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
