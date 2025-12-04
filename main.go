package main

import (
	"fmt"
)

func main() {
	res := ((25-1)%99+99)%99 + 1
	fmt.Println(res)
}
