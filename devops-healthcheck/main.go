package main

import (
	"fmt"
	"encoding/json"
	"net/http"
	"log"
	"io"
	"strconv"
	"time"
	"github.com/Bhakaresuraj/Go_language/devops-healthcheck/model"
	"github.com/Bhakaresuraj/Go_language/devops-healthcheck/database"
)
var db *database.Store

func HealthCheckHandler(w http.ResponseWriter,r *http.Request){
	fmt.Fprintln(w,"Welcome to DevOps Healthcheck.....")
	if r.Method != http.MethodPost{
		http.Error(w,"Method not allowed",http.StatusMethodNotAllowed)
		return
	}
	body,err:=io.ReadAll(r.Body)
	if err != nil{
		http.Error(w,"Error reading body",http.StatusInternalServerError)
		return
	}
	var service model.Service
	err = json.Unmarshal(body,&service)
	if err != nil{
		http.Error(w,"Error unmarshalling body",http.StatusInternalServerError)
		return
	}
	service.Healthy,service.StatusCode = service.CheckHealth()
	service.Checked_at=time.Now()
	err=db.Save(service)
	if err !=nil{
		http.Error(w,"Error Saving Service ",http.StatusInternalServerError)
		return 
	}
	json.NewEncoder(w).Encode(service)
}
func GetAllServices(w http.ResponseWriter ,r *http.Request){
	if r.Method != http.MethodGet{
		http.Error(w,"Method Not Allowed",http.StatusInternalServerError)
		return
	}
	body,err:=io.ReadAll(r.Body)
	var service model.Service
	 err =json.Unmarshal(body,&service)
	if err != nil{
		http.Error(w,"Error Unmarshalling Body :",http.StatusInternalServerError)
		return 
	}
	services,err:= db.GetAllServices(service.UserID)	
	if err!=nil{
		http.Error(w,"Error Geting Services : ",http.StatusInternalServerError)
		return 
	}
	json.NewEncoder(w).Encode(services)
}

func RunAll(w http.ResponseWriter ,r *http.Request){
	
	if r.Method != http.MethodPost{
		http.Error(w,"Method not allowed",http.StatusInternalServerError)
		return
	}
	body,err :=io.ReadAll(r.Body)
	if err!= nil{
		http.Error(w,"Error in reading Body",http.StatusInternalServerError)
		return 
	}
	var service model.Service
	err = json.Unmarshal(body,&service)
	if err!=nil{
		http.Error(w,"Error while Unmarshalling body",http.StatusInternalServerError)
		return
	}
	services,err :=db.GetAllServices(service.UserID)
	if err != nil{
		http.Error(w,"Error while fetching for update :",http.StatusInternalServerError)
		return
	}
	fmt.Println(services)
	for _,ser :=range services{

		ser.Healthy ,ser.StatusCode = ser.CheckHealth()
		ser.Checked_at =time.Now();
		err=db.UpdateServiceStatus(ser)
		if err!=nil{
			fmt.Println("Update Error :",err,)	
			continue
		}	
	}
	services,err = db.GetAllServices(service.UserID)
	if err!=nil{
		http.Error(w,"Error Geting Services : ",http.StatusInternalServerError)
		return 
	}
	json.NewEncoder(w).Encode(services)
		
}


func DeleteHandler(w http.ResponseWriter ,r *http.Request){
	if r.Method != http.MethodDelete{
		http.Error(w,"Method Not Allowed !",http.StatusMethodNotAllowed)
		return
	}
	user_id:=r.Header.Get("X-User-ID")
	if user_id == ""{
		http.Error(w,"User Id is required :" ,http.StatusBadRequest)
		return
	}
	userIdInt, err := strconv.Atoi(user_id)
	if err != nil {
		http.Error(w, "Invalid User ID", http.StatusBadRequest)
		return
	}
	service_id :=r.Header.Get("X-Service-ID")
	if service_id == ""{
		http.Error(w,"Service Id is required :" ,http.StatusBadRequest)
		return
	}
	serviceIdInt,err:=strconv.Atoi(service_id)
	if err != nil {
		http.Error(w, "Invalid User ID", http.StatusBadRequest)
		return
	}

	 err =db.DeleteService(userIdInt,serviceIdInt)
	if err != nil{
		http.Error(w,"Service not deleted ",http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w,"Successfully deleted Service")
}

func main(){
	fmt.Println("Server is Running on Port :8080")
	dbUrl := "postgres://suraj:Bhakare@localhost:5432/mydb"
	db =database.NewStore(dbUrl)
	fmt.Println("Database Connected Successfully")
	err:=db.RunMigration()
	if err != nil{
		log.Fatal("Unable to run migration :",err)
	}
	fmt.Println("Migration Run Successfully and Table Created...!")
	http.HandleFunc("/check",HealthCheckHandler)
	http.HandleFunc("/services",GetAllServices)
	http.HandleFunc("/runall",RunAll)
	http.HandleFunc("/delete",DeleteHandler)
	http.ListenAndServe(":8080", nil)
}
