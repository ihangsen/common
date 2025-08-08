package strs

import (
	"fmt"
	"regexp"
	"strings"
)

var printfReg = regexp.MustCompile(`\{(\w+)}`)

func Join(ss ...string) string {
	return strings.Join(ss, "")
}

func Sprintf(format string, params map[string]any) string {
	return printfReg.ReplaceAllStringFunc(format, func(str string) string {
		key := printfReg.FindStringSubmatch(str)[1]
		if val, ok := params[key]; ok {
			return fmt.Sprint(val)
		}
		return str
	})
}
