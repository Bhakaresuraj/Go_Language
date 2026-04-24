package main

import (
    "fmt"
)

func main(){
        fmt.Println("THis is for defer....:")

        defer fmt.Println("def 8")
         defer fmt.Println("def 7")
  defer       fmt.Println("def 6")
        defer fmt.Println("def 5")
  defer       fmt.Println("def 4")
        defer fmt.Println("def 3")
  defer       fmt.Println("def 2")
        defer fmt.Println("def 1")
        fmt.Println("End of main")
}
