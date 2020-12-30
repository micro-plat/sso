
#!/bin/sh

PATH=$PATH:$GOPATH/bin
rootdir=$(dirname $(pwd)) 

echo ""
echo "---------generate.db--------------------" 
echo ""
cd $rootdir

which phorcys > /dev/null
if [ $? -ne 0 ]; then
	echo "phorcys未安装"
	echo "请到https://gitlab.100bm.cn/devtools/phorcys/phorcys.git下载安装"
	exit 1	
fi

if [ -d $rootdir/0docs/2.设计/mysql/scheme ] ; then 
	echo  $rootdir/0docs/2.设计/mysql/scheme"文件夹已存在，如需生成，请删除该文件夹"
	echo ""
	echo ""
	exit 1 
fi 


echo "1. 生成数据库sql 脚本"
phorcys markdown create sql -db mysql -file $rootdir/0docs/2.设计/db.mysql.md -o $rootdir/out/mysql/scheme -c
if [ $? -ne 0 ]; then
	echo "phorcys 生成数据库脚本失败."
	exit 1
fi


rm -rf out/mysql/data/*
mkdir -p out/mysql/data

cp $rootdir/0docs/2.设计/mysql/data/*.sql $rootdir/out/mysql/data

mkdir -p  $rootdir/0docs/2.设计/mysql/scheme/

cp -r out/mysql/scheme/*.sql  $rootdir/0docs/2.设计/mysql/scheme/

rm -f out/mysql/scheme/all.sql

echo "2. 使用go-bindata 整合Scheme文件"
go-bindata -o=mgrserver/mgrapi/modules/const/sqls/mysql/scheme/scheme.go -pkg=scheme out/mysql/scheme/*  > /dev/null 
if [ $? -ne 0 ]; then
	echo "go-bindata 整合SQL出错"
	exit 1
fi
 

echo "3. 使用go-bindata 整合Data文件"
go-bindata -o=mgrserver/mgrapi/modules/const/sqls/mysql/data/data.go -pkg=data out/mysql/data/*   > /dev/null
if [ $? -ne 0 ]; then
	echo "go-bindata 整合SQL出错"
	exit 1
fi
rm -rf $rootdir/out/mysql
echo ""
echo "---------generate.db-success-----------------" 
echo ""