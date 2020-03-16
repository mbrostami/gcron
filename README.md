# gcron [In Development]
A go written tool to manage distributed cron jobs with centralized GUI. This will help you to monitor crons if you have multiple servers with crons enabled.

## Features
- Centralized logging
- GUI interface
  - Search outputs by tags/server/text
  - Check crons status (exit code, running time, last run, ...)
  - Resource usage per cron/server
- Accept syslog formats as input
- Auto detect input tags (specific syntax)
- Pipe input logs to remote syslog
- Pipe input logs to remote REST Api 

## Dev
Edit config.yml file and update log.path  
TCP Server `cd tcpserver && go run server.go cron.go`  
Agent `cd agent && go run main.go config.go cron.go --exec="echo 111 && sleep 1 && echo 222"`  

## TODO
- All Features! :D

# gcron client [In Development]
Client agent 