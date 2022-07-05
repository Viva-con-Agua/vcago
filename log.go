//Package handles error logs via nats message broker
package vcago

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/nats-io/nats.go"
)

//Handles error logs via nats message broker

//Logger struct for handling the nats connection.
type LoggingHandler struct {
	service string
	nats    *nats.EncodedConn
	output  string
}

//NewLogger creates a new Logger.

var Logger = new(LoggingHandler)

func (i *LoggingHandler) New(method string, uri string, err error) *LogError {
	return &LogError{
		Time:   time.Now().Format(time.RFC3339Nano),
		Error:  err,
		Method: method,
		Uri:    uri,
	}
}

//Write is an IOWriter for handling the middleware Logger output.
func (i *LoggingHandler) Write(data []byte) (n int, err error) {
	n = len(data)
	logError := new(LogError)
	if err = json.Unmarshal(data, logError); err != nil {
		if data != nil {
			fmt.Print(string(data) + "\n")
		}
		return
	}
	//temp := new(interface{})
	//json.Unmarshal([]byte(logError.Error.(string)), temp)
	//logError.Error = temp
	if logError.Status/100 == 2 {
		return
	}
	if i.output == "nats" {
		i.Log(logError)
	} else if i.output == "pretty" {
		t, _ := json.MarshalIndent(logError, "", "    ")
		fmt.Print(string(t) + "\n")
	} else {
		t, _ := json.Marshal(logError)
		fmt.Print(string(t) + "\n")
	}
	return
}

//Log publish the LogError to nats route "logger.log".
func (i *LoggingHandler) Log(logError *LogError) {
	Nats.Publish("logger.log", logError)
}

//Config for echo middleware Logger. Use logger for handle Nats connection.
func (i *LoggingHandler) Init(service string) echo.MiddlewareFunc {
	i.service = service
	i.output = Config.GetEnvString("LOGGING_OUTPUT", "w", "strout")
	return middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `{"time":"${time_rfc3339_nano}","id":"${id}","remote_ip":"${remote_ip}",` +
			`"host":"${host}","method":"${method}","uri":"${uri}","user_agent":"${user_agent}",` +
			`"status":${status},"error":"${error}","latency":${latency},"latency_human":"${latency_human}"` +
			`,"bytes_in":${bytes_in},"bytes_out":${bytes_out}}`,
		Output: i,
	})
}

func NewLoggingHandler(service string) *LoggingHandler {
	return &LoggingHandler{
		service: service,
		output:  Config.GetEnvString("LOGGING_OUTPUT", "w", "strout"),
	}
}

//LogError represents the an LogError for handling via nats and store into mongo databases. The struct matches the Config Format string as json.
type LogError struct {
	ID           string      `json:"id" bson:"_id"`
	Time         string      `json:"time" bson:"time"`
	RemoteIP     string      `json:"remote_ip" bson:"remote_ip"`
	Host         string      `json:"host" bson:"host"`
	Method       string      `json:"method" bson:"method"`
	Uri          string      `json:"uri" bson:"uri"`
	UserAgent    string      `json:"user_agent" bson:"user_agent"`
	Status       int         `json:"status" bson:"status"`
	Error        interface{} `json:"error" bson:"error"`
	Latency      int64       `json:"latency" bson:"latency"`
	LatencyHuman string      `json:"latency_human" bson:"latency_human"`
}
