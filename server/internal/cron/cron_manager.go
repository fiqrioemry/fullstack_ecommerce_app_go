package cron

import (
	"log"

	"server/internal/services"

	"github.com/robfig/cron/v3"
)

type CronManager struct {
	c                   *cron.Cron
	paymentService      services.PaymentService
	notificationService services.NotificationService
}

func NewCronManager(
	payment services.PaymentService,
	notification services.NotificationService,
) *CronManager {
	return &CronManager{
		c:                   cron.New(cron.WithSeconds()),
		paymentService:      payment,
		notificationService: notification,
	}
}

func (cm *CronManager) RegisterJobs() {
	cm.c.AddFunc("0 12 * * *", func() {
		log.Println("Cron: Checking expired pending payments...")
		if err := cm.paymentService.ExpireOldPendingPayments(); err != nil {
			log.Println("Error expiring payments:", err)
		} else {
			log.Println("Payment status updated (pending → failed)")
		}
	})
}

func (cm *CronManager) Start() {
	cm.c.Start()
	log.Println("Cron Manager started")
}
