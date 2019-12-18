#!/usr/bin/env bash

dep init
dep ensure -v


echo "Enter db password"
mysqladmin -u root -p create ESPGAME
mysql -u root ESPGAME < files/espdump.sql

echo "Use id=admin, password=admin"

echo "All Ready, Press any key to goto login page"
read -r line

open -a "Google Chrome" http://localhost:8080/login.html