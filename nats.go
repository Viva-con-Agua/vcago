package vcago

import (
	"log"

	"github.com/nats-io/nats.go"
)

//Nats represents the config struct for Nats service.
type NatsDAO struct {
	host       string
	port       string
	connection *nats.EncodedConn
}

var Nats = new(NatsDAO)

func (i *NatsDAO) LoadEnv() *NatsDAO {
	var l LoadEnv
	i.host = l.GetEnvString("NATS_HOST", "w", "localhost")
	i.port = l.GetEnvString("NATS_PORT", "w", "4222")
	return i
}

func (i *NatsDAO) Connect() (r *NatsDAO) {
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
	return i
}

func (i *NatsDAO) Publish(message string, body interface{}) {
	i.connection.Publish(message, body)
}

func (i *NatsDAO) Subscribe(message string, catch interface{}) {

	_, err := i.connection.Subscribe(message, catch)
	if err != nil {
		//TODO: nats logging message
		log.Print(err)
	}

}
