package main

import (
	"net/http"
    "github.com/gin-gonic/gin"
	"github.com/cpe-nuntawut/assignment-6-bobox/config"
    "github.com/cpe-nuntawut/assignment-6-bobox/controller"
	 "github.com/cpe-nuntawut/assignment-6-bobox/middlewares"
)

const PORT = "8001"

func main() {
   // open connection database
    config.ConnectionDB()

   // Generate databases
    config.SetupDatabase()

    r := gin.Default()
    r.Use(CORSMiddleware())

   // Auth Route
    r.POST("/signIn", controller.SignIn)

    router := r.Group("/")
    {   
        router.Use(middlewares.Authorizes())


}

    r.GET("/", func(c *gin.Context) {
        c.String(http.StatusOK, "API RUNNING... PORT: %s", PORT)
    })

    // Run the server
    r.Run("0.0.0.0:" + PORT)
}

func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }
        c.Next()
    }
}