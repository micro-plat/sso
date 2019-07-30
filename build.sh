#!/bin/sh

echo "sso五个项目的打包脚本...."

echo "请正确配置所有的config信息"

echo "1:生成apiserver数据"

#out_dir ="out"
#if [ ! -d "$out_dir" ]; then
#        mkdir $out_dir
#fi

apiserver_out_dir="$out/$apiserver"
if [ ! -d "$apiserver_out_dir" ]; then
        mkdir $apiserver_out_dir
fi

cd apiserver
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $apiserver_out_dir main.go
if [ $? -ne 0 ]; then
	echo "apiserver 项目编译出错,请检查"
	exit
fi


echo "2:生成lgapi数据"

echo "3:生成lgweb数据"

echo "4:生成mgrapi数据"

echo "5:生成mgrweb数据"

echo "生成数据完成"