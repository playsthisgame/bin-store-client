package main

import (
	"log/slog"
	"os"
	"time"

	"github.com/playsthisgame/bin-store-client/binstore"
)

// import (
// 	"github.com/playsthisgame/bin-store-client"
// )

func main() {
	c, err := binstore.Connect(&binstore.Config{
		Host: "localhost",
		Port: 3000,
	})
	if err != nil {
		slog.Error("Error connecting to server", "error", err)
		os.Exit(1)
	}

	defer c.Close()

	// load
	name := "my_data"
	err = c.Load(name)
	slog.Error("error loading data", "error", err)

	// write
	key := time.Now().Unix()
	data := []byte("test")
	err = c.Write(key, &data)
	if err != nil {
		slog.Error("Error writing data", "error", err)
	}

	// read
	read, err := c.Read(key)
	if err != nil {
		slog.Error("Error read data", "error", err)
	}

	slog.Info("Received data", "data", string(*read))

	// store
	err = c.Store("my_data")
	if err != nil {
		slog.Error("Error storing data", "error", err)
	}

}
