package main

import (
	"fmt"
	"time"
)

func main(){
 	//var wg sync.WaitGroup

 	c := make(chan string)


 	//wg.Add(1)
	go DoMyWork("Naveen", c)

 	for msg := range c{
 		fmt.Println(msg)
	}


 	//wg.Wait()
 	//time.Sleep(time.Second * 3)
}

func DoMyWork(name string, c chan string){
	for i:=1; i<=4 ; i++ {
		c<-fmt.Sprintf("Doing work of %v executed %v ", name,i)
		time.Sleep(time.Millisecond * 500)
	}
	close(c)
}
