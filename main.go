package main

import (
	"fmt"

	"github.com/takuya-yone/go-sample/subpkg"
)

func main() {
	fmt.Println(subpkg.Hello())
	fmt.Println(subpkg.Golang())
	fmt.Println(Goodbye())
}
