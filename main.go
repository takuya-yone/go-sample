package main

import (
	"fmt"
	"github.com/takuya-yone/sample-go/subpkg"
)

func main() {
	fmt.Println(subpkg.Hello())
	fmt.Println(subpkg.Golang())
	fmt.Println(Goodbye())
}
