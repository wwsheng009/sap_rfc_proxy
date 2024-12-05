package main

import (
	"log"
	"os"
	"os/signal"
	"sap_rfc_proxy/handlers"
	"sap_rfc_proxy/utils"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	utils.InitLogger()
	defer utils.CloseLogger()
	pool, err := handlers.NewSAPConnectionPool(10)
	if err != nil {
		utils.Logger.Fatalf("Error creating SAP connection pool: %v", err)
	}
	port := os.Getenv("PORT") // Get the port from the environment variable
	if port == "" {
		port = "8080"
	}

	r := gin.Default()
	r.POST("/rfc/call", handlers.RFCCall(pool))
	r.POST("/rfc/call1", handlers.RFCCall1)
	r.GET("/rfc/meta", handlers.RFCmeta)
	r.GET("/", func(ctx *gin.Context) {
		htmlContent := `
		<!DOCTYPE html>
		<html>
		<head>
			<title>SAP RFC PROXY</title>
		</head>
		<body>
			<h1>SAP RFC PROXY!</h1>
			<p>post request to url http://localhost:` + port + `/rfc/call?fname=<your_function_name>. the parameters set as payload map object</p>
		</body>
		</html>`
		ctx.Data(200, "text/html; charset=utf-8", []byte(htmlContent))
	})

	// Channel to listen for shutdown signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	// Run the server in a goroutine

	go func() {
		utils.Logger.Println("Starting server on port 8080...")
		if err := r.Run(":" + port); err != nil {
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
