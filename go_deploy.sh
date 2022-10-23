#!/bin/bash
curl http://laughing.plus:9527/stop
GOOS=linux GOARCH=amd64 go build -o laughing main.go
sshpass -p '@YangZhang^&#i>(Le*uUo09*%$C.,' scp laughing root@43.135.79.57:/usr/local/src
sshpass -p '@YangZhang^&#i>(Le*uUo09*%$C.,' scp run.sh root@43.135.79.57:/usr/local/src
sshpass -p '@YangZhang^&#i>(Le*uUo09*%$C.,' scp stop.sh root@43.135.79.57:/usr/local/src
echo "scp file to remote successfuly!"
rm -rf laughing
curl http://laughing.plus:9527/run
echo "deploy laughing complete！"