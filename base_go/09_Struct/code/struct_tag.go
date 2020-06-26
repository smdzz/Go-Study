package main

import (
	"fmt"
	"reflect"
	"time"
)

type TagType struct {
	Name string "商品名称"
	Price int  "商品价格"
}
func main() {
	tt := TagType{"iphon", 5200, }
	for i := 0; i < 2; i++ {
		refTag(tt, i)
	}
}

func refTag(tt TagType, i int) {
	time.Sleep(1 * time.Second)
	ttType := reflect.TypeOf(tt)
	ixField := ttType.Field(i)
	fmt.Printf("%v\n", ixField.Tag)
}
