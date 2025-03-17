package smtp_exporter

import (
	"fmt"
	"log"
	"net/smtp"
)

// EmailConfig содержит данные для отправки письма
type EmailConfig struct {
	SMTPServer string   // SMTP-сервер
	SMTPPort   int      // Порт (например, 25)
	From       string   // Адрес отправителя
	To         []string // Список получателей
	Subject    string   // Тема письма
	Body       string   // HTML-контент письма
}

// SendEmail отправляет письмо без аутентификации
func SendEmail(cfg EmailConfig) error {
	// Формируем заголовки письма
	headers := fmt.Sprintf(
		"From: %s\r\nTo: %s\r\nSubject: %s\r\nMIME-Version: 1.0\r\nContent-Type: text/html; charset=UTF-8\r\n\r\n",
		cfg.From,
		cfg.To,
		cfg.Subject,
	)

	// Полное письмо (заголовки + тело)
	message := []byte(headers + cfg.Body)

	// Адрес SMTP-сервера
	serverAddr := fmt.Sprintf("%s:%d", cfg.SMTPServer, cfg.SMTPPort)

	// Отправка без аутентификации
	err := smtp.SendMail(serverAddr, nil, cfg.From, cfg.To, message)
	if err != nil {
		log.Printf("Ошибка отправки письма: %v", err)
		return err
	}

	log.Println("Письмо успешно отправлено!")
	return nil
}
