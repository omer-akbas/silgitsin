package main

import "fmt"

func main() {
	ch := make(chan string, 1)
	process(ch)
	fmt.Println(<-ch)

	total := sum(8, 3)
	fmt.Println("total: ", total)
}


//Channel icerisindeki string degisir...
func process (ch chan string) error{
	ch<-"Sinan"
	return nil
}

func sum(n1, n2 int) int {
	return n1 + n2
}