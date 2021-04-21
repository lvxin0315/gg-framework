package tools

import (
	"fmt"
	"strconv"
)

/**
 * @Author lvxin0315@163.com
 * @Description float 保留两位小数
 * @Date 10:21 上午 2021/4/21
 * @Param float
 * @return float
 **/
func Decimal64E2(value float64) float64 {
	newVale, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	return newVale
}

func Decimal32E2(value float32) float32 {
	newVale, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", value), 32)
	return float32(newVale)
}
