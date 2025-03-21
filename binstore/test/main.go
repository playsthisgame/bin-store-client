package main

import (
	"crypto/rand"
	"io"
	"log/slog"
	"time"

	"github.com/playsthisgame/bin-store-client/binstore"
)

func main() {
	var client, err = binstore.Connect(&binstore.Config{"localhost", 3000})

	key := time.Now().Unix()

	file := make([]byte, 100)
	_, err = io.ReadFull(rand.Reader, file)
	if err != nil {
		slog.Error("Error creating file:", "error", err)
	}

	queueName := []byte("test_queue")

	// client.Write(key, &file)
	client.Write(key, &queueName)

	// client.Read(key)

}
