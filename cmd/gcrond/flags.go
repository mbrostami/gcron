package main

import "flag"

// InitFlags initialize flags
func InitFlags() *flag.FlagSet {
	flag.Bool("log.enable", false, "Enable file log")
	flag.String("log.path", "/var/log/gcron/gcron-server.log", "Log file path")
	flag.String("log.level", "", "Log level")
	flag.String("server.rpc.host", "", "RPC Server listening host")
	flag.String("server.rpc.port", "", "RPC Server listening port")
	flag.String("server.web.host", "", "Web server listening host")
	flag.String("server.web.port", "", "Web server listening port")
	flag.String("web.static", "", "Web serve path")
	return flag.CommandLine
}
