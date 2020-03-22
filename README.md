# gcron [In Development]
A go written tool to have a cronjob monitoring/reporting/stats. This will help you monitor outputs, status and timing and resource usage per cron.  
Generating report based on logs  
Stream cron outputs to remote servers (GCron server, Syslog server, logstash etc.)  

[gcron-server](https://github.com/mbrostami/gcron-server)
## TODO
- [ ] Support different log formats for write/stream purpose 
- [x] Mutex lock
- [ ] Run cron every given seconds for n times (e.g. every 10 seconds, in total 6 times)
- [x] Configurable tags (mem usage, cpu usage, systime, usertime, ...) (flag/config)
- [x] Trackable id for logs
- [x] Optional Regex status (Accept regex to change status of the cron to false or true)
  - [x] By default exitCode of the cron command will be used to detect if command was successful or failed
- [ ] Report table - stdout (read given logs and display how many crons are logged, how many times they run and show time consuming crons ...)
- [ ] Command line search by tags
- [ ] Alert based on status 

## FIXME
- Delete lock file

## Dev
Edit config.yml file and update log.path   
`go run main.go report ...`    
`go run main.go exec -c="echo 111 && sleep 1 && echo 222"`   
`go run main.go exec -c="git status"`  
```
      --c                         Command to execute (default "echo")
      --lock.enable               Enable mutex lock (prevent running same process twice)
      --lock.name                 Mutex name (if it's not provided, uid will be used)
      --out.clean                 Only command output
      --out.hide.duration         Hide duration tag (default false)
      --out.hide.systime          Hide system time tag (default false)
      --out.hide.uid              Hide uid tag (default false)
      --out.hide.usertime         Hide user time tag (default false)
      --out.tags                  Output tags (default false)
      --override                  Override command status by regex match in output (default "")
      --server.tcp.host           TCP Server host
      --server.tcp.port           TCP Server port
      --server.udp.host           UDP Server host
      --server.udp.port           UDP Server port
      --server.unix.path          UNIX socket path (default "/tmp/gcron-server.sock")
```
