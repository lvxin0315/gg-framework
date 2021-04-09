package tools

// 字符串相关操作
import (
	"regexp"
	"strings"
)

/**
 * @Author lvxin0315@163.com
 * @Description 截取字符串
 * @Date 2:47 下午 2021/4/2
 * @Param str
 * @Param start 第几位
 * @Param cutLen 截取长度
 * @return
 **/
func StrCut(str string, start, cutLen int) string {
	rs := []rune(str)
	if start >= len(str) {
		return ""
	}
	if cutLen >= len(rs) {
		cutLen = len(rs)
	}
	return string(rs[start:cutLen])
}

/**
 * @Author lvxin0315@163.com
 * @Description 过滤掉中文
 * @Date 4:02 下午 2021/4/8
 * @Param
 * @return
 **/
func FilterChinese(src string) string {
	var valid = regexp.MustCompile("[0-9]")
	rnArr := valid.FindAllStringSubmatch(src, -1)
	var rArr []string
	for _, v := range rnArr {
		rArr = append(rArr, strings.Join(v, ""))
	}
	return strings.Join(rArr, "")
}
