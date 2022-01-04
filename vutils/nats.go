package vutils

import (
	"log"

	"github.com/nats-io/nats.go"
)

//Nats represents the initial struct for an nats connection.
type Nats struct {
	Host string
	Port string
}

//LoadEnv loads the Host and Port From .env file.
//Host can be set via NATS_HOST
//Port can be set via NATS_PORT
func (i *Nats) LoadEnv() *Nats {
	var l LoadEnv
	i.Host, l = l.GetEnvString("NATS_HOST", "w", "localhost")
	i.Port, l = l.GetEnvString("NATS_PORT", "w", "4222")
	return i
}

//Connect connects nats client to server. The client is using the Nats.Host and Nats.Port parameters.
func (i *Nats) Connect() (r *nats.EncodedConn) {
	log.Print("nats connecting ... ")
	natsUrl := "nats://" + i.Host + ":" + i.Port
	nc, err := nats.Connect(natsUrl)
	if err != nil {
		log.Fatal(err)
	}
	r, err = nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("nats successfully connected!")
	return
}
