#!/bin/sh

PATH=$PATH:$GOPATH/bin
rootdir=$(dirname $(pwd)) 
pkg=$1
filename=login.static.zip

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
rm -f $filename
zip -r $filename  * > /dev/null
if [ $? -ne 0 ]; then
	echo "zip -r $filename * 出错"
	exit 1
fi

mkdir -p ${rootdir}/out

mv $filename ${rootdir}/out/
sleep 0.1

echo "3. 生成资源文件:loginserver/static.go"

if [ "$pkg" != "none" ] ; then 
	echo "3.1. 整合$filename文件"
	cd ${rootdir}/out
	go-bindata -o=${rootdir}/loginserver/static.go -pkg=main  $filename > /dev/null
	if [ $? -ne 0 ]; then
		echo "go-bindata 整合static出错"
		exit 1
	fi 
	sh $rootdir/build/empty.asset.sh ${rootdir}/loginserver 

fi
echo ""
echo "---------打包loginserver(loginweb)-success---------------" 
echo ""