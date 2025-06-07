package mqtt

import (
	"testing"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

func TestInitAndPublish(t *testing.T) {
	// Setup
	opts := MQTT.NewClientOptions()
	opts.AddBroker("tcp://localhost:1883")
	opts.SetClientID("test-client")
	opts.SetConnectTimeout(2 * time.Second)

	c := MQTT.NewClient(opts)
	token := c.Connect()
	if !token.WaitTimeout(3*time.Second) || token.Error() != nil {
		t.Fatalf("Failed to connect to MQTT broker: %v", token.Error())
	}
	defer c.Disconnect(250)

	topic := "irrigation/test"
	payloadCh := make(chan string, 1)

	// Subscribe to test topic
	if token := c.Subscribe(topic, 1, func(_ MQTT.Client, msg MQTT.Message) {
		payloadCh <- string(msg.Payload())
	}); token.Wait() && token.Error() != nil {
		t.Fatalf("Subscribe failed: %v", token.Error())
	}

	// Act: publish a test command using your package
	client = c // Override internal client for test
	PublishCommand("test", 42)

	// Assert: receive the message
	select {
	case msg := <-payloadCh:
		expected := "on:42"
		if msg != expected {
			t.Errorf("Unexpected payload: got %q, want %q", msg, expected)
		}
	case <-time.After(2 * time.Second):
		t.Error("Did not receive MQTT message")
	}
}
