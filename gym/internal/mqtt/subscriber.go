package mqtt

import (
	"gym/internal/business"
	"gym/pkg/logger"
)

// Subscriber routes each incoming message to the appropriate handler
//
// The main idea behind this struct is to separate the responsability of
// subscribing to the topics that we are going to use from
// the mqtt client itself
type Subscriber struct {
	mqttClient Client // mqttClient is used to subscribe to topics
}

func NewSubscriber(mqttClient Client) *Subscriber {
	return &Subscriber{mqttClient}
}

func (s *Subscriber) Setup() error {
	err := s.mqttClient.Subscribe("/gym/exercise", func(payload []byte) {
		logger.Logger().Println(string(payload[:]))
		(&business.AuthService{
			Test: 1,
		}).TestFunc()
	})

	if err != nil {
		return err
	}
	return nil
}
