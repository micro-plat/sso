rootdir=$(pwd)

rm -rf $rootdir/mgrserver/out

echo "----------生成mgrserver数据-----------"
cd $rootdir/mgrserver/
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -o "$rootdir/mgrserver/out/bin/mgrserver"
if [ $? -ne 0 ]; then
	echo "mgrserver 项目编译出错,请检查"
	exit 1
fi


echo "----------压缩并拷贝static文件到out-----------"
cd $rootdir/mgrserver/mgrweb/dist/static/
zip -r static.zip static/ index.html
if [ $? -ne 0 ]; then
	echo "压缩文件失败"
	exit 1
fi
mv ./static.zip $rootdir/mgrserver/out/bin

echo "----------启动mgrserver服务-----------"
cd $rootdir/mgrserver/out/bin
./mgrserver run -r zk://192.168.0.101
if [ $? -ne 0 ]; then
	echo "mgrserver 启动失败"
	exit 1
fi