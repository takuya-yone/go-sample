package subpkg

import (
	"strconv"
)

func Hello() (str string) {
	myMap := map[string]int{"a": 1, "b": 2}
	// fmt.Print(myMap["a"])
	// return quote.HelloV3()
	return strconv.Itoa(myMap["a"])
}
