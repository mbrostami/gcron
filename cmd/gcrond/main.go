package main

import (
	"github.com/mbrostami/gcron/internal/app/gcrond/grpc"
	"github.com/mbrostami/gcron/internal/config"
	"github.com/mbrostami/gcron/internal/db"
	"github.com/mbrostami/gcron/internal/logs"
	"github.com/mbrostami/gcron/web"
)

var cfg config.GeneralConfig

func main() {
	cfg = config.GetConfig("gcrond", InitFlags())
	// initialize logs
	fd := logs.Initialize(cfg)
	defer fd.Close()

	var dbAdapter db.DB
	dataDir := cfg.GetKey("db.dataDir").(string)
	dbAdapter = db.NewLedis(dataDir)

	// Run in second thread
	go web.Listen(dbAdapter, cfg)

	// Run in main thread
	//taskCollection := dbAdapter.Get(1446109160, 0, 5)
	host := cfg.GetKey("server.rpc.host").(string)
	port := cfg.GetKey("server.rpc.port").(string)
	grpc.Run(host, port, dbAdapter)
}
