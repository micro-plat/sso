#!/bin/sh

rm -rf ./out 

echo "------------------------------------"
echo "sso五个项目的打包脚本...."
if [ $# != 1 ]; then 
	echo "请输入要生成的环境[prod, dev]"
	exit 1
fi

publishenv=$1

if [ $publishenv != "prod" ] && [ $publishenv != "dev" ]; then 
	echo "请在命令行中输入要生成的环境[prod, dev]"
	exit 1
fi

echo "当前生成的环境为: $publishenv" 

tags=""
if [ $publishenv != "dev" ]; then
	tags=" -tags prod "
fi

echo "----------1:生成apiserver数据-----------"
cd apiserver/
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  $tags -o "../out/sso_apiserver/api/bin/sso_apiserver"
if [ $? -ne 0 ]; then
	echo "apiserver 项目编译出错,请检查"
	exit 1
fi
cd ../


echo "----------2:生成lgapi数据------------"
cd lgapi/
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build $tags -o "../out/sso_login/api/bin/sso_lgapi"
if [ $? -ne 0 ]; then
	echo "lgapi 项目编译出错,请检查"
	exit 1
fi
cd ../

echo "----------3:生成lgweb数据-----------"
cd lgweb/
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build $tags -o "../out/sso_login/web/sso_lgweb"
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

#mkdir -p ../out/sso_login/web/web/static/ && cp -r ./dist/static/ "$_"
cp -r ./dist/static/ ../out/sso_login/web/static/

echo "--------------------------------------"
cd ../


echo "----------4:生成mgrapi数据----------"
cd mgrapi/
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  $tags -o "../out/sso_mgr/api/bin/sso_mgrapi"
if [ $? -ne 0 ]; then
	echo "mgrapi 项目编译出错,请检查"
	exit 1
fi
cd ../

echo "-------创建mgrapi图片临时目录--------"
mkdir -p ./out/sso_mgr/api/bin/static/img

echo "----------5:生成mgrweb数据------------"
cd mgrweb/
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  $tags -o "../out/sso_mgr/web/sso_mgrweb"
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
cp -r ./dist/static/ ../out/sso_mgr/web/static/
echo "--------------------------------------"

cd ../
echo "-----------6:生成数据完成--------------"


cd out/
echo "-----------打包相关文件(zip)-----------"
zip -r sso_apiserver sso_apiserver
if [ $? -ne 0 ]; then
	echo "打包ssoapi调用(sso_apiserver)出错,请检查"
	exit
fi

zip -r sso_login sso_login
if [ $? -ne 0 ]; then
	echo "打包登录中心(sso_login)出错,请检查"
	exit
fi

zip -r sso_mgr sso_mgr
if [ $? -ne 0 ]; then
	echo "打包用户管理系统(sso_mgr)出错,请检查"
	exit
fi

echo "-----------------------------------"
echo "sso_login里面包含api,web | sso_mgr里面包含api,web"
echo "-----------打包完成(zip)-------------"



# read -p "所有配置参数都改好了？确认请输入[y],否则输入[n]: " flag
# temp=$(echo $flag | tr [A-Z] [a-z])
# if [ $temp != "y" ]; then
# 	echo "请前去修改"
# 	exit
# fi