package notifications

import "github.com/sirupsen/logrus"

// Sender represents a delivery mechanism for user notifications.
type Sender interface {
	Send(userID uint, channel, message string) error
}

// LogSender writes notifications to the application log for visibility in local environments.
type LogSender struct{}

// NewLogSender constructs a logging notification sender.
func NewLogSender() *LogSender {
	return &LogSender{}
}

// Send logs the outgoing notification.
func (l *LogSender) Send(userID uint, channel, message string) error {
	logrus.WithFields(logrus.Fields{"user_id": userID, "channel": channel}).Info(message)
	return nil
}
