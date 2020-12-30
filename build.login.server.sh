#!/bin/sh

rootdir=$(pwd)

echo ""
echo "---------打包loginserver--------------------" 
echo "" 

echo "1. 生成loginserver"
cd $rootdir/loginserver
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o "$rootdir/out/loginserver/bin/loginserver"
if [ $? -ne 0 ]; then
	echo "loginserver 项目编译出错,请检查"
	exit 1
fi

echo ""
echo "---------打包loginserver--success------------------" 
echo ""