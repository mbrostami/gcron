package main

import (
	"flag"
	"os"
	"strconv"

	localmutex "github.com/mbrostami/gcron/cmd/gcron/mutex"
	gcronout "github.com/mbrostami/gcron/cmd/gcron/out"
	"github.com/mbrostami/gcron/internal/config"
	grpc "github.com/mbrostami/gcron/internal/grpc"
	"github.com/mbrostami/gcron/internal/logs"
	"github.com/mbrostami/gcron/internal/mutex"
	"github.com/mbrostami/gcron/internal/task"
	"github.com/mbrostami/gcron/pkg/bash"
	"github.com/mbrostami/gcron/pkg/validators"
	log "github.com/sirupsen/logrus"
)

var cfg config.GeneralConfig
var mtx mutex.Interface

func main() {
	cfg = config.GetConfig(InitFlags())
	// initialize logs
	fd := logs.Initialize(cfg)
	defer fd.Close()

	if cfg.GetKey(FlagHelp).(bool) {
		flag.Usage()
		os.Exit(1)
	}
	crontask := task.Task{
		Command:   cfg.GetKey(FlagCommand).(string),
		FLock:     cfg.GetKey(FlagLockEnable).(bool),
		FLockName: cfg.GetKey(FlagLockName).(string),
		FOverride: cfg.GetKey(FlagOverride).(string),
		FDelay:    cfg.GetKey(FlagDelay).(int),
	}
	if _, err := crontask.Validate(); err != nil {
		log.Fatal(err)
	}
	crontask.SetBasics()

	grpcHandler := &gcronout.GrpcHandler{}

	grpcEnable := cfg.GetKey("server.rpc.enable").(bool)
	remoteLock := cfg.GetKey("lock.remote").(bool)
	lockTimeout := cfg.GetKey("lock.timeout").(int)
	lockName := strconv.FormatUint(uint64(crontask.UID), 10)
	if crontask.FLockName != "" {
		lockName = crontask.FLockName
	}
	if grpcEnable {
		var err error
		grpcHandler, err = gcronout.NewHandler(cfg)
		if err != nil {
			log.Fatalf("RPC server is not available %s", err.Error())
		}
		//defer grpcHandler.Close()
		if remoteLock { // remote lock can only be used with rpc
			mtx = &localmutex.GrpcMutex{Server: grpcHandler}
		}
	} else if remoteLock {
		log.Fatalf("Remote lock requires enabled grpc server (see server.rpc.enable)")
	}
	if crontask.FLock && !remoteLock { // override grpcMutex if is set
		mtx = &localmutex.FileMutex{}
	}

	if crontask.FLock {
		locked, err := mtx.Lock(lockName, lockTimeout)
		if !locked {
			log.Fatalf("Task is already running... %s", err.Error())
		}
	}

	outputCh := make(chan []byte)
	cmd, err := bash.MakeCommand(crontask.Command, crontask.FDelay, outputCh)
	if err != nil {
		log.Fatalf("Command failed: %s", err.Error())
	}

	var stream grpc.Gcron_StartLogClient
	if grpcEnable {
		stream, err = grpcHandler.StartLogStream()
		if err != nil {
			log.Fatalf("Stream failed %s", err.Error())
		}
	}

	// Start execution
	err = cmd.Execute()
	if err != nil {
		log.Fatalf("Command execution failed: %s", err.Error())
	}

	crontask.StartTime = cmd.CmdInfo.StartTime
	crontask.Success = cmd.CmdInfo.Status
	crontask.Pid = cmd.CmdInfo.Pid

	var statusByRegex = false
	for output := range outputCh {
		log.Debugf("%s", string(output))
		if crontask.FOverride != "" {
			statusByRegex = statusByRegex || validators.NewRegex(crontask.FOverride).Validate(string(output))
		}
		// Stream output
		if grpcEnable {
			stream.Send(grpcHandler.GetLogEntry(crontask.GUID, output))
		}
		output = append(output, []byte("\n")...) // Add new line
		crontask.Output = append(crontask.Output, output...)
	}
	// Execution is finished now
	if grpcEnable {
		stream.CloseAndRecv()
	}
	if crontask.FLock {
		mtx.Release(lockName)
	}
	crontask.Success = cmd.CmdInfo.Status
	if crontask.FOverride != "" {
		crontask.Success = statusByRegex
	}

	crontask.Parent = cmd.CmdInfo.Parent
	crontask.Username = cmd.CmdInfo.Username
	crontask.SystemTime = cmd.CmdInfo.SystemTime
	crontask.UserTime = cmd.CmdInfo.UserTime
	crontask.ExitCode = cmd.CmdInfo.ExitCode
	crontask.EndTime = cmd.CmdInfo.EndTime

	if grpcEnable {
		grpcHandler.Done(crontask)
	}
	// Log tags
	if cfg.GetKey("out.tags").(bool) {
		// var customOutput string
		fields := log.Fields{}
		if !cfg.GetKey("out.hide.uid").(bool) {
			fields["uid"] = crontask.UID
		}
		if !cfg.GetKey("out.hide.systime").(bool) {
			fields["systime"] = crontask.SystemTime.Seconds()
		}
		if !cfg.GetKey("out.hide.usertime").(bool) {
			fields["usertime"] = crontask.UserTime.Seconds()
		}
		if !cfg.GetKey("out.hide.duration").(bool) {
			fields["duration"] = crontask.EndTime.Sub(crontask.StartTime).Seconds()
		}
		fields["status"] = crontask.Success
		log.WithFields(fields).Info("tags")
	}
}
