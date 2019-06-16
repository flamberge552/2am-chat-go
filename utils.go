package main

import (
	"fmt"
	"os"
)

func currentContext() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		fmt.Printf("$PORT either undeclared or empty, falling back to default port %s", port)
	}
	return ":" + port
}
