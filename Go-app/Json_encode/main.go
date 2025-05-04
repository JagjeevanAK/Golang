package main

import (
	"encoding/json"
	"fmt"
)

type course struct{
	Name string `json:"name"`
	Password string `json:"-"`
	Email string `json:"email"`
	Tags []string `json:"tags,omitempty"`
	Price int16 `json:"price"`
}

func main()  {
	fmt.Println("Raw data to JSON");

	JSON();
}

func JSON()  {
	courses:=[]course{
		{"React","11223344","jagjeevankashid97@gmail.com",[]string{"web-dev","frontend"},199},
		{"C","1234","demo123@gmail.com",[]string{"backend","beginner"},99},
	}
	list,err:=json.MarshalIndent(courses,"","\t");

	if err!=nil{
		panic(err)
	}
	fmt.Printf("%s",list);
}