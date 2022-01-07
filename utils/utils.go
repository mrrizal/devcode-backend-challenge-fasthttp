package utils

import (
	"encoding/json"

	"github.com/valyala/fasthttp"
)

func GenerateErrorMessage(status, message string, data interface{}) map[string]interface{} {
	errorMessage := make(map[string]interface{})
	errorMessage["status"] = status
	errorMessage["message"] = message
	errorMessage["data"] = data
	return errorMessage
}

func ResponseHandler(ctx *fasthttp.RequestCtx, statusCode int, message map[string]interface{}) {
	ctx.Response.Header.Set("Content-Type", "application/json")
	ctx.Response.SetStatusCode(statusCode)
	if err := json.NewEncoder(ctx).Encode(message); err != nil {
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
	}
}
