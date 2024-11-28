package main

import (
	"sap_rfc_proxy/handlers"
	"sap_rfc_proxy/utils"

	"github.com/gin-gonic/gin"
)


func main() {
	utils.InitLogger()
	defer utils.CloseLogger()

	r := gin.Default()
	r.POST("/rfc/call", handlers.RFCCall)
	r.GET("/rfc/meta", handlers.RFCmeta)

	utils.Logger.Println("Starting server on port 8080...")
	r.Run(":8080")
}
