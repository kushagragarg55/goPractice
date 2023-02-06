package main

import "fmt"

func zeroval(i int){
    i=0
}

func zeroptr(iptr *int){
    *iptr=0
}


func main(){
    i:=1
    fmt.Println("initial value of i")
    zeroval(i)
    fmt.Println("Value of i after calling zeroval function: ", i)
    fmt.Println("value of i address of i ",&i)
    zeroptr(&i)
    fmt.Println("Value of i after calling zeroptr function: ", i)
    fmt.Println("value of address of i after calling zeroptr function",&i)
}