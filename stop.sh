#!/bin/bash
pids="`ps -ef |grep laughing |grep -v -e grep -e killkeys |awk '{print $2}'`"
echo "kill pid:"$pids
kill -9 $pids