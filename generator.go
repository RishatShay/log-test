package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	fmt.Println("Starting generator")

	levels := []string{"INFO", "WARN", "ERROR", "DEBUG"}
	services := []string{"auth-api", "payment-worker", "frontend-app", "admin-panel"}
	messages := []string{
		"Message 1",
		"Message 2",
		"Message 3",
		"Message 4",
	}

	url := "http://localhost:9002/_bulk"
	client := &http.Client{Timeout: 2 * time.Second}

	for {
		lvl := levels[rand.Intn(len(levels))]
		svc := services[rand.Intn(len(services))]
		msg := messages[rand.Intn(len(messages))]

		payload := fmt.Sprintf("{\"index\":{}}\n{\"level\":\"%s\", \"service\":\"%s\", \"message\":\"%s\"}\n", lvl, svc, msg)

		req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(payload)))
		if err != nil {
			fmt.Println("Error compiling the request", err)
			continue
		}
		req.Header.Set("Content-Type", "application/json")

		resp, err := client.Do(req)
		if err != nil {
			fmt.Printf("Ошибка отправки! %v\n", err)
		} else {
			fmt.Printf("Отправлен лог [%5s] от %-15s (HTTP Status: %d)\n", lvl, svc, resp.StatusCode)
			resp.Body.Close()
		}

		time.Sleep(1 * time.Second)
	}
}
