#!/bin/sh

#############################################
# ./builid.sh pkg
#############################################

#获取当前目录
rootdir=$(pwd)
PATH=$PATH:$GOPATH/bin

pkg=$1

rm -rf $rootdir/out 

echo ""
echo "---------打包-开始-------------------" 
echo ""

sh check.require.sh  
if [ $? -ne 0 ]; then
	echo "check.require出错"
	exit 1
fi
#------------------------------------" 
sh generate.db.sh   
if [ $? -ne 0 ]; then
	echo "generate.db出错"
	exit 1
fi
#------------------------------------" 
sh build.mgr.web.sh $1
if [ $? -ne 0 ]; then
	echo "build.mgr.web出错"
	exit 1
fi
#------------------------------------" 
sh build.mgr.server.sh $1
if [ $? -ne 0 ]; then
	echo "build.mgr.server出错"
	exit 1
fi
#------------------------------------" 
sh build.login.web.sh $1
if [ $? -ne 0 ]; then
	echo "build.login.web出错"
	exit 1
fi
#------------------------------------" 
sh build.login.server.sh $1
if [ $? -ne 0 ]; then
	echo "build.login.server出错"
	exit 1
fi

echo ""
echo "---------打包-success-------------------" 
echo "---------目录:${rootdir}/out"
echo ""
