package main
 
import (
    "fmt"
    "github.com/fatih/color"
    "os"
    "strings"
    "strconv"

)

func main(){

    fmt.Println("Hello this is Go program for learning package :")
    color.Green("This is printed using color package")
    color.Red("This is printed using color package")

    // Welcome to  os package 
    port:=8080  //create and assign  ,integer type

    msg :=fmt.Sprintf("Port : %d",port)
    fmt.Println(msg);

    // os.Getenv function returns string value
    portStr :=os.Getenv("MYAPP_PORT")
    fmt.Println("Port : ",portStr)


    // Use of string convert package

    // strconv.Atoi()  returns two values one is result and another is error 

    portInt ,err :=strconv.Atoi(portStr)

    if err != nil {
        fmt.Println("Error :",err)
        os.Exit(1)
    }

    fmt.Println("Port Int  :",portInt)


    //host variable

    host:=os.Getenv("MYAPP_HOST")

    fmt.Println("Host : ",host)
    
    // accept our service name

    service :=os.Getenv("MYAPP_NAME")

    fmt.Println("Name:",service)

    // string manipulation

    fmt.Println(strings.ToUpper(service))



}
