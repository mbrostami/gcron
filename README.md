[![Build Status](https://travis-ci.com/mbrostami/gcron.svg?branch=master)](https://travis-ci.com/mbrostami/gcron)
# gcron [In Development] 
A go written tool to have a cronjob monitoring/reporting/stats. This will help you monitor outputs, status and timing and resource usage per cron.  
Generating report based on logs  
Stream cron outputs to remote servers (GCron server, Syslog server, logstash etc.)  

[gcron-server](https://github.com/mbrostami/gcron-server)

## Usage
```
// Delay running command  
* * * * * gcron -c="echo HelloWorld" --delay=5  
* * * * * gcron -c="echo HelloWorld" --delay=10  
* * * * * gcron -c="echo HelloWorld" --delay=15

// Delay running command but avoid duplicate running 
* * * * * gcron -c="echo HelloWorld" --delay=10 --lock.enable
* * * * * gcron -c="echo HelloWorld" --delay=20 --lock.enable
* * * * * gcron -c="echo HelloWorld" --delay=30 --lock.enable

// Mutex lock to prevent overlap same command
* * * * * gcron -c="sleep 61 && echo HelloWorld" --lock.enable

// Enable logging and log level
gcron -c="echo HelloWorld" --log.enable --log.level=trace

// Display tags (systime, usertime, duration, etc)
gcron -c="echo HelloWorld" --log.level=info --out.tags

// Override command status which is stored in tags.status [case sensitive]
gcron -c="echo HelloWorld" --log.level=info --out.tags --override=".*World$" 
-- [INFO] [status:true]
gcron -c="echo HelloWorld" --log.level=info --out.tags --override=".*Worl$" 
-- [INFO] [status:false]

// Remote mutex lock to prevent overlap same command in multiple servers [gcron-server required]
* * * * * gcron -c="echo Server1HelloWorld" --lock.enable --lock.remote


// Using combination of gcrons
* * * * * gcron -c="echo SendLogsToGcronServer" --lock.enable --lock.remote --server.rpc.enabled 
* * * * * gcron -c="echo KeepLogsOnlyInLocal" --log.enable --server.rpc.enabled=0
* * * * * gcron -c="echo TraceOutputs" --log.enable --log.level=trace --server.rpc.enabled=0
* * * * * gcron -c="echo RunWithDelay" --delay=5

```  

## TODO
- [ ] Clean code!!
- [ ] Test
- [ ] Support different log formats for write/stream purpose 
- [ ] Ignore errors (Run command even if connection is not established)
- [ ] Using workers pool 
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

## Production
Download binary file  
Create a config file   
`mkdir /etc/gcron && touch /etc/gcron/config.yml`  
Links  
```ln -s `pwd`/gcron /usr/local/bin/gcron```
## Development
Edit config.yml file and update log.path   
`go run main.go -c="echo 111 && sleep 1 && echo 222"`   
`go run main.go -c="git status"`  
```
  -c, --c string                 Command to execute
      --delay int                Delay running command in seconds
      --lock.enable              Enable mutex lock
      --lock.name string         Mutex name
      --lock.remote              Use rpc mutex lock
      --lock.timeout int         Mutex timeout (default 60)
      --log.enable               Enable log
      --log.level string         Log level (default "warning")
      --out.hide.duration        Hide duration tag
      --out.hide.systime         Hide system time tag
      --out.hide.uid             Hide uid tag
      --out.hide.usertime        Hide user time tag
      --out.tags                 Output tags
      --override pattern         Override command status by regex match in output
      --server.rpc.host string   RPC Server host
      --server.rpc.port string   RPC Server port
```
