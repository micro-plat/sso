#!/bin/sh

#############################################
# sh build.sh 
#############################################

#获取当前目录
rootdir=$(dirname $(pwd)) 

PATH=$PATH:$GOPATH/bin

pkg=$1
 
echo ""
echo "---------打包-start-----------------" 
echo ""
  
#------------------------------------" 
echo "1. loginserver 项目编译" 
cd $rootdir/loginserver
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o "$rootdir/out/loginserver"
if [ $? -ne 0 ]; then
	echo "loginserver 项目编译出错,请检查"
	exit 1
fi

#------------------------------------" 
echo "2. mgrserver 项目编译"
cd $rootdir/mgrserver
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o "$rootdir/out/mgrserver"
if [ $? -ne 0 ]; then
	echo "mgrserver 项目编译出错,请检查"
	exit 1
fi

echo ""
echo "---------打包-success----------------" 
echo "---------目录:${rootdir}/out"
echo ""
