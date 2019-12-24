#!/usr/bin/env bash

echo "Setting GOPATH..."
export GOPATH=$(go env GOPATH)
apppath="$GOPATH/src/github.com/ESPGAME"

echo "Creating workspace..."
mkdir -p $apppath

tput setaf 1;echo "Copying folder to workspace..."
echo "use current path from nowon"
pwd
cp -R . $apppath
cd $apppath

tput sgr0
echo "Initiating dependencies....	"
dep init
dep ensure -v


echo "Enter db password"
mysqladmin -u root -p create ESPGAME
mysql -u root ESPGAME < files/espdump.sql

echo "Use id=admin, password=admin"

echo "Starting Server..."
if go run src/app.go; then
	echo "All Ready, Press any key to goto login page"
	read -r line
	open -a "Google Chrome" http://localhost:8080/login.html
else
	echo "Something went wrong"
fi

tput setaf 1
echo "Workspace moved to new location,"
pwd
