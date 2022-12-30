package main

import (
	"fmt"
	"strconv"
)

func main() {
	const (
		a = iota
		b
		c
		d
	)
	cate_id, _ := strconv.Atoi("123")
    fmt.Println(cate_id)

}
