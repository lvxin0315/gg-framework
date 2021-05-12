package tools

// 字符串相关操作
import (
	"regexp"
	"strings"
)

// StrCut
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

// FilterChinese
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

// CamelString camel string, xx_yy to XxYy
func CamelString(s string) string {
	data := make([]byte, 0, len(s))
	j := false
	k := false
	num := len(s) - 1
	for i := 0; i <= num; i++ {
		d := s[i]
		if k == false && d >= 'A' && d <= 'Z' {
			k = true
		}
		if d >= 'a' && d <= 'z' && (j || k == false) {
			d = d - 32
			j = false
			k = true
		}
		if k && d == '_' && num > i && s[i+1] >= 'a' && s[i+1] <= 'z' {
			j = true
			continue
		}
		data = append(data, d)
	}
	return string(data[:])
}

// Capitalize 如果是小写字母, 则变换为大写字母
func Capitalize(str string) string {
	var upperStr string
	vv := []rune(str) // 后文有介绍
	for i := 0; i < len(vv); i++ {
		if i == 0 {
			if vv[i] >= 97 && vv[i] <= 122 { // 后文有介绍
				vv[i] -= 32 // string的码表相差32位
				upperStr += string(vv[i])
			} else {
				return str
			}
		} else {
			upperStr += string(vv[i])
		}
	}
	return upperStr

}
