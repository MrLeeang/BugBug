#!/bin/bash

chmod +x /www/wwwroot/BugBug/BugBug

ps -ef |grep BugBug |grep -v grep |awk '{print $2}' |xargs kill -9

nohup /www/wwwroot/BugBug/BugBug >/var/log/api.log 2>&1 &
