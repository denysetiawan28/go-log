package util

import (
	"github.com/denysetiawan28/go-log/src/constanta"
	"github.com/denysetiawan28/go-log/src/payload/response"
	"net/http"
)

func CreateBadRequestResponse(dt constanta.ErrMsg, data interface{}) (code int, resp response.GlobalResponse) {
	code = http.StatusBadRequest
	resp = response.GlobalResponse{
		Status:  dt.Code,
		Message: dt.Msg,
		Data:    data,
	}
	return
}

func CreateFailedResponse(dt constanta.ErrMsg, data interface{}) (code int, resp response.GlobalResponse) {
	code = http.StatusInternalServerError
	resp = response.GlobalResponse{
		Status:  dt.Code,
		Message: dt.Msg,
		Data:    data,
	}
	return
}

func CreateSuccessResponse(data interface{}) (code int, resp response.GlobalResponse) {
	code = http.StatusOK
	resp = response.GlobalResponse{
		Status:  constanta.Success.Code,
		Message: constanta.Success.Msg,
		Data:    data,
	}
	return
}
