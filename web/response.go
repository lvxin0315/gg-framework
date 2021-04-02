package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	OKMsg    = "ok"
	ErrorMsg = "error"
)

const (
	ErrorCode = 400
	OKCode    = 200
)

type Output struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

/**
 * @Author lvxin0315@163.com
 * @Description 成功基础返回值格式
 * @Date 5:55 下午 2020/12/17
 * @Param data interface
 * @Param msg ...string
 **/
func SuccessJson(ctx *gin.Context, data interface{}, msg ...string) {
	message := ""
	if len(msg) == 0 {
		message = OKMsg
	} else {
		message = strings.Join(msg, " ")
	}
	ctx.JSON(http.StatusOK, Output{
		Code: OKCode,
		Msg:  message,
		Data: data,
	})
}

/**
 * @Author lvxin0315@163.com
 * @Description 失败基础返回值格式
 * @Date 5:55 下午 2020/12/17
 * @Param data interface
 * @Param msg ...string
 **/
func FailJson(ctx *gin.Context, data interface{}, msg ...string) {
	message := ""
	if len(msg) == 0 {
		message = ErrorMsg
	} else {
		message = strings.Join(msg, " ")
	}
	ctx.JSON(http.StatusOK, Output{
		Code: ErrorCode,
		Msg:  message,
		Data: data,
	})
}
