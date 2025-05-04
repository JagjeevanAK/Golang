package main

import (
	"fmt"
	"sync"
)

func main(){
	fmt.Println("Welcome to channels")

	wg:=&sync.WaitGroup{}
	Ch:=make(chan int)

	wg.Add(2)
	go func (wg *sync.WaitGroup,ch chan int)  {
		defer wg.Done()
		ch<-5
	}(wg,Ch)
	go func (wg *sync.WaitGroup,ch chan int)  {
		defer wg.Done()
		fmt.Println(<-ch)
		// ch<-5
	}(wg,Ch)

	wg.Wait()
}