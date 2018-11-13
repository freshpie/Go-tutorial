package main

import (
	"fmt"
)

func main() {
	aMap := map[string]interface{}{
		"GOOG": "Google Inc",
		"MSFT": 32234234,
		"FB":   4949.34590,
	}
	bMap := map[string]string{
		"GOOG": "Google Inc",
		"MSFT": "awef",
		"FB":   "432",
	}

	var cMap map[string]interface{}

	dMap := make(map[string]interface{})
	dMap["key1"] = "value1"

	fmt.Println(aMap)
	fmt.Println(bMap)
	fmt.Println(cMap)
	fmt.Println(dMap)
	fmt.Println(dMap["key1"])
	fmt.Println(dMap["key2"])
	fmt.Println("===============================")

	fmt.Printf("aMap['GOOG'] : %T\n", aMap["GOOG"])
	fmt.Printf("aMap['MSFT'] : %T\n", aMap["MSFT"])
	fmt.Printf("aMap['FB'] : %T\n", aMap["FB"])
	fmt.Printf("%T\n", dMap)
	fmt.Printf("%T\n", dMap["key1"])
	fmt.Printf("%T\n", dMap["key2"])

	if _, ok := aMap["MSFT"].(string); ok {
		fmt.Println("aMap['MSFT'] is string")
	} else if _, ok := aMap["MSFT"].(int); ok {
		fmt.Println("aMap['MSFT'] is int")
	} else {
		fmt.Println("aMap['MSFT'] is not string")
	}

	switch aMap["GOOG"].(type) {
	case int:
		fmt.Println("int")
	case float64:
		fmt.Println("float64")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("?")
	}

	fmt.Println("===============================")
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64) // panic
	fmt.Println(f)
	fmt.Println("===============================")

}
