package main

import (
	"github.com/ISSuh/my-stream-auth/internal"
)

func main() {
	server := internal.InitializeServer()
	server.Handle()
	server.Run()
}
