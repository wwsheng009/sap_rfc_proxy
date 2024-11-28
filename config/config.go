package config

import (
	"log"
	"os"

	"sap_rfc_proxy/gorfc"

	"github.com/joho/godotenv"
)

func LoadConfig() gorfc.ConnectionParameters {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return gorfc.ConnectionParameters{
		"Dest":      os.Getenv("SAP_DEST"),
		"Client":    os.Getenv("SAP_CLIENT"),
		"User":      os.Getenv("SAP_USER"),
		"Passwd":    os.Getenv("SAP_PASSWD"),
		"Lang":      os.Getenv("SAP_LANG"),
		"Ashost":    os.Getenv("SAP_ASHOST"),
		"Sysnr":     os.Getenv("SAP_SYSNR"),
		"Saprouter": os.Getenv("SAP_SAPROUTER"),
	}
}
