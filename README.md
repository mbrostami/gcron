# gcron [In Development]
A go written tool to have a cronjob monitoring/reporting/stats. This will help you monitor outputs, status and timing and resource usage per cron.  
Generating report based on logs  
Stream cron outputs to remote servers (GCron server, Syslog server, logstash etc.)  

[gcron-server](https://github.com/mbrostami/gcron-server)
## TODO
- [ ] Clean code!!
- [ ] Test
- [ ] Support different log formats for write/stream purpose 
- [ ] Ignore errors (Run command even if connection is not established)
- [x] Run cron after given seconds
- [x] Implement gRPC
- [x] Send output to remote server (tcp/udp/unix)
- [x] Configurable tags (mem usage, cpu usage, systime, usertime, ...) (flag/config)
- [x] Trackable id for logs
- [x] Optional Regex status (Accept regex to change status of the cron to false or true)
  - [x] By default exitCode of the cron command will be used to detect if command was successful or failed
- [x] Stream logs over rpc
- [x] Remote mutex lock
- [x] Local mutex
- [x] Remote lock based on command
- [x] Remote lock timeout

## FIXME
- Delete local lock file

## Dev
Edit config.yml file and update log.path   
`go run main.go -c="echo 111 && sleep 1 && echo 222"`   
`go run main.go -c="git status"`  
```
  -c, --c string                  Command to execute (default "")
      --delay int                 Delay running command in seconds
      --lock.enable               Enable mutex lock
      --lock.remote               Use RPC for mutex lock
      --lock.name string          Mutex name
      --log.enable                Enable log
      --log.level string          Log level (default "warning")
      --out.hide.duration         Hide duration tag
      --out.hide.systime          Hide system time tag
      --out.hide.uid              Hide uid tag
      --out.hide.usertime         Hide user time tag
      --out.tags                  Output tags
      --override string           Override command status by regex match in output
      --server.rpc.enable         Enable RPC Server
      --server.rpc.host string    RPC Server host
      --server.rpc.port string    RPC Server port
      --server.tcp.enable         Enable TCP Server
      --server.tcp.host string    TCP Server host
      --server.tcp.port string    TCP Server port
      --server.udp.enable         Enable UDP Server
      --server.udp.host string    UDP Server host
      --server.udp.port string    UDP Server port
      --server.unix.enable        Enable Unix socket
      --server.unix.path string   UNIX socket path (default "/tmp/gcron-server.sock")
```
