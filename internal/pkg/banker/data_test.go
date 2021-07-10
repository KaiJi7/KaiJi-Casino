package banker

import (
	"testing"
)

func Test_banker_GetGambles(t *testing.T) {
	aList := []string{"h", "e", "l", "l", "o"}

	var tList = make([]string, len(aList))
	copy(tList, aList)
	for i, w := range aList {
		if w == "l" {
			tList = remove(tList, i)
		}
	}
	t.Log(tList)
}

func remove(s []string, i int) []string {
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	return s[:len(s)-1]
}
