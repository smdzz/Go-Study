package main

import (
	"fmt"
	"reflect"
)

//func main() {
//	var x float64 = 5.6
//	fmt.Println("type:", reflect.TypeOf(x))
//	v := reflect.ValueOf(x)
//	v.CanSet()
//	fmt.Println("value: ", v)
//	fmt.Println("type: ", v.Type())
//	fmt.Println("kind: ", v.Kind())
//	fmt.Println(v.Kind() == reflect.Float64)
//	fmt.Println("value: ", v.Float())
//	fmt.Println(v.Interface())
//	y := v.Interface().(float64)
//	fmt.Println(y)
//}


//func main() {
//	var x float64 = 5.6
//	fmt.Println("type:", reflect.TypeOf(x))
//	v := reflect.ValueOf(x)
//	v.CanSet()
//	v = reflect.ValueOf(&x) // 这里要取x的地址
//	// CanSet可以用来测试value值是否可设置
//	fmt.Println("v的类型: ", v.CanSet())
//	fmt.Println("v: 是否可设置", v.CanSet())
//	v = v.Elem()
//	fmt.Println("v的Elem是: ",v)
//	fmt.Println("v: 是否可设置", v.CanSet())
//	v.SetFloat(8.996)
//	fmt.Println(v.Interface())
//	fmt.Println(v)
//}


//type T struct {
//	s1, s2, s3 string
//}
//
//func (n T) String() string {
//	return n.s1 + "-" + n.s2 + "-" + n.s3
//}

//func main() {
//	var secret interface{} = T{"Google", "Go", "Python"}
//	value := reflect.ValueOf(secret)
//	typ := reflect.TypeOf(secret)
//	fmt.Println("type: ", typ)
//	knd := value.Kind()
//	fmt.Println("kind: ", knd)
//
//	for i := 0; i < value.NumField(); i++ {
//		fmt.Printf("Field %d: %v\n", i, value.Field(i))
//		//value.Field(i).SetString("我是")
//		// 返回错误,结构中只有被导出的字段才能够被设置
//	}
//
//	results := value.Method(0).Call(nil)
//	fmt.Println(results)
//}

type T struct {
	A int
	B string
}

func main() {
	t := T{23, "hello"}
	s := reflect.ValueOf(&t).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i, typeOfT.Field(i).Name, f.Type(), f.Interface())
	}

	s.Field(0).SetInt(77)
	s.Field(1).SetString("hello world")
	fmt.Println("现在的t值是--->", t)

}