# gcron [In Development]
A go written tool to manage cron jobs. This will help you monitor outputs, status and timing and resource usage per cron.  
Generating report based on logs  
Stream cron outputs to remote servers (GCron server, Syslog server, logstash etc.)  

[gcron-server](https://github.com/mbrostami/gcron-server)
## TODO
- Support different log formats for write/stream purpose 
- ~~Mutex lock~~
- Run cron every given seconds for n times (e.g. every 10 seconds, in total 6 times)
- ~~Configurable tags (mem usage, cpu usage, systime, usertime, ...) (flag/config)~~
- ~~Trackable id for logs~~
- ~~Optional Regex status (Accept regex to change status of the cron to false or true)~~
  - ~~By default exitCode of the cron command will be used to detect if command was successful or failed~~
- Report table - stdout (read given logs and display how many crons are logged, how many times they run and show time consuming crons ...)
- Dry run
- Command line search by tags
- Alert based on status 

## Dev
Edit config.yml file and update log.path   
`go run main.go -exec="echo 111 && sleep 1 && echo 222"`  
`go run main.go -exec="git status"`  
```
      --exec string               Command to execute (default "echo")
      --lock.enable               Enable mutex lock
      --lock.name string          Mutex name
      --out.clean                 Only command output
      --out.hide.duration         Hide duration tag (default false)
      --out.hide.systime          Hide system time tag (default false)
      --out.hide.uid              Hide uid tag (default false)
      --out.hide.usertime         Hide user time tag (default false)
      --out.tags                  Output tags (default false)
      --override string           Override command status by regex match in output (default "")
      --server.tcp.host string    TCP Server host
      --server.tcp.port string    TCP Server port
      --server.udp.host string    UDP Server host
      --server.udp.port string    UDP Server port
      --server.unix.path string   UNIX socket path (default "/tmp/gcron-server.sock")
```
