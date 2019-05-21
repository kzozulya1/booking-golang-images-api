package api

import (
    "github.com/gin-gonic/gin"  
)

/*
*  Apply headers
*/
func ApplyHeaders(c *gin.Context) {
    c.Header("Access-Control-Allow-Origin", "*")   
    c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
    c.Header("Access-Control-Allow-Methods", "HEAD, GET, PUT, POST, OPTIONS")
}