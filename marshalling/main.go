package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func pingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Pong")
}
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is serving from Port :8080")
	fmt.Println("Method : ", r.Method)
	fmt.Println("Query :", r.URL.RawQuery)
	fmt.Println("Path :", r.URL.Path)
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "You called  : %s  %s\n", r.Method, r.URL.Path)

}

func statusReader(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, r.Body)
	data, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error Reading Body :", err)
	}
	fmt.Fprintf(w, "Data in the form of Byte Stream : \n ", data)
	fmt.Fprintf(w, "\nString Data :--\n", string(data))

}

type UserInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Location string `json:"location"`
}

type Result struct {
	Math int `json:"math"`
	Mar  int `json:"mar"`
	Eng  int `json:"eng"`
}

func login(i UserInfo) (Result, error) {
	var result Result
	if i.Username == "suraj" && i.Password == "Pass" {
		result = Result{Math: 99, Mar: 50, Eng: 91}
		return result, nil
	}
	return result, fmt.Errorf("Please login by  Valid User details....")
}

func Marshalling(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error while reading Body :", err)
	}
	var data UserInfo
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("Error during Unmarshalling :", err)
	}
	result, err := login(data)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintln(w, err)
		return
	}

	resultJson, err := json.Marshal(result)
	w.WriteHeader(200)
	fmt.Fprintln(w, string(resultJson))
}

func main() {
	// http.HandleFunc("/", homeHandler)
	// http.HandleFunc("/read", statusReader)
	// http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/result", Marshalling)
	fmt.Println("Server is Running on Port :8080")
	http.ListenAndServe(":8080", nil)

}

