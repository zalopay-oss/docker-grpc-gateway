package main

import (
	"context"
	"log"
	"os"
	cmd "service/pkg/cmd/server"
)

func main() {
	if err := cmd.RunServer(context.Background(), "8000"); err != nil {
		log.Fatal("run server err: ", err)
		os.Exit(1)
	}
}
