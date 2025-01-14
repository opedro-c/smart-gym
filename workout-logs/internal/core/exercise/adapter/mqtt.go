package adapter

import (
	"cloud-gym/internal/core/exercise"
	"cloud-gym/internal/core/exercise/usecases"
	"cloud-gym/internal/mongo"
	utils "cloud-gym/pkg"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"log/slog"
)

func CreateExerciseMQTTHandler(client MQTT.Client, msg MQTT.Message) {
	var input []exercise.ExerciseRecord

	if err := utils.ParseMQTTPayload(msg, &input); err != nil {
		slog.Error(err.Error())
		return
	}

	for _, data := range input {
		if err := utils.ValidateJsonStruct(&data); err != nil {
			slog.Error(err.Error())
			return
		}
	}

	useCase := usecases.NewCreateExercises(exercise.NewMongoExerciseRepository(mongo.GetConnection()))

	_, err := useCase.Execute(input)

	if err != nil {
		slog.Error(err.Error())
		return
	}
}
