package main

import "flag"

// FlagCommand Command to execute
const FlagCommand string = "c"

// FlagHelp Show usage
const FlagHelp string = "help"

// FlagLockEnable Mutex lock
const FlagLockEnable string = "lock.enable"

// FlagLockName Custom lock name
const FlagLockName string = "lock.name"

// FlagOverride Override command status
const FlagOverride string = "override"

// FlagDelay Delay running command
const FlagDelay string = "delay"

// InitFlags initialize flags
func InitFlags() *flag.FlagSet {
	flag.String(FlagCommand, "", "Command to execute")
	flag.String(FlagLockName, "", "Mutex name")
	flag.String(FlagOverride, "", "Override command status by regex match in output")
	flag.String("server.rpc.host", "", "remote RPC host")
	flag.String("server.rpc.port", "", "remote RPC port")
	flag.Bool("server.rpc.enable", false, "enable RPC")
	flag.String("log.level", "info", "Log level")
	flag.Int("lock.timeout", 60, "Mutex timeout")
	flag.Int(FlagDelay, 0, "Delay running command in seconds")
	flag.Bool(FlagLockEnable, false, "Enable mutex lock")
	flag.Bool("lock.remote", false, "Use rpc mutex lock")
	flag.Bool("out.tags", false, "Output tags")
	flag.Bool("out.hide.systime", false, "Hide system time tag")
	flag.Bool("out.hide.usertime", false, "Hide user time tag")
	flag.Bool("out.hide.duration", false, "Hide duration tag")
	flag.Bool("out.hide.uid", false, "Hide uid tag")
	flag.Bool("log.enable", false, "Enable log")
	flag.Bool(FlagHelp, false, "Usage")
	return flag.CommandLine
}
