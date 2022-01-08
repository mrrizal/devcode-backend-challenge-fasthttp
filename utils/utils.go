package utils

import (
	"encoding/json"

	"github.com/valyala/fasthttp"
)

type MessageStruct struct {
	Status  interface{} `json:"status"`
	Message interface{} `json:"message"`
	Data    interface{} `json:"data"`
}

func GenerateResponse(status, message string, data interface{}) MessageStruct {
	response := MessageStruct{
		Status:  status,
		Message: message,
		Data:    data,
	}
	return response
}

func ResponseHandler(ctx *fasthttp.RequestCtx, statusCode int, message MessageStruct) {
	ctx.Response.Header.Set("Content-Type", "application/json")
	ctx.Response.SetStatusCode(statusCode)
	if err := json.NewEncoder(ctx).Encode(message); err != nil {
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
	}
}
