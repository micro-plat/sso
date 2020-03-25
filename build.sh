#!/bin/sh

#############################################
# ./builid.sh 线上环境， ./build.sh dev 线下环境
#############################################

#echo $@ 多个参数时 动态判断

# for i in $@;  
# do  
# if [ $i = "oci" ]; then 
# 	echo "oci"
# fi;

# if [ $i = "prod" ]; then 
# 	echo "prod"
# fi
# done  


#获取当前目录
rootdir=$(pwd)

rm -rf $rootdir/out 
echo ""
echo "-----------(默认生成 prod环境, 开发环境传 dev)--------"
echo ""
echo "----此脚本不能打包17ebs_sso相关的包,要做修改才行 ！！！！！！！---"
echo "----17ebs_sso只需要打包三个文件(mgrserver/mgrapi, mgrserver/mgrweb, apiserver/apiserver)---"

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
echo "tags: $tags"
echo ""

echo "----------1:生成apiserver数据-----------"
cd $rootdir/apiserver/apiserver/
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  $tags -o "$rootdir/out/sso/apiserver/bin/apiserver"
if [ $? -ne 0 ]; then
	echo "apiserver 项目编译出错,请检查"
	exit 1
fi


echo "----------2:生成lgapi数据------------"
cd $rootdir/loginserver/lgapi/
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build $tags -o "$rootdir/out/sso/loginserver/bin/lgapi"
if [ $? -ne 0 ]; then
	echo "lgapi 项目编译出错,请检查"
	exit 1
fi


echo "----------3:生成lgweb数据-----------"
cd $rootdir/loginserver/lgweb/
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build $tags -o "$rootdir/out/sso/loginserver/bin/lgweb"
if [ $? -ne 0 ]; then
	echo "lgweb golang 项目编译出错,请检查"
	exit 1
fi

rm -rf $rootdir/loginserver/lgweb/dist/
echo ""
echo "--------------------------------------"
npm run build
if [ $? -ne 0 ]; then
	echo "lgweb vue项目编译出错,请检查"
	exit 1
fi

cp -r $rootdir/loginserver/lgweb/dist/static/ $rootdir/out/sso/loginserver/bin/static/
echo "--------------------------------------"



echo "----------4:生成mgrapi数据----------"
cd $rootdir/mgrserver/mgrapi/
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  $tags -o "$rootdir/out/sso/mgrserver/bin/mgrapi"
if [ $? -ne 0 ]; then
	echo "mgrapi 项目编译出错,请检查"
	exit 1
fi

echo "-------创建mgrapi图片临时目录--------"
mkdir -p $rootdir/out/sso/mgrserver/image

echo "----------5:生成mgrweb数据------------"
cd $rootdir/mgrserver/mgrweb/
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  $tags -o "$rootdir/out/sso/mgrserver/bin/mgrweb"
if [ $? -ne 0 ]; then
	echo "mgrweb golang 项目编译出错,请检查"
	exit
fi

rm -rf $rootdir/mgrserver/mgrweb/dist/

echo ""
echo "--------------------------------------"
npm run build
if [ $? -ne 0 ]; then
	echo "mgrweb vue项目编译出错,请检查"
	exit
fi
cp -r $rootdir/mgrserver/mgrweb/dist/static/ $rootdir/out/sso/mgrserver/bin/static/
echo "--------------------------------------"

echo "-----------6:生成数据完成--------------"

cd $rootdir/out/sso/loginserver/bin/
echo "-----------压缩lgweb前端static文件夹-----------"
zip -r static static
if [ $? -ne 0 ]; then
	echo "压缩lgweb前端static文件夹出错,请检查"
	exit
fi
rm -rf static

cd $rootdir/out/sso/mgrserver/bin/
echo "-----------压缩mgrweb前端static文件夹-----------"
zip -r static static
if [ $? -ne 0 ]; then
	echo "压缩mgrweb前端static文件夹出错,请检查"
	exit
fi
rm -rf static

echo "-----------都放在out目录中-------------"




