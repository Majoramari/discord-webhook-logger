package webhook

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/majoramari/visitor-logger/visitor"
)

func SendToWebhook(visitorInfo visitor.VisitorInfo) error {
	err := godotenv.Load(".env")
	if err != nil {
		return fmt.Errorf("error loading .env file: %v", err)
	}

	webhookURL := os.Getenv("WEBHOOK_URL")
	if webhookURL == "" {
		return fmt.Errorf("webhook URL not found in environment variables")
	}

	payload := CreatePayload(visitorInfo)

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to encode payload as JSON: %v", err)
	}

	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("failed to send data to webhook: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected response from webhook: %s", resp.Status)
	}

	fmt.Println("Visitor information sent to webhook successfully")
	return nil
}
