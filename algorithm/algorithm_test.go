package algorithm

import (
	"fmt"
	"testing"
)

func TestRandomString(t *testing.T) {
	s := GetRandomStr()
	fmt.Println(s)
	h, err := GetFromString(s)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(h)
}
