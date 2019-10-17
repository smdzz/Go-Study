package main

import "fmt"

type Books struct {
	title string
	author string
	price int
	book_id int
}

func main() {
	var book1 Books
	book1 = Books{"小人书", "哈利波特", 23, 34567}
	fmt.Println((Books{title: "小黄人", author: "大连湾", book_id: 3456, price: 34}))
	fmt.Println(Books{title: "钢铁是怎样炼成的"})
	fmt.Println(book1.title, book1.author, book1.price)
	book1.price = 78
	// 值传递
	fmt.Println(book1.price)
	printBook(book1)
	fmt.Println(book1)
	// 指针传递
	fmt.Println(book1)
	changeBookName(&book1)
	fmt.Println(book1)
}

// 将结构体类型作为参数传递给函数,此处为值传递,不改变原来变量的结果
func printBook(book Books) {
	fmt.Println(book)
	book.title = "西游记"
	fmt.Println(book)
	fmt.Println("进入函数")
}

// 结构体指针传递
func changeBookName(book *Books) {
	fmt.Println(&book)
	book.author = "金角大王"
	fmt.Println(&book)
}

