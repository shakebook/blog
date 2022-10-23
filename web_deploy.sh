#!/bin/bash
curl http://laughing.plus:9527/removeWeb &&
curl http://laughing.plus:9527/stop &&
sshpass -p '' scp -r web/styles root@43.135.79.57:/usr/local/src/web/styles &&
sshpass -p '' scp -r web/views root@43.135.79.57:/usr/local/src/web/views &&
curl http://laughing.plus:9527/run &&
echo "deploy web complete!"
