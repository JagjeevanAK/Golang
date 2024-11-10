package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://www.google.com"

func main() {
	fmt.Println("Calling google");

	response,err:=http.Get(url);

	if err!=nil{
		fmt.Println("There is problem while calling the external server ")
		panic(err);
	}
	fmt.Printf("Response type: %T\n",response);
	defer response.Body.Close();

	databytes,err:=ioutil.ReadAll(response.Body);
	if err!=nil{
		panic(err);
	}

	content:=string(databytes);
	fmt.Println(content);
}