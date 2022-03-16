package vcago

import (
	"log"

	"github.com/nats-io/nats.go"
)

//Nats represents the config struct for Nats service.
type NatsDAO struct {
	host       string
	port       string
	skip       bool
	connection *nats.EncodedConn
}

//Nats used for Nats connection
var Nats = new(NatsDAO)

func (i *NatsDAO) Connect() {
	i.skip = Config.GetEnvBool("NATS_SKIP", "n", false)
	if i.skip {
		return
	}
	i.host = Config.GetEnvString("NATS_HOST", "w", "localhost")
	i.port = Config.GetEnvString("NATS_PORT", "w", "4222")
	natsUrl := "nats://" + i.host + ":" + i.port
	nc, err := nats.Connect(natsUrl)
	if err != nil {
		log.Fatal(err, " ", "NatsUrl: ", natsUrl)
	}
	i.connection, err = nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("nats successfully connected!")
	return
}

func (i *NatsDAO) Publish(message string, body interface{}) {
	if i.skip {
		return
	}
	i.connection.Publish(message, body)
}

func (i *NatsDAO) Subscribe(message string, catch interface{}) {
	if i.skip {
		return
	}
	_, err := i.connection.Subscribe(message, catch)
	if err != nil {
		//TODO: nats logging message
		log.Print(err)
	}

}
