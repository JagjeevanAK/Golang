package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func main()  {
	fmt.Println("Welcome to gorutine")
	url_list:=[]string{
		"https://www.google.com",
		"https://www.github.com",
		"https://www.fb.com",
	}
	for _,i:=range url_list{
		wg.Add(1)
		go func (u string)  {
			defer wg.Done()
			getstatuscode(u)
		}(i)
	}
	wg.Wait()
}
 
func getstatuscode(url string){
	out,err:=http.Get(url);

	if err!=nil{
		fmt.Println("Failed")
	}else{
		fmt.Printf("%d status code of %s\n",out.StatusCode,url)
	}
}