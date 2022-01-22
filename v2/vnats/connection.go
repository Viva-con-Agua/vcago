package vnats

import (
	"log"

	"github.com/Viva-con-Agua/vcago/verr"
	"github.com/nats-io/nats.go"
)

type Connection struct {
	host       string
	port       string
	connection *natsEncodedConn
}

var Nats = new(Connection)

func Connect() {
	Nats.host = l.GetEnvString("NATS_HOST", "w", "localhost")
	Nats.port = l.GetEnvString("NATS_PORT", "w", "4222")
	uri := "nats://" + i.host + ":" + i.port
	connection, err := nats.Connect(url)
	if err != nil {
		log.Fatal(verr.ErrorWithColor, err, " ", "NatsUrl: ", uri)
	}
	Nats.connection, err = nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	if err != nil {
		log.Fatal(verr.ErrorWithColor, err)
	}
	log.Print("nats successfully connected!")
}
