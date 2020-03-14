# gcron client [In Development]
A go written tool to manage distributed cron jobs with centralized GUI. This will help you if you have multiple servers with crons enabled.

## Features
- Monitor resource usage by each cron

## Usage
Edit config.yml file and update log.path  
`go run *.go --exec="git status"`  
`go run main.go cron.go config.go --exec="echo 222 && sleep1 && echo 333"`  

## TODO
- gcron server to store log streams  
- gcron client support syslog format (2 RFC)   
- gcron server configurable to use tcp/udp or socket connection  

# gcron server [In Development]

## Features
- Centralized logging
- GUI interface
  - Search outputs by tags/server/text
  - Check crons status (exit code, running time, last run, ...)