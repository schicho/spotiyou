package telegram

import (
	"net/http"
	"strings"

	notification "github.com/schicho/spotiyou/pkg/notification"
)

// Telegram is a Telegram notifier.
// It implements the Notifier interface.
//
// It uses the Telegram Bot API to send notifications.
// No external libraries are used. We send the requests
// directly to the Telegram API.
type TelegramNotifier struct {
	// token is the Telegram Bot API token.
	token string
	// chatID is the Telegram chat ID.
	chatID string
	// StringBuilder is used to build the notification message.
	sb *strings.Builder
}

// NewTelegramNotifier creates a new TelegramNotifier.
func NewTelegramNotifier(token string, chatID string) *TelegramNotifier {
	return &TelegramNotifier{
		token:  token,
		chatID: chatID,
		sb:     &strings.Builder{},
	}
}

// Notify sends the notification to the user.
func (t *TelegramNotifier) Notify(notification notification.Notification) error {
	if len(notification.Playlists) == 0 {
		return nil
	}
	return t.sendMessage(t.buildMessage(notification))
}

func (t *TelegramNotifier) sendMessage(message string) error {
	resp, err := http.DefaultClient.Post("https://api.telegram.org/bot"+t.token+"/sendMessage",
		"application/json",
		strings.NewReader(`{"chat_id":`+t.chatID+`,"text":"`+message+`"}`))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

func (t *TelegramNotifier) buildMessage(notification notification.Notification) string {
	t.sb.Reset()

	t.sb.WriteString("🚨New Playlists Alert!🚨\n\n")
	t.sb.WriteString(notification.Playlists[0].OwnerName)
	t.sb.WriteString(" has created the following playlists:\n")

	for _, p := range notification.Playlists {
		t.sb.WriteString("\n")
		t.sb.WriteString(p.Name)
	}

	return t.sb.String()
}
