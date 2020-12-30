
#!/bin/sh

PATH=$PATH:$GOPATH/bin
rootdir=$(pwd)

echo ""
echo "---------generate.db--------------------" 
echo ""

echo "1. 生成数据库sql 脚本"
phorcys markdown create sql -db mysql -file 0docs/2.设计/db.mysql.md -o out/mysql/scheme -c > /dev/null
if [ $? -ne 0 ]; then
	echo "phorcys 生成数据库脚本失败."
	exit 1
fi

rm -rf out/mysql/data/*
mkdir -p out/mysql/data

cp 0docs/2.设计/*.sql out/mysql/data

echo ""
echo "---------generate.db-success-----------------" 
echo ""