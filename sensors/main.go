package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	mqttLib "github.com/Lilanga/sensor-data-processing-service/pkg/mqtt"
	"github.com/MichaelS11/go-dht"
	"github.com/joho/godotenv"

	_ "embed"
)

//go:embed index.htm
var indexHTMLData []byte

type SensorData struct {
	Humidity    string `json:"humidity"`
	Temperature string `json:"temperature"`
	Pressure    string `json:"pressure"`
	SensorID    string `json:"sensor_id"`
	Timestamp   string `json:"timestamp"`
}

var currentData SensorData

func main() {
	godotenv.Load(".env")

	if err := initializeHardware(); err != nil {
		log.Fatalf("Failed to initialize hardware: %v", err)
	}

	dhtSensor, err := dht.NewDHT("GPIO2", dht.Celsius, "DHT22")
	if err != nil {
		fmt.Println("NewDHT error:", err)
		return
	}

	interval, err := strconv.Atoi(os.Getenv("REFRESH_INTERVAL"))

	if err != nil {
		interval = 30
	}

	ticker := time.NewTicker(time.Duration(interval) * time.Second)
	defer ticker.Stop()

	client, topic, sensorID := setupMQTT()
	// defer client.Disconnect(250)

	go publishSensorData(dhtSensor, ticker, client, topic, sensorID)

	// Serve HTTP
	http.HandleFunc("/", webpageHandler)
	http.HandleFunc("/data", dataHandler)
	go func() {
		log.Fatal(http.ListenAndServe(os.Getenv("PORT"), nil))
	}()

	handleShutdown()
}

func initializeHardware() error {
	if err := dht.HostInit(); err != nil {
		return fmt.Errorf("host initialization failed: %w", err)
	}
	return nil
}

func setupMQTT() (*mqttLib.MqttClient, string, string) {
	clientID := os.Getenv("MQTT_CLIENT_ID")
	topic := os.Getenv("MQTT_TOPIC")
	sensorID := os.Getenv("ID")

	if clientID == "" || topic == "" || sensorID == "" {
		log.Fatalf("MQTT configuration is missing. Check .env file")
	}

	client := mqttLib.GetMqttClient(clientID)
	return client, topic, sensorID
}

func publishSensorData(sensor *dht.DHT, ticker *time.Ticker, client *mqttLib.MqttClient, topic, sensorID string) {
	for range ticker.C {
		humidity, temperature, err := sensor.ReadRetry(11)
		if err != nil {
			fmt.Println("Read error:", err)
			continue
		}

		currentData = SensorData{
			Humidity:    fmt.Sprintf("%v", humidity),
			Temperature: fmt.Sprintf("%v", temperature),
			Pressure:    "0",
			SensorID:    sensorID,
			Timestamp:   time.Now().Format(time.RFC3339),
		}

		jsonData, err := json.Marshal(currentData)
		if err != nil {
			log.Printf("Error marshalling JSON: %v", err)
			continue
		}

		client.Publish(topic, 0, jsonData)
	}
}

func webpageHandler(w http.ResponseWriter, r *http.Request) {
	// Serve the HTML page
	// indexHTMLData, err := emdbFsData.ReadFile("index.html")

	if indexHTMLData == nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.Write(indexHTMLData)
}

func dataHandler(w http.ResponseWriter, r *http.Request) {
	// Serve the JSON data
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(currentData)
}

func handleShutdown() {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		log.Printf("Received signal: %s", sig)
		done <- true
	}()

	<-done
	log.Println("Shutting down gracefully...")
}
