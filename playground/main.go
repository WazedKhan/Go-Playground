package main

import (
	"fmt"
	"strings"
)

func main() {
	// convert slice of ints into comma separated string
	nums := []int64{1,2,3,4,5,6}
	result := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(nums)), ","), "[]")
	fmt.Println(result)
}