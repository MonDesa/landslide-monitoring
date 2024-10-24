package app

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/rubenszinho/landslide-monitoring/internal/mqttclient"
	// "github.com/rubenszinho/landslide-monitoring/internal/sensors"
)

func RunCoordinator() {
	err := godotenv.Load("config/config.env")
	if err != nil {
		panic("Error loading config.env file")
	}

	broker := os.Getenv("MQTT_SERVER")
	clientID := "coordinator"

	client := mqttclient.NewClient(broker, clientID)

	// Coordinator are not responsible to publish sensor data yet, they directly publish to broker through usb driver
	// go sensors.PublishSensorData(client)

	go PublishHealthMetrics(client)
	SubscribeControlTopics(client)

	// Keep the main routine running
	select {}
}
