package stringger

import (
	"fmt"
	"testing"
)

func TestEqualStringAndInt(t *testing.T) {
	fmt.Println(IntString[0])
	fmt.Println(string(IntString[0]))
	fmt.Println(int(IntString[0]) == 0)
	fmt.Println(string(IntString[0]) == "0")
}
