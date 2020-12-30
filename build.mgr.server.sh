#!/bin/sh

PATH=$PATH:$GOPATH/bin
rootdir=$(pwd)

echo ""
echo "---------打包mgrserver--------------------" 
echo ""


if [ ! -d out/mysql/scheme ] ; then 
	echo "请先执行 sh generate.db.sh "
	exit 1
fi  
if [ ! -d out/mysql/data ] ; then 
	echo "请先执行 sh generate.db.sh "
	exit 1
fi 

rm -f out/mysql/all.sql

echo "1. 使用go-bindata 整合Scheme文件"
go-bindata -o=mgrserver/mgrapi/modules/const/sqls/mysql/scheme/scheme.go -pkg=scheme out/mysql/scheme/*  > /dev/null 
if [ $? -ne 0 ]; then
	echo "go-bindata 整合SQL出错"
	exit 1
fi
 

echo "2. 使用go-bindata 整合Data文件"
go-bindata -o=mgrserver/mgrapi/modules/const/sqls/mysql/data/data.go -pkg=data out/mysql/data/*   > /dev/null
if [ $? -ne 0 ]; then
	echo "go-bindata 整合SQL出错"
	exit 1
fi


echo "3. 生成mgrserver"
cd $rootdir/mgrserver
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o "$rootdir/out/mgrserver/bin/mgrserver"
if [ $? -ne 0 ]; then
	echo "mgrserver 项目编译出错,请检查"
	exit 1
fi

echo ""
echo "---------打包mgrserver-success-------------------" 
echo ""