package main
 
import (
	"github.com/gin-gonic/gin"  
     api "app/pkg/api"
     "os"
)

/*
* Setup main router
*/
func setupRouter() *gin.Engine {
	router := gin.Default()
	
    //Upload OPTIONS API Entry
    router.OPTIONS("/upload", func(c *gin.Context) {
        api.ApplyHeaders(c)
	})
    //Download GET API Entry
    router.GET("/download", func(c *gin.Context) {
        api.Download(c)
    })
    //Upload POST API Entry
    router.POST("/upload", func(c *gin.Context) {
        api.Upload(c)
	})
	return router
}

func main() {
    port := os.Getenv("APP_API_PORT")
    r := setupRouter()
	r.Run(":" + port) 
}