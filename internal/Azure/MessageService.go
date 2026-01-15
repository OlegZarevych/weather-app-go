package azure

import (
	"context"
	"log"
	"os"

	azservicebus "github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
)

type MessageService struct {
	queueName   string
	client      *azservicebus.Client
	messageChan chan map[string]string
}

func NewMessageService() MessageService {
	client, err := azservicebus.NewClientFromConnectionString(getConnectionString(), nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	ms := MessageService{
		queueName:   "weather-queue",
		client:      client,
		messageChan: make(chan map[string]string, 10), // buffered channel because capacity is 10
	}
	// start a background goroutine to process messages
	go func() {
		for data := range ms.messageChan {
			ms.sendWeatherMessage(data)
		}
	}()
	return ms
}
func (m *MessageService) SendWeatherMessage(weatherData map[string]string) {
	log.Println("Preparing to send weather message...")
	// Instead of spawning a new goroutine per call, it now sends the weatherData to the channel
	//  is non-blocking (due to the buffered channel)
	m.messageChan <- weatherData
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
