# gcron [In Development]
A go written tool to manage cron jobs. This will help you monitor outputs, status and timing and resource usage per cron.  
Generating report based on logs  
Stream cron outputs to remote servers (GCron server, Syslog server, logstash etc.)  

[gcron-server](https://github.com/mbrostami/gcron-server)
## TODO
- Support different log formats for write/stream purpose 
- Run cron every given seconds for n times (e.g. every 10 seconds, in total 6 times)
- Configurable tags (mem usage, cpu usage, systime, usertime, ...) (flag/config)
- Support single line log (convert newlines to specific character) (flag/config)
- Trackable id for logs
- Optional Regex status (Accept regex to change status of the cron to false or true)
  - By default exitCode of the cron command will be used to detect if command was successful or failed
- Report table - stdout (read given logs and display how many crons are logged, how many times they run and show time consuming crons ...)
- Dry run
- Command line search by tags
- Alert based on status 

## Dev
Edit config.yml file and update log.path   
`go run main.go -exec="echo 111 && sleep 1 && echo 222"`  
`go run main.go -exec="git status"`  
```
  -exec string
        Command to execute
  -config string
        Config file path (default ".")
  -o-clean
        Clean output
  -o-notime
        Do not show datetime
```
