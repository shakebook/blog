#!/bin/bash
curl http://laughing.plus:9527/removeWeb &&
curl http://laughing.plus:9527/stop &&
sshpass -p '@YangZhang^&#i>(Le*uUo09*%$C.,' scp -r web/styles root@43.135.79.57:/usr/local/src/web/styles &&
sshpass -p '@YangZhang^&#i>(Le*uUo09*%$C.,' scp -r web/views root@43.135.79.57:/usr/local/src/web/views &&
curl http://laughing.plus:9527/run &&
echo "deploy web complete!"
