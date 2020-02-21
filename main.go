package main

import "fmt"

func main() {
	ch := make(chan string, 1)
	process(ch)
	fmt.Println(<-ch)
}


//Channel icerisindeki string degisir...
func process (ch chan string) error{
	ch<-"Sinan"
	return nil
}