package main

import "fmt"

func print(val interface{}) {
	switch v := val.(type) {
	case string:
		fmt.Println("string")
		fmt.Println(v)
	case int:
		fmt.Println("int")
		fmt.Println(fmt.Sprint(v))
	case int64:
		fmt.Println("int64")
		fmt.Println(fmt.Sprint(v))
	}
}

func main() {
	s := "Hello, World"
	result := add(9223372036854775806, 1)
	print(s)
	print(result)
	mymap := createMap()
	print(mymap["hello"])
}

func add(num1 int64, num2 int64) int64 {
	return num1 + num2
}

func createMap() map[string]string {
	mymap := map[string]string{"hello": "10", "name": "23"}
	return mymap
}
