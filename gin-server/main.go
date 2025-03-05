package main 
import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type RequestBody struct{
	Name string `json:"name"`
}
type ResponseBody struct{
	Message string `json:"message"`
}
func AddUser(c *gin.Context){
	var reqBody RequestBody
	err := c.ShouldBindBodyWithJSON(&reqBody)
	if(err != nil){
		c.JSON(http.StatusBadRequest , gin.H{
			"error": "invalid json format",
		})
		return 
	}
	response := ResponseBody{
		Message:   "User " + reqBody.Name + "added successfully",
	}
	c.JSON(http.StatusOK , response)

}
func main(){
	r := gin.Default()
	r.POST("/api/v1/adduser" , AddUser)
	port := ":8080"
	r.Run(port)
}