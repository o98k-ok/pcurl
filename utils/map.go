package utils

import "strings"

func Convert(src map[string]string, innerDelim, outerDelim string) string {
	arr := make([]string, 0, len(src))
	for k, v := range src {
		arr = append(arr, k+innerDelim+v)
	}
	return strings.Join(arr, outerDelim)
}
