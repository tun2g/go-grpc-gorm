package main

import (
	"fmt"
	"net"

	server "app/src"
	"app/src/config"
	"app/src/database"
	"app/src/lib/logger"
)

func main() {
	
	log := logger.NewLogger("Main")

	connection := database.InitDB()

	s, err := server.NewServer(connection, &config.AppConfiguration)

	if(err != nil) {
		log.Errorf("Error creating server: %v", err)
	}

	s.RegisterService()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", config.AppConfiguration.AppPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Println("Server is running on port", config.AppConfiguration.AppPort)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
