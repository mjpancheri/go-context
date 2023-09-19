package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	// if this timeout is less than the server's timeout, it will fail
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*6)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:8080", nil)
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Error executing request: %v", err)
	}
	defer res.Body.Close()
	io.Copy(os.Stdout, res.Body)
}
