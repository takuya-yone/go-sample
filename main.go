package main

import (
	// "go-sample/controllers"

	"github.com/takuya-yone/go-sample/controllers"
)

func main() {
	// fmt.Println(subpkg.Hello())
	// fmt.Println(subpkg.Golang())
	// fmt.Println(Goodbye())
	controllers.StartWebServer()
}
