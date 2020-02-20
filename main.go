package main

import (
	"fmt"
	"log"
	"strings"
)

func main() {
	fmt.Println("B...")
	fmt.Println(sum(3, 8))

	hiName, err := hi("hafsa")
	if err != nil {
		log.Println("hi err:", err.Error())
	}else {
		fmt.Println(hiName)
	}
}

func sum(n1, n2 int) int {
	return n1 + n2
}

func hi(name string) (string, error) {
	name = strings.ToTitle("Selam " + name)
	return name, nil
}

func process() {

}