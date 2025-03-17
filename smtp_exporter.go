package smtp_exporter

import (
	"fmt"
	"log"
	"net/smtp"
)

// EmailConfig содержит данные для отправки письма
type EmailConfig struct {
	SMTPServer string   // SMTP-сервер
	SMTPPort   int      // Порт
	From       string   // Адрес отправителя
	To         []string // Список получателей
	Subject    string   // Тема письма
	Body       string   // HTML-контент письма
}

// SendEmail отправляет письмо БЕЗ TLS и аутентификации
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

	// Подключаемся к SMTP-серверу без TLS
	client, err := smtp.Dial(serverAddr)
	if err != nil {
		log.Printf("Ошибка соединения с SMTP-сервером: %v", err)
		return err
	}
	defer client.Close()

	// Указываем отправителя
	if err := client.Mail(cfg.From); err != nil {
		log.Printf("Ошибка MAIL FROM: %v", err)
		return err
	}

	// Добавляем получателей
	for _, addr := range cfg.To {
		if err := client.Rcpt(addr); err != nil {
			log.Printf("Ошибка RCPT TO (%s): %v", addr, err)
			return err
		}
	}

	// Записываем данные письма
	wc, err := client.Data()
	if err != nil {
		log.Printf("Ошибка при передаче данных: %v", err)
		return err
	}
	defer wc.Close()

	_, err = wc.Write(message)
	if err != nil {
		log.Printf("Ошибка записи письма: %v", err)
		return err
	}

	log.Println("Письмо успешно отправлено без TLS!")
	return nil
}
