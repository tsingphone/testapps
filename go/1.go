package main

import (
	"fmt"
	"strconv"
)


func t(p1,p2 int,p3 string) string {
	var s = strconv.Itoa (p1 + p2)+p3;
	return s;
}


func main() {
	i := 5
	j := &i
	k := *j


	v := t(i,k,"aaa")
	fmt.Println("i=",v)
	fmt.Println("j=",j)
	fmt.Println("k=",k)
	//fmt.Println("l=",l)
}