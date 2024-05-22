package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/buabaj/fairness-engine/pkg/server"
)

func main() {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8080
	}

	srv := server.NewServer(port)
	err = srv.Run()
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
		os.Exit(1)
	}
}
