#!/bin/sh

rm -rf ./out 
echo ""
echo "-----------(默认生成 prod环境, 开发环境传 dev)--------"

read -p "所有配置参数都改好了？确认请输入[y],否则输入[n]: " flag
temp=$(echo $flag | tr [A-Z] [a-z])
if [ $temp != "y" ]; then
	echo "请前去修改"
	exit 1
fi

echo "------------------------------------"

publishenv="prod"
if [ $# -eq 1 ] && [ $1 = "dev" ]; then 
	publishenv=$1
fi

echo "当前生成的环境为: $publishenv" 
echo "------------------------------------"
echo "---------------打包开始--------------"

tags=""
if [ $publishenv != "dev" ]; then
	tags=" -tags prod "
fi

echo "----------1:生成apiserver数据-----------"
cd apiserver/apiserver/
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  $tags -o "../out/sso/apiserver/bin/apiserver_sso"
if [ $? -ne 0 ]; then
	echo "apiserver 项目编译出错,请检查"
	exit 1
fi
cd ../../


echo "----------2:生成lgapi数据------------"
cd logingserver/lgapi/
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build $tags -o "../out/sso/logingserver/bin/lgapi_sso"
if [ $? -ne 0 ]; then
	echo "lgapi 项目编译出错,请检查"
	exit 1
fi
cd ../

echo "----------3:生成lgweb数据-----------"
cd lgweb/
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build $tags -o "../out/sso/logingserver/bin/lgweb_sso"
if [ $? -ne 0 ]; then
	echo "lgweb golang 项目编译出错,请检查"
	exit 1
fi

rm -rf ./dist/
echo ""
echo "--------------------------------------"
npm run build
if [ $? -ne 0 ]; then
	echo "lgweb vue项目编译出错,请检查"
	exit 1
fi

cp -r ./dist/static/ ../out/sso/logingserver/bin/static/

echo "--------------------------------------"
cd ../../


echo "----------4:生成mgrapi数据----------"
cd mgrserver/mgrapi/
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  $tags -o "../out/sso/mgrserver/bin/mgrapi_sso"
if [ $? -ne 0 ]; then
	echo "mgrapi 项目编译出错,请检查"
	exit 1
fi
cd ../

echo "-------创建mgrapi图片临时目录--------"
mkdir -p ./out/sso/mgrserver/bin/image

echo "----------5:生成mgrweb数据------------"
cd mgrweb/
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  $tags -o "../out/sso/mgrserver/bin/mgrweb_sso"
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
cp -r ./dist/static/ ../out/sso/mgrserver/bin/static/
echo "--------------------------------------"

cd ../
echo "-----------6:生成数据完成--------------"
echo "-----------都放在out目录中-------------"


# cd out/
# echo "-----------打包相关文件(zip)-----------"
# zip -r sso_apiserver sso_apiserver
# if [ $? -ne 0 ]; then
# 	echo "打包ssoapi调用(sso_apiserver)出错,请检查"
# 	exit
# fi

# zip -r sso_login sso_login
# if [ $? -ne 0 ]; then
# 	echo "打包登录中心(sso_login)出错,请检查"
# 	exit
# fi

# zip -r sso_mgr sso_mgr
# if [ $? -ne 0 ]; then
# 	echo "打包用户管理系统(sso_mgr)出错,请检查"
# 	exit
# fi

# echo "-----------------------------------"
# echo "sso_login里面包含api,web | sso_mgr里面包含api,web"
# echo "-----------打包完成(zip)-------------"
