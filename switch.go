package main

import (
    "fmt"
    "time"
)

func main(){
    i:=2
    fmt.Print("Write ",i," as ")
    switch i {
    case 1:
        fmt.Println("one")
    case 2:
        fmt.Println("two")
    case 3:
        fmt.Println("three")
    }

    switch time.Now().Weekday(){
        case time.Saturday, time.Sunday:
            fmt.Println("It's a weekend")
        default:
            fmt.Println("It's a weekday")
    }

    t:=time.Now()
    switch {
    case t.Hour()<12:
        fmt.Println("It's before noon")
    default:
        fmt.Println("It's after noon")
    }

    whatAmI := func(i interface{}){
        switch t := i.(type){
            case bool:
                fmt.Println("I am a bool")
            case int:
                fmt.Println("I am a integer")
            default:
                fmt.Printf("I not aware about %T\n",t)
        }
    }

    whatAmI(true)
    whatAmI(false)
    whatAmI(1)
    whatAmI("lolzz")
    whatAmI(2.44)

}