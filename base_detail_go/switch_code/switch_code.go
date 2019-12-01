package main

import (
	"fmt"
)

func main() {
	// 表达式switch,在表达式switch中,case包含与switch表达式的值进行比较的表达式
	// 类型switch, 在类型switch中,case包含与特殊注释的switch表达式的类型进行比较的类型
	// 1. 表达式sqitch
	grade := "B"
	var score int
	fmt.Scanln(&score)
	fmt.Printf("%d", score)
	switch {
	case score >= 90:
		grade = "A"
	case score >= 80:
		grade = "B"
	case score >= 70:
		grade = "C"
	case score >= 60:
		grade = "D"
	default:
		grade = "E"
	}
	fmt.Printf("成绩为%s\n", grade)

	switch {
	case grade == "A":
		fmt.Printf("成绩优秀! \n")
		//fallthrough // 这个关键词可以继续执行后面case中的语句
	case grade == "B":
		fmt.Printf("再接再厉! \n")
		//fallthrough
	//case grade == "C":
	//	fmt.Printf("再接再厉! \n")
	//case grade == "D":
	//	fmt.Printf("再接再厉! \n")
	// 多表达式写法
	case grade == "C", grade == "D":
		fmt.Println("再接再厉!")
	}

	// 类型switch
	// 类型switch的初始化子语句中需要判断的类型变量必须是具有接口类型的变量(如果是固定类型的变量就没有判断的意义了)
	var x interface{}
	switch x.(type) {
	case nil:
		fmt.Println("这里是nil")
	case int:
		fmt.Println("这里是int")
	case float64:
		fmt.Println("这里是float64")
	case bool:
		fmt.Println("这里是bool")
	case string:
		fmt.Println("这里是string")
	default:
		fmt.Println("未知类型")
	}
}
