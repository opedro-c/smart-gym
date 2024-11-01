package main

import (
	"fmt"
	"os"
	MQTT "github.com/eclipse/paho.mqtt.golang"
)

func main() {
	opts := MQTT.NewClientOptions()
	opts.AddBroker("tcp://mosquitto:1883")
	opts.SetClientID("1")
	opts.SetUsername("admin")
	opts.SetPassword("admin")
	opts.SetCleanSession(false)

	receiveCount := 0
	choke := make(chan [2]string)
	opts.SetDefaultPublishHandler(func(client MQTT.Client, msg MQTT.Message) {
		choke <- [2]string{msg.Topic(), string(msg.Payload())}
	})

	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	qos := 0
	if token := client.Subscribe("/hello", byte(qos), nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	for receiveCount < 2 {
		incoming := <-choke
		fmt.Printf("RECEIVED TOPIC: %s MESSAGE: %s\n", incoming[0], incoming[1])
		receiveCount++
	}
}
