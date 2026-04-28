package main

import (
    "fmt"
   "io" 
    "net/http"
)

func homeHandler(w http.ResponseWriter , r *http.Request){
    fmt.Println("=================Incoming Request=============")
    fmt.Println("method :", r.Method)
    fmt.Println("Path :" ,r.URL.Path)
    fmt.Println("Query :",r.URL.RawQuery)
    fmt.Println("Headers :",r.Header)
    fmt.Println("content Length :",r.ContentLength)
    fmt.Println("Host  :",r.Host)
    fmt.Println("Body :", r.Body)

    body,err :=io.ReadAll(r.Body)
    if err != nil{
        fmt.Println("Error Reading Body :",err)
        return
    }
    fmt.Println("Body data :" ,string(body))

    fmt.Fprintln(w,"Hello from http-server......")

}

func main(){
        fmt.Println("This is Http server......")

        http.HandleFunc("/",homeHandler)

        fmt.Println("Server is running on Port :8080")
        http.ListenAndServe(":8080",nil)
}
