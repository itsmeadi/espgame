#!/usr/bin/env bash

dep ensure -v
mysql -u root < files/espdump.sql

echo "use id=admin, password=admin"

open -a "Google Chrome" http://l:8080/login.html