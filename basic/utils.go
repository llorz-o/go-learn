package basic

import "strings"

func Divide(tit string, f func()) {
	println()
	println("*****************", tit)
	f()
}

func Split(s, sep string) (result []string) {
	i := strings.Index(s, sep)

	for i > -1 {
		result = append(result, s[:i])
		s = s[i+1:]
		i = strings.Index(s, sep)
	}
	result = append(result, s)
	return
}