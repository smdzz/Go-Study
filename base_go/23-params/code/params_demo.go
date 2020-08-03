package main

import (
	"flag"
	"fmt"
)

var (
	user string
	pwd  string
	host string
	port int
)

func init() {
	flag.StringVar(&user, "u", "root", "用户名")
	flag.StringVar(&pwd, "p", "root", "密码")
	flag.StringVar(&host, "h", "127.0.0.1", "host")
	flag.IntVar(&port, "P", 3306, "端口号")
}

func main() {
	flag.Parse()
	fmt.Printf("host: %v, pwd: %v, host: %v, port: %v\n", user, pwd, host, port)
}
