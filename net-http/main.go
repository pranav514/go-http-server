package main
import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const port string = ":8080"
type RequestBody struct{
	Name string `json:"name"` // maping of the struct to the expected json structure in the image
}
type ResponseBody struct{
	Message string `json:"message"`
}

func AddUser(w http.ResponseWriter , r *http.Request){
	if(r.Method != http.MethodPost){
		http.Error(w , "method not allowed" , http.StatusMethodNotAllowed)
		return
	}
	var reqBody RequestBody
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil{
		http.Error(w , "invalid json format",http.StatusBadRequest)
		return 
	}
	response := ResponseBody{
		Message: fmt.Sprintf("User %s added successfully", reqBody.Name),
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)

}
func main(){
	http.HandleFunc("/api/v1/adduser" , AddUser)
	log.Println("Server started at port" + port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}