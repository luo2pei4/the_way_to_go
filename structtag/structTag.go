package main

import (
	"fmt"
	"reflect"
)

// TagType is a test struct for struct tag
// 变量后面的字符串就是变量的标签（tag）
type TagType struct {
	field1 bool   `A:"An important answer"`
	field2 string `B:"The name of the thing"`
	field3 int    `C:"How much there are"`
}

func main() {

	tt := TagType{true, "Barak Obama", 1}
	for i := 0; i < 3; i++ {
		refTag(tt, i)
	}
}

func refTag(tt TagType, ix int) {

	fmt.Printf("address: %p\n", &tt)
	ttType := reflect.TypeOf(tt)

	fmt.Printf("Type: %s\n", ttType)

	ixField := ttType.Field(ix)

	fmt.Printf("Field: %s\n", ixField.Name)

	fmt.Printf("%v\n", ixField.Tag)
}
