package sendemailusingcron

import (
	"fmt"
	"net/smtp"
	"time"

	"github.com/robfig/cron/v3"
)

var startTime time.Time

func everyMinute() {
	now := time.Now()
	duration := now.Sub(startTime)
	fmt.Printf("Time since start: %.0f minutes\n", duration.Minutes())
}

func sendEmail() {
	from := "bahatoho.1709@gmail.com"
	password := "123456789"
	to := []string{"runtoyou1709@gmail.com"}

	// SMTP server configuration
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Authentication
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Email content
	message := []byte("To: runtoyou1709@gmail.com\r\n" +
		"Subject: Hello\r\n" +
		"\r\n" +
		"Hello, this is a test email from Go!")

	// Sending email
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email sent!")
}

func Oneminutes() {
	startTime = time.Now()
	c := cron.New()
	c.AddFunc("* * * * *", func() { everyMinute() })
	c.AddFunc("0 5 * * *", func() { sendEmail() })
	c.Start()
	defer c.Stop()
	select {}
}
