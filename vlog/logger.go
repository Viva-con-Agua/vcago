package vlog

import (
	"encoding/json"
	"fmt"

	"github.com/Viva-con-Agua/vcago/vmdb"
	"github.com/Viva-con-Agua/vcago/vutils"
	"github.com/labstack/echo/v4/middleware"
	"github.com/nats-io/nats.go"
)

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
	Error        interface{}   `json:"error" bson:"error"`
	Latency      int64         `json:"latency" bson:"latency"`
	LatencyHuman string        `json:"latency_human" bson:"latency_human"`
	ByteIn       string        `json:"byte_in" bson:"byte_in"`
	ByteOut      string        `json:"byte_out" bson:"byte_out"`
	Modified     vmdb.Modified `json:"modified" bson:"modified"`
}

//Logger struct for handling the nats connection.
type LoggingHandler struct {
	service string
	nats    *nats.EncodedConn
	output  string
}

//Logger can be used for log in echo projects
var Logger = new(LoggingHandler).Init()

func (i *LoggingHandler) Init() *LoggingHandler {
	i.service = vutils.Config.GetEnvString("SERVICE_NAME", "w", "default")
	i.output = vutils.Config.GetEnvString("LOGGING_OUTPUT", "w", "strout")
	return i
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
	temp := new(interface{})
	json.Unmarshal([]byte(logError.Error.(string)), temp)
	logError.Error = temp
	if logError.Status/100 == 2 {
		return
	}
	if i.output == "nats" {
		i.Log(logError)
	} else {
		t, _ := json.Marshal(logError)
		fmt.Print(string(t) + "\n")
	}
	return
}

//Log publish the LogError to nats route "logger.log".
func (i *LoggingHandler) Log(logError *LogError) {
	logError.Service = i.service
	i.nats.Publish("logger.log", logError)
}

//Config for echo middleware Logger. Use logger for handle Nats connection.
func (i *LoggingHandler) Config() *middleware.LoggerConfig {
	return &middleware.LoggerConfig{
		Format: `{"time":"${time_rfc3339_nano}","id":"${id}","remote_ip":"${remote_ip}",` +
			`"host":"${host}","method":"${method}","uri":"${uri}","user_agent":"${user_agent}",` +
			`"status":${status},"error":"${error}","latency":${latency},"latency_human":"${latency_human}"` +
			`,"bytes_in":${bytes_in},"bytes_out":${bytes_out}}`,
		Output: i,
	}
}
