package main

import (
	"encoding/json"
	"log"
	"server/internal/config"
	"server/internal/repositories"
	"server/internal/services"

	"github.com/joho/godotenv"
)

type NotificationEvent struct {
	UserID  string `json:"userId"`
	Type    string `json:"type"`
	Title   string `json:"title"`
	Message string `json:"message"`
}

func main() {
	// Load env dan init koneksi
	_ = godotenv.Load()
	// config.InitDatabase()
	config.InitMailer()
	config.InitRabbitMQ()

	// Init repositori dan service
	db := config.DB
	notificationRepo := repositories.NewNotificationRepository(db)
	notificationSvc := services.NewNotificationService(notificationRepo)

	// Konsumsi pesan dari queue
	msgs, err := config.Channel.Consume(
		"notification_checkout", "", true, false, false, false, nil,
	)
	if err != nil {
		log.Fatalf("❌ Failed to consume: %v", err)
	}

	log.Println("📡 [Worker] Listening for checkout notifications...")

	for msg := range msgs {
		var payload NotificationEvent
		if err := json.Unmarshal(msg.Body, &payload); err != nil {
			log.Println("❌ Failed to parse message:", err)
			continue
		}

		log.Printf("🔔 Received notification for user %s: %s\n", payload.UserID, payload.Title)

		err := notificationSvc.SendToUser(payload.UserID, payload.Type, payload.Title, payload.Message)
		if err != nil {
			log.Printf("❌ Failed to send to user %s: %v\n", payload.UserID, err)
		}
	}
}
