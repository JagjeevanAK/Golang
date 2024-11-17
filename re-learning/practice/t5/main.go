package main

import "fmt"

func main() {
    //Array
	var intarr [3]int;
	fmt.Println(intarr[0]);
    fmt.Println(intarr[1:3]);

    //Slice
    var intslice = []int{1,2,3};
    fmt.Println(intslice);

    intslice = append(intslice,4,5);
    fmt.Println(intslice);
}
