#!/bin/sh

PATH=$PATH:$GOPATH/bin
rootdir=$(dirname $(pwd)) 
pkg=$1
filename=login.static.zip

echo ""
echo "---------打包loginserver(loginweb)--------------------" 
echo ""




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
 

rm -f ${rootdir}/loginserver/static.go
rm -f ${rootdir}/loginserver/web.go

echo "3. 生成资源文件:loginserver/web.go"
if [ "$pkg" != "none" ] ; then 
	echo "3.1. 整合$filename文件"
	cp  ${rootdir}/out/$filename  $rootdir/loginserver
	sh $rootdir/build/empty.asset.sh ${rootdir}/loginserver $filename

fi
echo ""
echo "---------打包loginserver(loginweb)-success---------------" 
echo ""