#!/bin/sh

PATH=$PATH:$GOPATH/bin
rootdir=$(dirname $(pwd)) 
pkg=$1

echo ""
echo "---------打包loginserver(loginweb)--------------------" 
echo ""



echo "0. 检查 go-bindata " 
which go-bindata > /dev/null
if [ $? -ne 0 ]; then
 	echo "go-bindata 未安装"
	echo "请到https://github.com/go-bindata/go-bindata.git下载安装"
	exit 1	  
fi 


cd $rootdir/loginserver/loginweb

echo "1. 打包项目：npm run build"
npm run build > /dev/null
if [ $? -ne 0 ]; then
	echo "npm run build 出错"
	exit 1
fi

echo "2. 压缩：loginweb/dist/static"
cd dist/static
rm -f login.static.tar.gz
tar -zcvf login.static.tar.gz * > /dev/null
if [ $? -ne 0 ]; then
	echo "tar -zcvf login.static.tar.gz dist/static/* 出错"
	exit 1
fi

mkdir -p ${rootdir}/out

mv login.static.tar.gz ${rootdir}/out/
sleep 0.1

echo "3. 生成资源文件:loginserver/loginapi/web/static.go"

if [ "$pkg" = "none" ] ; then 
	echo "3.1. 生成空文件文件"
	cd $rootdir
	sh $rootdir/scripts/empty.asset.sh ${rootdir}/loginserver/loginapi/web 
else
	echo "3.1. 整合static.tar.gz文件"
	cd ${rootdir}/out
	go-bindata -o=${rootdir}/loginserver/loginapi/web/static.go -pkg=web  login.static.tar.gz > /dev/null
	if [ $? -ne 0 ]; then
		echo "go-bindata 整合static出错"
		exit 1
	fi 
	echo "3.2. 请重新执行sh build.sh生成loginserver二进制文件"

fi
echo ""
echo "---------打包loginserver(loginweb)-success---------------" 
echo ""