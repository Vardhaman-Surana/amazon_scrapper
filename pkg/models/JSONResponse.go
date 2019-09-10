package models

type JSONResponse struct{
	Msg string	`json:"msg"`
	Err error	`json:"err"`
}

func NewJsonResponse(msg string,err error) JSONResponse{
	return JSONResponse{
		Msg: msg,
		Err: err,
	}
}
