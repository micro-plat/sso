rootdir=$(pwd)

rm -rf $rootdir/loginserver/out

echo "----------生成loginserver数据-----------"
cd $rootdir/loginserver/
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -o "$rootdir/loginserver/out/bin/loginserver"
if [ $? -ne 0 ]; then
	echo "loginserver 项目编译出错,请检查"
	exit 1
fi

echo "----------压缩并拷贝static文件到out-----------"
cd $rootdir/loginserver/loginweb/dist/static/
zip -r static.zip static/ index.html
if [ $? -ne 0 ]; then
	echo "压缩文件失败"
	exit 1
fi
mv ./static.zip $rootdir/loginserver/out/bin

echo "----------启动loginserver服务-----------"
cd $rootdir/loginserver/out/bin
./loginserver run -r zk://192.168.0.101
if [ $? -ne 0 ]; then
	echo "loginserver 启动失败"
	exit 1
fi