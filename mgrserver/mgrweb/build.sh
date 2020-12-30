#!/bin/sh


rootdir=$(pwd)

npm run build 





cd dist/static 
rm -f static.tar.gz
rm -f $rootdir/../../out/mgrserver/bin/static.tar.gz

tar -zcvf static.tar.gz *
mv ./static.tar.gz  $rootdir/../../out/mgrserver/bin

