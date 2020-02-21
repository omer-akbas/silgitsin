package main

import (
	"fmt"
	"github.com/omer-akbas/errsend"
	"github.com/omer-akbas/gocr"
	"log"
)

func main() {
	ch := make(chan string, 1)

	err := process(ch)
	if err != nil {
		log.Println("process func err: ", err.Error())
		e := errsend.ErrSend{To: []string{"izmirlieleman@gmail.com"}, Subject: "Test", Body: err.Error(), IsTelegram:true}
		_ = e.Send()
	}
	fmt.Println(<-ch)

	total := sum(8, 3)
	fmt.Println("total: ", total)

	err = gocr.ImageToText("./temp/in/1.tif", "./temp/out/test")
	if err != nil{
		log.Println("err:", err.Error())
	}
}

//Channel icerisindeki string degisir...
func process (ch chan string) error{
	ch<-"Sinan"
	return fmt.Errorf("Cok acayp seyler oluyor...")
}

func sum(n1, n2 int) int {
	return n1 + n2
}