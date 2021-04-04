package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	godotenv "github.com/joho/godotenv"
)

func main() {

	f, err := os.OpenFile("readings.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)

	err = godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	broker := os.Getenv("MQTT_URL")
	username := os.Getenv("MQTT_USERNAME")
	channel := os.Getenv("MQTT_CHANNEL")
	opts := mqtt.NewClientOptions()

	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", broker, 1883))
	opts.SetClientID("go_mqtt_client")
	opts.SetUsername(username)
	client := mqtt.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}

	output, err := exec.Command("./SHTC3").Output()

	if err != nil {
		log.Fatal(err.Error())
	}

	data := string(output)
	log.Println(data)
	client.Publish(channel, 1, false, data)
	client.Disconnect(250)
}
