package main

import (
	"fmt"
	"net/http"
	"sync"
)

var statusCode []int
var mu sync.Mutex
var wg sync.WaitGroup


func main()  {
	fmt.Println("Welcome to gorutine")
	urls:=[]string{
		"http://www.youtube.com",
		"http://www.fb.com",
		"http://www.google.com",
	}

	for _,url:=range urls{
		wg.Add(1)
		go func (u string)  {
			defer wg.Done()
			status(u)
		}(url)
	}
	defer print()
	wg.Wait()
}

func print(){
	fmt.Println(statusCode);
}
func status(url string)  {
	res,err:=http.Get(url)
	if err!=nil{
		fmt.Println("Sorry server is down !!!")
	}else{
		mu.Lock()
		statusCode = append(statusCode, res.StatusCode);
		mu.Unlock()

		fmt.Printf("%d is the status code of %s\n",res.StatusCode,url)
	}
}