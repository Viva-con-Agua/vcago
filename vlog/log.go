//Package handles error logs via nats message broker
package vlog

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/Viva-con-Agua/vcago/venv"
	"github.com/Viva-con-Agua/vcago/verr"
	"github.com/Viva-con-Agua/vcago/vmod"
	"github.com/labstack/echo/v4/middleware"
	"github.com/nats-io/nats.go"
)

//Handles error logs via nats message broker

//Config for echo middleware Logger. Use logger for handle Nats connection.
func Config(logger *Logger) *middleware.LoggerConfig {
	return &middleware.LoggerConfig{
		Format: `{"time":"${time_rfc3339_nano}","id":"${id}","remote_ip":"${remote_ip}",` +
			`"host":"${host}","method":"${method}","uri":"${uri}","user_agent":"${user_agent}",` +
			`"status":${status},"error":"${error}","latency":${latency},"latency_human":"${latency_human}"` +
			`,"bytes_in":${bytes_in},"bytes_out":${bytes_out}}`,
		Output: logger,
	}
}

//Logger struct for handling the nats connection.
type Logger struct {
	Service string
	Nats    *nats.EncodedConn
	Type    string
}

//NewLogger creates a new Logger.

func NewLogger(service string) (logger *Logger) {
	logger = &Logger{
		Service: service,
		Type:    venv.Logger,
	}
	if venv.Logger == "NATS" {
		natsUrl := "nats://" + venv.NatsHost + ":" + venv.NatsPort
		nc, err := nats.Connect(natsUrl)
		if err != nil {
			log.Fatal(verr.ErrorWithColor, err, " ", "NatsUrl: ", natsUrl)
		}
		logger.Nats, err = nats.NewEncodedConn(nc, nats.JSON_ENCODER)
		if err != nil {
			log.Fatal(verr.ErrorWithColor, err)
		}
		log.Print("nats successfully connected!")
	}
	return
}

//Write is an IOWriter for handling the middleware Logger output.
func (i *Logger) Write(data []byte) (n int, err error) {
	n = len(data)
	logError := new(LogError)
	if err = json.Unmarshal(data, logError); err != nil {
		if data != nil {
			fmt.Print(string(data) + "\n")
		}
		return
	}
	if i.Type == "NATS" {
		i.Log(logError)
	} else {
		fmt.Print(string(data) + "\n")
	}
	return
}

//Log publish the LogError to nats route "logger.log".
func (i *Logger) Log(logError *LogError) {
	logError.Service = i.Service
	i.Nats.Publish("logger.log", logError)
}

//LogError represents the an LogError for handling via nats and store into mongo databases. The struct matches the Config Format string as json.
type LogError struct {
	ID           string        `json:"id" bson:"_id"`
	Service      string        `json:"service" bson:"service"`
	Time         string        `json:"time" bson:"time"`
	RemoteIP     string        `json:"remote_ip" bson:"remote_io"`
	Host         string        `json:"host" bson:"host"`
	Method       string        `json:"method" bson:"method"`
	Uri          string        `json:"uri" bson:"uri"`
	UserAgent    string        `json:"user_agent" bson:"user_agent"`
	Status       int           `json:"status" bson:"status"`
	Error        string        `json:"error" bson:"error"`
	Latency      int64         `json:"latency" bson:"latency"`
	LatencyHuman string        `json:"latency_human" bson:"latency_human"`
	ByteIn       string        `json:"byte_in" bson:"byte_in"`
	ByteOut      string        `json:"byte_out" bson:"byte_out"`
	Modified     vmod.Modified `json:"modified" bson:"modified"`
}
