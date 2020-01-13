#!/bin/sh

netstat -vanp tcp | grep 4104
sudo lsof -i tcp:4104
echo "kill -9 "

exit 0

