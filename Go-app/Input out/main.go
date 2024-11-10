package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader:=bufio.NewReader(os.Stdin);

	fmt.Println("Enter your name : ");
	input,_:=reader.ReadString('\n');
	input= strings.TrimSpace(input);

	age,err:=strconv.Atoi(input);

	if err!=nil{
		fmt.Println("Please enter the valid input !!");
		return
	}
	if age>=18{
		fmt.Println("You are eligible to vote");
	} else{
		fmt.Println("You are not eligible to vote");
	}
}
