package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	//"fmt"
	//"encoding/json"
	//"log"
	"os"
)

type Product struct {
	Name      string
	Price     int
	Published bool
}


func main(){
	productos,err:=os.Open("./productos.json")
	if err!=nil {
		log.Fatal(err)
	}
	fmt.Print(productos)

	router:=gin.Default()
	router.GET("/hello-world", func(c *gin.Context) {
		c.JSON(200, gin.H{
		"message": "Hello World!",
		})
		})
		router.GET("/productos", func(c *gin.Context) {
			c.JSON(200, gin.H{
			"message": "Productos",
			})
			})
		router.Run()  
}