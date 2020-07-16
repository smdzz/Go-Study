package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	for x, y := range []rune("11ddfkj接发空格") {
		fmt.Println(x, string(y))
	}
	str := strconv.FormatInt(123, 2)
	fmt.Println(str)
	bool_ := strings.Contains("woe们尽快发动机", "woe们")
	fmt.Println(bool_)
	num := strings.Count("aaaasss", "a")
	fmt.Println(num)
	fmt.Println(strings.EqualFold("aa", "AA"))
	fmt.Println(strings.Index("name", "9"))
	fmt.Println(strings.LastIndex("name as", "a"))
	fmt.Println(strings.Replace("e name", "e", " 替换 ", -1))
	fmt.Println(strings.Split("name,time,hello", ","))
	fmt.Println(strings.ToUpper("name"), strings.ToLower("RUjj"))
	fmt.Println(strings.TrimSpace("  hdjksfah     "))
	fmt.Println(strings.Trim("a hdjksfah     a  ", "a "))
	fmt.Println(strings.TrimSuffix("hdjksfah.jpg ", " "))
	fmt.Println(strings.HasPrefix("https://", "https"))
	fmt.Println(strings.HasSuffix("https://", "//"))

}
