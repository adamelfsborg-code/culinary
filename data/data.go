package data

import (
	"github.com/adamelfsborg-code/food/culinary/config"
	"github.com/go-pg/pg/v10"
	"github.com/nats-io/nats.go"
)

type DataConn struct {
	Env  config.Environments
	DB   pg.DB
	Nats *nats.Conn
	JS   nats.JetStreamContext
}

var Data *DataConn
