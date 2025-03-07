package leetcodetop100

import (
	"strconv"
	"strings"
	"testing"
)

func DecodeString(str string) string {
	var stack []string
	count := 0
	var builder strings.Builder
	for index, value := range str {
		println("in loop ", index)
		if value >= '0' && value <= '9' {
			count = (count*10 + int(value-'0'))
			println("in loop ", index, " the count is ", count)
		} else if (value >= 'a' && value <= 'z') || (value >= 'A' && value <= 'Z') {
			builder.WriteRune(value)
		} else if value == '[' {
			stack = append(stack, builder.String())
			stack = append(stack, strconv.Itoa(count))
			builder.Reset()
			count = 0
		} else {
			countStr := stack[len(stack)-1]
			preStr := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			tempCount, _ := strconv.Atoi(countStr)
			var curStr strings.Builder
			for i := 0; i < tempCount; i++ {
				curStr.WriteString(builder.String())
			}

			preStr += curStr.String()
			builder.Reset()
			builder.WriteString(preStr)
		}
	}
	return builder.String()
}

func TestDecodeString(t *testing.T) {
	str := "3[a]2[bc]"
	str = "100[leetcode]"
	var formatStr = DecodeString(str)

	println(formatStr, " the count is ", len(strings.Split(formatStr, "leetcode")))
}
