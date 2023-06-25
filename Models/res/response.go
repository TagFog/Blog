package res

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

//实例化
func Result(code int, data interface{}, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}

const (
	Success = 0
	Error   = 1
)

func Ok(data interface{}, msg string, c *gin.Context) {
	Result(Success, data, msg, c)
}

func OkWithData(data interface{}, c *gin.Context) {
	Result(Success, data, "成功", c)
}

func OkWithMessage(data interface{}, msg string, c *gin.Context) {
	Result(Success, map[string]interface{}{}, msg, c)
}

func Fail(data interface{}, c *gin.Context) {
	Result(Success, data, "成功", c)
}

func FileWithMessage(data interface{}, msg string, c *gin.Context) {
	Result(Success, map[string]interface{}{}, msg, c)
}
func FailWithCode(code ErrorCode, c *gin.Context) {
	msg, ok := ErrorMap[code]
	if ok {
		Result(int(code), map[string]interface{}{}, msg, c)
		return
	}
	Result(Error, map[string]interface{}{}, "未知错误", c)
}
