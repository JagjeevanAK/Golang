//Switch statement

package main

import (
	"errors"
	"fmt"
)

func main(){
    var result,remainder,err =divide(10,0);
    switch{
        case err!=nil:
            fmt.Println(err.Error());
    case remainder==0:
        fmt.Println("Result: ",result);
    default:
        fmt.Println("Result: ",result," Remainder: ",remainder);
    }

    switch (remainder){
        case 0:
            fmt.Println("Result: ",result);
        default:
            fmt.Println("Result: ",result," Remainder: ",remainder);
    }
}

func divide(numinator int,denominator int)(int,int,error){
    if denominator==0 {
        return 0,0,errors.New("Cannot divide by zero")
    }
    return numinator/denominator,numinator%denominator,nil;
}
