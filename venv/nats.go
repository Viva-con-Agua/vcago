package venv

//Nats represents the config struct for Nats service.
type Nats struct {
	Host string
	Port string
}

func (i *Nats) Load() *Nats {
	var l LoadEnv
	i.Host, l = l.GetEnvString("NATS_HOST", "w", "localhost")
	i.Port, l = l.GetEnvString("NATS_PORT", "w", "4222")
	return i
}
