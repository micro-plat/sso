#!/bin/sh

rm -rf ./out 

echo "sso五个项目的打包脚本...."

read -p "所有配置参数都改好了？确认请输入[yes],否则输入[no]: " flag
temp=$(echo $flag | tr [A-Z] [a-z])
if [ $temp != "yes" ]; then
	echo "请前去修改"
	exit
fi


echo "----------1:生成apiserver数据-----------"
cd apiserver/
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o "../out/apiserver/apiserver"
if [ $? -ne 0 ]; then
	echo "apiserver 项目编译出错,请检查"
	exit
fi
cd ../

echo "----------2:生成lgapi数据------------"
cd lgapi/
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o "../out/lg/bin/lgapi"
if [ $? -ne 0 ]; then
	echo "lgapi 项目编译出错,请检查"
	exit
fi
cd ../

echo "----------3:生成lgweb数据-----------"
cd lgweb/
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o "../out/lg/web/lgweb"
if [ $? -ne 0 ]; then
	echo "lgweb golang 项目编译出错,请检查"
	exit
fi

rm -rf ./dist/
echo ""
echo "--------------------------------------"
npm run build
if [ $? -ne 0 ]; then
	echo "lgweb vue项目编译出错,请检查"
	exit
fi
cp -r ./dist/static/ ../out/lg/web/static/
echo "--------------------------------------"
cd ../


echo "----------4:生成mgrapi数据----------"
cd mgrapi/
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o "../out/mgr/bin/mgrapi"
if [ $? -ne 0 ]; then
	echo "mgrapi 项目编译出错,请检查"
	exit
fi
cd ../

echo "----------5:生成mgrweb数据------------"
cd mgrweb/
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o "../out/mgr/web/mgrweb"
if [ $? -ne 0 ]; then
	echo "mgrweb golang 项目编译出错,请检查"
	exit
fi
rm -rf ./dist/
echo ""
echo "--------------------------------------"
npm run build
if [ $? -ne 0 ]; then
	echo "mgrweb vue项目编译出错,请检查"
	exit
fi
cp -r ./dist/static/ ../out/mgr/web/static/
echo "--------------------------------------"

cd ../
echo "生成数据完成"