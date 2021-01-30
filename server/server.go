package server

import (
	"../controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/check", controllers.Check())
	r.Run(":8080")
}
