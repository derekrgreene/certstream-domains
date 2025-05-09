package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/coder/websocket"
)

const outputFile = "domains_data/domains.jsonl"

type IncomingMessage struct {
	Data interface{} `json:"data"`
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	go func() {
		<-sig
		log.Println("Shutting down...")
		cancel()
	}()

	// Connection to certstream-server-go websocket - update to localhost or otherwise if not running in Docker
	conn, _, err := websocket.Dial(ctx, "ws://certstream:8080/domains-only", nil)
	if err != nil {
		log.Fatalf("WebSocket connection failed: %v", err)
	}
	defer conn.Close(websocket.StatusNormalClosure, "done")
	log.Println("Connected to ws://localhost:8080/")
	file, err := os.OpenFile(outputFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	count := 0 // domain counter

	for {
		_, data, err := conn.Read(ctx)
		if err != nil {
			log.Printf("WebSocket read error: %v", err)
			break
		}

		var msg IncomingMessage
		if err := json.Unmarshal(data, &msg); err != nil {
			continue
		}

		if rawList, ok := msg.Data.([]interface{}); ok {
			for _, d := range rawList {
				if domain, ok := d.(string); ok {
					jsonDomain, err := json.Marshal(domain)
					if err != nil {
						continue
					}
					if _, err := file.Write(append(jsonDomain, '\n')); err != nil {
						log.Printf("Write error: %v", err)
					}
					count++
					fmt.Printf("\rDomains collected: %d", count)
					fmt.Printf("%*s", 0, "")
				}
			}
		}
	}
}
