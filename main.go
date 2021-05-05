package main

import (
	"datastruct/set"
	"fmt"
)

func main() {
	l := make([]int, 5)
	s := set.NewSet()
	s.Insert(l)
	s.Insert("Abc")
	fmt.Println(s.Find("Abc"))
	fmt.Println(s.Find("bcd"))
}



