package main

import (
	"ShelfQL/cmd"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	//homeDir, _ := os.UserHomeDir()
	//envPath := filepath.Join(homeDir, ".shelfql.env")
	//
	//err := godotenv.Load(envPath)

	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Warning: No .env file found, using system environment variables")
	}

	cmd.Execute()

	if cmd.DB != nil {
		cmd.DB.Close()
	}
}
