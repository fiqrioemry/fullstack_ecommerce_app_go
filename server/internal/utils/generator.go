package utils

import (
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

func RandomUserAvatar(avatar string) string {
	return fmt.Sprintf("https://api.dicebear.com/6.x/initials/svg?seed=%s", avatar)
}

func GenerateOTP(length int) string {
	digits := "0123456789"
	var sb strings.Builder

	for range length {
		sb.WriteByte(digits[rand.Intn(len(digits))])
	}

	return sb.String()
}

func GenerateSlug(input string) string {

	slug := strings.ToLower(input)
	re := regexp.MustCompile(`[^a-z0-9]+`)
	slug = re.ReplaceAllString(slug, "-")
	slug = strings.Trim(slug, "-")

	suffix := strconv.Itoa(rand.Intn(1_000_000))
	slug = slug + "-" + leftPad(suffix, "0", 6)

	return slug
}

func leftPad(s string, pad string, length int) string {
	for len(s) < length {
		s = pad + s
	}
	return s
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func GenerateInvoiceNumber(orderID uuid.UUID) string {
	timestamp := time.Now().Format("20060102")
	shortID := orderID.String()[:8]
	return fmt.Sprintf("INV/%s/%s", timestamp, shortID)
}
