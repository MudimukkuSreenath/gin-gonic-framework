package main
import(
	"fmt"
	"github.com/gin-gonic/gin"
)
func Homepage(c *gin.Context){
c.JSON(200,gin.H{
"message":"Hello World1",
})
}
func postHomepage(c *gin.Context){
c.JSON(200,gin.H{
"message":"post Homepage",
})
}
func QueryString(c *gin.Context){
name:=c.Query("name")
age:=c.Query("age")
c.JSON(200,gin.H{
"name":name,
"age":age,
})
}
func pathparameters(c *gin.Context){
name:=c.Param("name")
age:=c.Param("age")
c.JSON(200,gin.H{
"name":name,
"age":age,
})
}
func main(){
fmt.Println("Hello World")
r:=gin.Default()
r.GET("/",Homepage)
r.POST("/",postHomepage)
r.GET("/query",QueryString)
r.GET("/path/:name/:age",pathparameters)
r.Run()
}
