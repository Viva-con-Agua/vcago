package vcago

import (
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

// Nats represents the config struct for Nats service.
type NatsDAO struct {
	url        string
	skip       bool
	connection *nats.EncodedConn
}

// Nats used for Nats connection
var Nats = new(NatsDAO)

// Connect connects the NatsDAO to a nats-server
func (i *NatsDAO) Connect() {
	i.skip = Settings.Bool("NATS_SKIP", "n", false)
	if i.skip {
		return
	}
	i.url = Settings.String("NATS_URL", "w", "localhost")
	nc, err := nats.Connect(i.url)
	if err != nil {
		log.Fatal(err, " ", "NatsUrl: ", i.url)
	}
	i.connection, err = nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("nats successfully connected!")
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

func (i *NatsDAO) Request(message string, request interface{}, response interface{}) (err error) {
	if i.skip {
		return
	}
	err = i.connection.Request(message, request, response, 1000*time.Millisecond)
	return
}
