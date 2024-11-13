package mqtt

import (
	"encoding/json"
	"fmt"
	"gym/internal/business"
	"gym/internal/database/sqlite"
	"gym/pkg/abstract"
	"log/slog"
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

func (s *Subscriber) Setup() {

	s.mqttClient.Subscribe("/gym/exercise", func(payload []byte) {

		db := sqlite.GetConnection()
		tx, err := db.Begin()

		if err != nil {
			slog.Error(fmt.Sprintf("Error starting transaction: %v", err))
		}

		var input business.InputSaveExerciseUseCase
		json.Unmarshal(payload, &input)

		abstract.NewUseCaseTransaction(
			tx,
			business.NewSaveExerciseUseCase(sqlite.NewExerciseRepository(tx)),
			input,
		).Execute()
	})
}
