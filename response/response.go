package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int                    `json:"code,omitempty"`
	Status  string                 `json:"status,omitempty"`
	Message string                 `json:"message,omitempty"`
	Data    map[string]interface{} `json:"data,omitempty"`
	Error   error                  `json:"error,omitempty"`
}

func (resp *Response) SendError(ctx *gin.Context, code int, err error) {
	resp.Code = code
	resp.Status = "fail"
	resp.Error = err
	resp.Message = err.Error()
	ctx.AbortWithStatusJSON(resp.Code, resp)
}
func (resp *Response) SendData(ctx *gin.Context, data map[string]interface{}) {
	resp.Code = http.StatusOK
	resp.Status = "Success"
	resp.Data = data
	ctx.JSON(resp.Code, resp)
}
