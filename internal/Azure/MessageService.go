package azure

import (
	"context"
	"log"
	"os"

	azservicebus "github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
)

type MessageService struct {
	queueName string
	client    *azservicebus.Client
}

func NewMessageService() MessageService {
	client, err := azservicebus.NewClientFromConnectionString(getConnectionString(), nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	return MessageService{
		queueName: "weather-queue",
		client:    client,
	}
}
func (m *MessageService) SendWeatherMessage(weatherData map[string]string) {
	log.Println("Preparing to send weather message...")
	go m.sendWeatherMessage(weatherData)
}

func (m *MessageService) sendWeatherMessage(weatherData map[string]string) {

	var temp string = weatherData["temp"]
	log.Printf("Sending weather message with temp: %s", temp)

	msg := &azservicebus.Message{
		Body: []byte(temp),
	}

	sender, err := m.client.NewSender(m.queueName, nil)
	if err != nil {
		log.Fatalf("failed to create sender: %v", err)
	}

	defer sender.Close(context.TODO())
	defer m.client.Close(context.TODO())

	err = sender.SendMessage(context.TODO(), msg, nil)
	if err != nil {
		log.Fatalf("failed to send message: %v", err)
	}
	log.Println("Message sent!")
}

func getConnectionString() string {
	value := os.Getenv("SB_CONNECTION_STRING")
	if value == "" {
		log.Println("⚠ env var SB_CONNECTION_STRING not set")
		value = "default-value"
	}
	log.Println("Env value:", value)
	return value
}
