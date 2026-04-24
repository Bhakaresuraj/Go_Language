p
ckage main
import (
    "fmt"
    "net/http"
)


func main(){
    
    ch := make(chan string)

    go func(){
        resp,err :=http.Get("https://httpbin.org/status/200")

         if err != nil {
             ch <- "❌ failed ....."
         }   
        defer resp.Body.Close()
        ch <- fmt.Sprintf("✅status: %d", resp.StatusCode) 
    }()



    result := <- ch
    fmt.Println(result);

}
