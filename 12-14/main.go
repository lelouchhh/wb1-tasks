package main

import (
	"fmt"
	"reflect"
)

func main() {
	//12
	s := []string{"cat", "cat", "dog", "tree"}
	m := make(map[string]struct{})
	for _, v := range s {
		m[v] = struct{}{}
	}
	fmt.Println(m)
	//13
	var a, b int = 1, 2
	a, b = b, a
	fmt.Println(a, b)
	//14
	var data interface{}
	data = 10
	switch value := data.(type) {
	case int:
		fmt.Printf("Тип: int, Значение: %v\n", value)
	case string:
		fmt.Printf("Тип: string, Значение: %v\n", value)
	case bool:
		fmt.Printf("Тип: bool, Значение: %v\n", value)
	case chan int:
		fmt.Printf("Тип: канал, Значение: %v\n", reflect.TypeOf(value))
	default:
		fmt.Printf("Неизвестный тип\n")
	}
}
