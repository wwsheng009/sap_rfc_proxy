package main

import (
	"os"
	"os/signal"
	"sap_rfc_proxy/handlers"
	"sap_rfc_proxy/utils"
	"syscall"

	"github.com/gin-gonic/gin"
)

func main() {

	utils.InitLogger()
	defer utils.CloseLogger()
	pool, err := handlers.NewSAPConnectionPool(10)
	if err != nil {
		utils.Logger.Fatalf("Error creating SAP connection pool: %v", err)
	}
	r := gin.Default()
	r.POST("/rfc/call", handlers.RFCCall(pool))
	r.POST("/rfc/call1", handlers.RFCCall1)
	r.GET("/rfc/meta", handlers.RFCmeta)

	// Channel to listen for shutdown signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	// Run the server in a goroutine
	go func() {
		utils.Logger.Println("Starting server on port 8080...")
		if err := r.Run(":8080"); err != nil {
			utils.Logger.Fatalf("Server failed to start: %v", err)
		}
	}()

	// Wait for shutdown signal
	<-quit
	utils.Logger.Println("Shutting down server...")

	// Gracefully close connections and release resources
	pool.CloseAllConnections()

	utils.Logger.Println("Server stopped")
}
