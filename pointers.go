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

    zeroptr(&i)
    fmt.Println("Value of i after calling zeroptr function: ", i)
}