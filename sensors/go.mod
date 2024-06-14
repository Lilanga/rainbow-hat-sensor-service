module github.com/Lilanga/rainbow-hat-sensor-service

go 1.17

require github.com/Lilanga/sensor-data-processing-service v0.0.0-20220412065528-0dc56dd3eb8b

require (
	periph.io/x/conn/v3 v3.7.0 // indirect
	periph.io/x/host/v3 v3.8.2 // indirect
)

require (
	github.com/MichaelS11/go-dht v0.1.1
	github.com/eclipse/paho.mqtt.golang v1.3.5 // indirect
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/joho/godotenv v1.4.0
	golang.org/x/net v0.0.0-20210614182718-04defd469f4e // indirect
	periph.io/x/periph v3.6.8+incompatible
)
