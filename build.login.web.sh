#!/bin/sh

PATH=$PATH:$GOPATH/bin
rootdir=$(pwd)
pkg=$1

echo ""
echo "---------打包loginserver(loginweb)--------------------" 
echo ""
 
cd $rootdir/loginserver/loginweb

echo "1. 打包项目：npm run build"
npm run build > /dev/null 2>&1
if [ $? -ne 0 ]; then
	echo "npm run build 出错"
	exit 1
fi

echo "2. 压缩：loginweb/dist/static"
cd dist/static
rm -f static.tar.gz
tar -zcvf static.tar.gz * > /dev/null
if [ $? -ne 0 ]; then
	echo "tar -zcvf static.tar.gz dist/static/* 出错"
	exit 1
fi

mkdir -p ${rootdir}/out/loginserver/bin

mv static.tar.gz ${rootdir}/out/loginserver/bin
sleep 0.1

echo "3. 生成资源文件:loginserver/loginapi/web/static.go"

if [ "$pkg" = "pkg" ] ; then 
	echo "a. 整合static.tar.gz文件"
	cd ${rootdir}/out/loginserver/bin
	go-bindata -o=${rootdir}/loginserver/loginapi/web/static.go -pkg=web  static.tar.gz > /dev/null
	if [ $? -ne 0 ]; then
		echo "go-bindata 整合static出错"
		exit 1
	fi
else
	echo "b. 生成空文件文件"
	cd $rootdir
	sh empty.asset.sh ${rootdir}/loginserver/loginapi/web
fi
echo ""
echo "---------打包loginserver(loginweb)-success---------------" 
echo ""