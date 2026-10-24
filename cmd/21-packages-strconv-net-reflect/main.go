package main

import (
	"fmt"
	"net"
	"reflect"
	"strconv"
)

func main() {
	n, _ := strconv.Atoi("123")
	ip := net.ParseIP("127.0.0.1")
	fmt.Println(n, ip, reflect.TypeOf(n))
}
