#!/bin/sh

#############################################
# ./builid.sh 
#############################################

#获取当前目录
rootdir=$(pwd)

rm -rf $rootdir/out 
echo "" 
echo "-----------打包(mgrserver,loginserver)---"
echo ""
echo "1. 生成数据库sql 脚本"
phorcys markdown create sql -db mysql -file 0docs/2.设计/db.mysql.md -o out/mysql -c
if [ $? -ne 0 ]; then
	echo "phorcys 生成数据库脚本失败."
	exit 1
fi

rm -f out/mysql/all.sql

echo "2. 使用go-bindata 整合SQL文件到"
go-bindata -o=mgrserver/mgrapi/modules/const/sqls/mysql/dbscheme.go -pkg=mysql out/mysql/*
if [ $? -ne 0 ]; then
	echo "go-bindata 整合SQL出错"
	exit 1
fi

rm -rf out/mysql 


echo "3. 打包处理mgrserver"
cd $rootdir/mgrserver/mgrweb
echo "a. 下载npm 数据包：npm install"
#npm install 
if [ $? -ne 0 ]; then
	echo "npm install 出错"
	exit 1
fi

echo "b. 打包项目：npm run build"
npm run build 
if [ $? -ne 0 ]; then
	echo "npm run build 出错"
	exit 1
fi

echo "c. 压缩：dist/static"
cd dist/static
tar -zcvf static.tar.gz *
if [ $? -ne 0 ]; then
	echo "tar -zcvf static.tar.gz dist/static/* 出错"
	exit 1
fi

mkdir -p ${rootdir}/out/mgrserver/

mv static.tar.gz ${rootdir}/out/mgrserver/

sleep 0.1

cd $rootdir
echo "d. 使用go-bindata 整合static文件"
go-bindata -o=mgrserver/mgrapi/web/static.go -pkg=web $rootdir/out/mgrserver/static.tar.gz
if [ $? -ne 0 ]; then
	echo "go-bindata 整合static出错"
	exit 1
fi


echo "e. 生成mgrserver"
cd $rootdir/mgrserver
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o "$rootdir/out/mgrserver/bin/mgrserver"
if [ $? -ne 0 ]; then
	echo "mgrserver 项目编译出错,请检查"
	exit 1
fi

echo ""
echo "---------打包mgrserver-success-------------------" 
echo ""
echo "4. 打包处理loginserver"

cd $rootdir/loginserver/loginweb
echo "a. 下载npm 数据包：npm install"
#npm install 
if [ $? -ne 0 ]; then
	echo "npm install 出错"
	exit 1
fi

echo "b. 打包项目：npm run build"
npm run build 
if [ $? -ne 0 ]; then
	echo "npm run build 出错"
	exit 1
fi

echo "c. 压缩：dist/static"
cd dist/static
tar -zcvf static.tar.gz *
if [ $? -ne 0 ]; then
	echo "tar -zcvf static.tar.gz dist/static/* 出错"
	exit 1
fi

mkdir -p ${rootdir}/out/loginserver/

mv static.tar.gz ${rootdir}/out/loginserver/

sleep 0.1

cd $rootdir
echo "d. 使用go-bindata 整合static文件"
go-bindata -o=loginserver/loginapi/web/static.go -pkg=web  ${rootdir}/out/loginserver/static.tar.gz
if [ $? -ne 0 ]; then
	echo "go-bindata 整合static出错"
	exit 1
fi


echo "e. 生成loginserver"
cd $rootdir/loginserver
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o "$rootdir/out/loginserver/bin/loginserver"
if [ $? -ne 0 ]; then
	echo "loginserver 项目编译出错,请检查"
	exit 1
fi


#rm -rf ${rootdir}/out/mgrserver/static.tar.gz
#rm -rf ${rootdir}/out/loginserver/static.tar.gz
#rm -rf ${rootdir}/mgrserver/mgrweb/dist
#rm -rf ${rootdir}/loginserver/loginweb/dist


echo ""
echo "---------打包loginserver-success-------------------" 
echo ""
echo "-----------打包内容${rootdir}/out目录中-------------"