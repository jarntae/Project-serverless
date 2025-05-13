package main
import (
    "net/http"

    "github.com/cpe-nuntawut/assignment-6-bobox/config"
    "github.com/cpe-nuntawut/assignment-6-bobox/controller"
    "github.com/cpe-nuntawut/assignment-6-bobox/middlewares"

    "github.com/gin-gonic/gin"
)
const PORT = "8001"
func main() {
   // open connection database
    config.ConnectionDB()

   // Generate databases
    config.SetupDatabase()

    r := gin.Default()
    r.Use(CORSMiddleware())

    router := r.Group("/")
    {   
        router.Use(middlewares.Authorizes())

        // Member Route
        r.GET("/Member/:id", controller.GetMemberByID)
        r.POST("/Member", controller.CreateMember)
        r.PATCH("/UpdateMember/:id", controller.UpdateMember)

        // Ledger Route
        r.GET("/LedgerByMemberID/:id", controller.GetLedgerByMemberID)
        r.POST("/Ledger", controller.CreateLedger)
        r.PATCH("/UpdateLedger/:id", controller.UpdateLeder)
        r.DELETE("/Deleteledger/:id", controller.DeleteLedger)

        // MoneyType Routes
        r.GET("/MoneyType", controller.GetMoneyType) //PetAllow

        // CategoryType Routes
        r.GET("/CategoryType", controller.GetCategoryType)


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