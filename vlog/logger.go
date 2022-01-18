package vlog

import (
	"github.com/Viva-con-Agua/vcago/vmod"
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
	Modified     vmod.Modified `json:"modified" bson:"modified"`
}

//Logger struct for handling the nats connection.
type LoggingHandler struct {
	service string
	nats    *nats.EncodedConn
	output  string
}

//Logger can be used for log in echo projects
var Logger = new(LoggingHandler)

func Init() {
	Logger.service = Config.GetEnvString("SERVICE_NAME", "w", "default")
	Logger.output = Config.GetEnvString("LOGGING_OUTPUT", "w", "strout")
}
