#!/bin/bash
curl http://laughing.plus:x/stop
GOOS=linux GOARCH=amd64 go build -o laughing main.go
sshpass -p '' scp laughing root@43.135.x.x:/usr/local/src
sshpass -p '' scp run.sh root@43.135.x.x:/usr/local/src
sshpass -p '' scp stop.sh root@43.135.x.x:/usr/local/src
echo "scp file to remote successfuly!"
rm -rf laughing
curl http://laughing.plus:x/run
echo "deploy laughing complete！"
