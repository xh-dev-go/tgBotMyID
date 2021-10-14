# Installation
```shell
go get github.com/xh-dev-go/tgBotMyID
go install github.com/xh-dev-go/tgBotMyID
```
# Demo
## Version
```shell
tgBotMyID -version
```

## Check User TG ID
Before you can test the tg id, you need to have existing api token for tg bot.
Run below script to start up ip check service. 
```shell
tgBotMyID --token {token}
```
Send message the bot.
The bot will return message including the name, tg id and is bot of the caller.

