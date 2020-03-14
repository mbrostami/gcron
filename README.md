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
`go run *.go --exec="git status"`  
`go run main.go cron.go config.go --exec="echo 222 && sleep1 && echo 333"`  

## TODO
- All Features! :D

# gcron client [In Development]
Client agent 