package main

import (
	"errors"
	"fmt"
)

func main(){
    printhello("Jagjeevan");
    fmt.Println(add(5,6));
    qutiont,remainder,err:=divide(10,0);
    if (err!=nil){
        fmt.Println(err.Error());
    }else{
        fmt.Print("Quotient: ",qutiont," Remainder: ",remainder);
    }
}

func printhello(str string){
    fmt.Println("Hello "+str);
}

func add (a int, b int) int{
    return a+b;
}

func divide (a int, b int ) (float64,int,error){
    if(b==0){
        return 0,0,errors.New("Cannot divide by zero")
    }

    return float64(a/b),a%b,nil;
}
