# gcron  
Edit config.yml file and update log.path  
`go run *.go --exec="git status"`  
`go run main.go cron.go config.go --exec="echo 222 && sleep1 && echo 333"`  

## TODO
[] gcron server to store log streams  
[] gcron client support syslog RFC format  
[] gcron server configurable to use tcp/udp or socket connection  
