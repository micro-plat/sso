#!/bin/sh

PATH=$PATH:$GOPATH/bin

rootdir=$(pwd)

echo "" 
echo "---------打包检查-----------------"
echo "" 
echo "1. 检查 phorcys "
which phorcys > /dev/null
if [ $? -ne 0 ]; then
	echo "1.a. phorcys 不存在，执行安装"
	
	dir=$GOPATH/src/gitlab.100bm.cn/devtools/phorcys/phorcys
	if [ ! -d  $dir ] ; then
		cd $GOPATH/src/
		mkdir -p gitlab.100bm.cn/devtools/phorcys
		cd gitlab.100bm.cn/devtools/phorcys

		echo "1.b. 下载phorcys"
		git clone https://gitlab.100bm.cn/devtools/phorcys/phorcys.git
		
		git checkout dev
	fi
	cd $dir
	git pull 
	echo "1.c. 编译phorcys"
	go install 
	echo "" 

fi 

cd $rootdir
echo "2. 检查 go-bindata "

which go-bindata > /dev/null
if [ $? -ne 0 ]; then
	echo "2.a. go-bindata 不存在，执行安装"
	dir=$GOPATH/src/github.com/go-bindata/go-bindata
	if [ ! -d  $dir ] ; then
		cd $GOPATH/src
		mkdir -p github.com/go-bindata
		cd github.com/go-bindata
		
		echo "2.b. 下载go-bindata"
		git clone https://github.com/go-bindata/go-bindata.git
	fi
	
	cd $dir
	git pull 
	cd go-bindata
 	echo "2.c. 编译go-bindata"
	go install 
fi 

echo "" 
echo "---------打包检查-success---------"
echo "" 