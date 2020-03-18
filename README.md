# gcron [In Development]
A go written tool to manage cron jobs. This will help you monitor outputs, status and timing and resource usage per cron.  
Generating report based on logs  
Stream cron outputs to remote servers (GCron server, Syslog server, logstash etc.)  

## TODO
- Support different log formats for write/stream purpose   
- Configurable tags (mem usage, cpu usage, systime, usertime, ...) (flag/config)
- Support single line log (convert newlines to specific character) (flag/config)
- Trackable id for logs
- Optional Regex status (Accept regex to change status of the cron to false or true)
  - By default exitCode of the cron command will be used to detect if command was successful or failed
- Report table - stdout (read given logs and display how many crons are logged, how many times they run and show time consuming crons ...)
- Dry run
- Command line search by tags
- Alert based on status 

# gcron server [In Development]
A go written tool to manage distributed cron jobs with centralized GUI. This will help you monitor crons if you have multiple servers with crons enabled.

## Features
- Centralized logging
- GUI interface
  - Search outputs by tags/server/text
  - Check crons status (exit code, running time, last run, ...)
  - Graphs 
  - Realtime monitoring
  - Resource usage per cron/server
- Accept syslog formats as input
- Auto detect input tags (specific syntax)
- Pipe input logs to remote syslog
- Pipe input logs to remote REST Api 

## Dev
Edit config.yml file and update log.path  
TCP Server `cd server && go run tcpserver.go cron.go`  
UDP Server `cd server && go run udpserver.go cron.go`  
UNIX Server `cd server && go run unixserver.go cron.go`  

Agent `cd agent && go run main.go config.go cron.go --exec="echo 111 && sleep 1 && echo 222"`  

## TODO
- All Features! :D



