#!/bin/sh

PATH=$PATH:$GOPATH/bin
rootdir=$(dirname $(pwd)) 
pkg=$1
filename=mgr.static.zip

echo ""
echo "---------打包mgrserver(mgrweb)-------------------" 
echo ""

cd $rootdir/mgrserver/mgrweb

echo "1. 打包项目：npm run build"
npm run build  > /dev/null 
if [ $? -ne 0 ]; then
	echo "npm run build 出错"
	exit 1
fi

echo "2. 压缩：dist/static"
cd dist/static
rm -f $filename
zip -r  $filename * > /dev/null
if [ $? -ne 0 ]; then
	echo "tar -zcvf $filename dist/static/* 出错"
	exit 1
fi

mkdir -p ${rootdir}/out

mv $filename ${rootdir}/out/

sleep 0.1
echo "3. 生成资源文件:loginserver/loginapi/web/static.go" 
if [ "$pkg" != "none" ] ; then 
	echo "3.1. 整合static.tar.gz文件"
	sleep 0.1
	cd $rootdir/out
	go-bindata -o=${rootdir}/mgrserver/static.go -pkg=main $filename > /dev/null
	if [ $? -ne 0 ]; then
		echo "go-bindata 整合static出错"
		exit 1
	fi
	sh $rootdir/build/empty.asset.sh ${rootdir}/mgrserver 
fi

echo ""
echo "---------打包mgrserver(mgrweb)-success------------------" 
echo ""