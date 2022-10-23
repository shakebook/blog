#!/bin/bash
curl http://laughing.plus:x/removeWeb &&
curl http://laughing.plus:x/stop &&
sshpass -p '' scp -r web/styles root@43.135.xx.xx:/usr/local/src/web/styles &&
sshpass -p '' scp -r web/views root@43.135.xx.xx:/usr/local/src/web/views &&
curl http://laughing.plus:x/run &&
echo "deploy web complete!"
