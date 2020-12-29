#!/bin/sh

#############################################
# ./builid.sh pkg
#############################################

#获取当前目录
rootdir=$(pwd)
PATH=$PATH:$GOPATH/bin

pkg=$1

build_static(){
	filepath=$1
	echo "" > $filepath/static.go
	echo "package web" >> $filepath/static.go
	echo "import \"fmt\"" >> $filepath/static.go
	echo "func Asset(name string) ([]byte, error) { return nil, fmt.Errorf(\"Asset %s not found\", name) }" >> $filepath/static.go
	echo "func AssetNames() []string              { return []string{} }" >> $filepath/static.go
}


echo "" 
echo "-----------打包检查（phorcys， go-bindata) ---"
echo "" 
echo "1. 检查 phorcys 存在"
which phorcys
if [ $? -ne 0 ]; then
	echo "1.a. phorcys 不存在，执行安装"
	
	dir=$GOPATH/src/gitlab.100bm.cn/devtools/phorcys/phorcys
	if [ ! -d  $dir ] ; then
		cd $GOPATH/src/
		mkdir -p gitlab.100bm.cn/devtools/phorcys
		cd gitlab.100bm.cn/devtools/phorcys

		echo "1.b. 下载phorcys"
		git clone https://gitlab.100bm.cn/devtools/phorcys/phorcys.git
		
		git checkout dev
	fi
	cd $dir
	git pull 
	echo "1.c. 编译phorcys"
	go install 

fi 

cd $rootdir
echo "1. 检查 go-bindata 存在"

which go-bindata
if [ $? -ne 0 ]; then
	echo "2.a. go-bindata 不存在，执行安装"
	dir=$GOPATH/src/github.com/go-bindata/go-bindata
	if [ ! -d  $dir ] ; then
		cd $GOPATH/src
		mkdir -p github.com/go-bindata
		cd github.com/go-bindata
		
		echo "2.b. 下载go-bindata"
		git clone https://github.com/go-bindata/go-bindata.git
	fi
	
	cd $dir
	git pull 
	cd go-bindata
 	echo "2.c. 编译go-bindata"
	go install 
fi 

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


echo "3. 打包处理mgrserver ${rootdir}/mgrserver/mgrweb"
cd $rootdir/mgrserver/mgrweb
#echo "a. 下载npm 数据包：npm install"
#npm install 
# if [ $? -ne 0 ]; then
# 	echo "npm install 出错"
# 	exit 1
# fi

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

if [ "$pkg" = "pkg" ] ; then 
	echo "d. 使用go-bindata 整合static文件"
	sleep 0.1
	cd $rootdir/out/mgrserver/
	go-bindata -o=${rootdir}/mgrserver/mgrapi/web/static.go -pkg=web static.tar.gz
	if [ $? -ne 0 ]; then
		echo "go-bindata 整合static出错"
		exit 1
	fi
else
	echo "build_static 1"
	build_static ${rootdir}/mgrserver/mgrapi/web
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
#echo "a. 下载npm 数据包：npm install"
#npm install 
# if [ $? -ne 0 ]; then
# 	echo "npm install 出错"
# 	exit 1
# fi

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

if [ "$pkg" = "pkg" ] ; then 
	echo "d. 使用go-bindata 整合static文件"
	sleep 0.1
	cd ${rootdir}/out/loginserver/
	go-bindata -o=${rootdir}/loginserver/loginapi/web/static.go -pkg=web  static.tar.gz
	if [ $? -ne 0 ]; then
		echo "go-bindata 整合static出错"
		exit 1
	fi
else
	echo "build_static 2"
	build_static ${rootdir}/loginserver/loginapi/web
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