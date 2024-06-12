# Rainbow HAT with Raspberry Pi on BalenaOS

This repository contains a Go application for reading sensor data from the Rainbow HAT on a Raspberry Pi running BalenaOS. The application reads temperature, humidity, and pressure data, and publishes it via MQTT while also serving a simple web page to display the current sensor data.

## Table of Contents

- [Requirements](#requirements)
- [Installation](#installation)
- [Configuration](#configuration)
- [Running the Application](#running-the-application)
- [Accessing the Web Server via Balena Public URL](#accessing-the-web-server-via-balena-public-url)
- [Endpoints](#endpoints)
- [Shutdown](#shutdown)
- [License](#license)

## Requirements

- Raspberry Pi with BalenaOS
- Rainbow HAT
- Go 1.18+
- MQTT broker
- `.env` file with necessary configurations
- Docker

## Installation

1. **Clone the repository:**

    ```sh
    git clone https://github.com/yourusername/rainbowhat-sensor-reader.git
    cd rainbowhat-sensor-reader
    ```

2. **Ensure you have Go installed:**

    Follow the instructions on the [official Go website](https://golang.org/doc/install) to install Go.

3. **Install dependencies:**

    This project uses a `go.mod` file to manage dependencies. Ensure you are in the project directory and run:

    ```sh
    go mod tidy
    ```

4. **Create a `.env` file:**

    ```sh
    touch .env
    ```

    Populate the `.env` file with the following variables:

    ```env
    ID=RBH01
    MQTT_HOST=tls://[xyz......tyu].s1.eu.hivemq.cloud
    MQTT_PORT=8883
    MQTT_USER=test
    MQTT_PASS=test
    MQTT_CLIENT_ID=test
    MQTT_TOPIC=rainbowhat/weather
    PORT=:80
    REFRESH_INTERVAL=60
    ```

## Configuration

- `ID`: A unique ID for your sensor.
- `MQTT_HOST`: The host URL of your MQTT broker.
- `MQTT_PORT`: The port number for your MQTT broker.
- `MQTT_USER`: The username for your MQTT broker.
- `MQTT_PASS`: The password for your MQTT broker.
- `MQTT_CLIENT_ID`: Your MQTT client ID (can be left empty if not needed).
- `MQTT_TOPIC`: The MQTT topic to publish sensor data to.
- `PORT`: The port on which the web server will run (e.g., `:80`).
- `REFRESH_INTERVAL`: The interval (in seconds) at which sensor data is read and published.

## Running the Application

### Using Balena to Run on Raspberry Pi

1. **Install Balena CLI:**

    Follow the instructions on the [official Balena CLI documentation](https://www.balena.io/docs/reference/cli/#installation) to install the Balena CLI.

2. **Log in to Balena:**

    ```sh
    balena login
    ```

3. **Initialize the project:**

    ```sh
    balena push <your-app-name>
    ```

    Replace `<your-app-name>` with the name of your Balena application. This command will build and deploy the application to your Raspberry Pi.
    You need to have a Balena application set up with a Raspberry Pi device before running this command to deploy the application.

### Running Locally

1. **Build and run the application using Docker:**

    Make sure Docker is installed and running on your machine. Then, from the project directory, run:

    ```sh
    docker build -t rainbow-hat-sensor-service .
    docker run --env-file .env -p 8080:8080 rainbow-hat-sensor-service
    ```

## Accessing the Web Server via Balena Public URL

1. **Enable Public URL:**

    Go to the Balena dashboard, select your application, and enable the Public Device URL for your device. This will provide you with a URL through which you can access your web server.

2. **Configure Environment Variables:**

    Set the `PORT` environment variable to `80` in your Balena application configuration. This ensures that the web server listens on the correct port for the Balena public URL.

    ```env
    PORT=:80
    ```

3. **Access the Web Server:**

    Once the application is running, you can access the web server using the public URL provided by Balena.

## Endpoints

- **Root Endpoint (`/`)**: Serves an HTML page displaying the current sensor data.
- **Data Endpoint (`/data`)**: Serves the current sensor data in JSON format.

## Shutdown

The application is designed to handle graceful shutdown upon receiving an interrupt signal (e.g., `Ctrl+C` or a termination signal).

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
