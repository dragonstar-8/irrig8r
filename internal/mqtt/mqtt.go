package mqtt

import (
	"fmt"
	"log"
	"os"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

var client MQTT.Client

func Init() {
	broker := os.Getenv("MQTT_BROKER")
	if broker == "" {
		broker = "tcp://mqtt:1883"
	}

	opts := MQTT.NewClientOptions().AddBroker(broker).SetClientID("irrigation-server")
	client = MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("MQTT connect error: %v", token.Error())
	}
	log.Println("Connected to MQTT broker at", broker)
}

func PublishCommand(zone string, duration int) {
	topic := "irrigation/" + zone
	payload := fmt.Sprintf("on:%d", duration)
	token := client.Publish(topic, 1, false, payload)
	token.Wait()
	if token.Error() != nil {
		log.Printf("Failed to publish to %s: %v", topic, token.Error())
	}
}
