package main

import (
	"fmt"
	"sort"
)

// 声明一个Hero结构体
type Hero struct {
	Name string
	Age int
}

type HeroSlice []Hero

func (hs HeroSlice) Len() int {
	return len(hs)
}

func (hs HeroSlice) Less(i, j int) bool {
	return hs[i].Age < hs[j].Age
}

func (hs HeroSlice) Swap(i, j int) {
	hs[i], hs[j] = hs[j], hs[i]
}

func main() {
	intSlice := []int{1, 6, 3, 9, 4}
	sort.Ints(intSlice)
	fmt.Println(intSlice)

	// 只要HeroSlice实现了Len(),Less(),Swap()这几个方法,就可以使用sort.Sort排序
	heroes := HeroSlice{{Age: 20, Name: "a"}, {Age: 15, Name: "b"}, {Age: 30, Name: "c"}, {Age: 6, Name: "d"}}
	fmt.Println(heroes)
	sort.Sort(heroes)
	fmt.Println(heroes)
}