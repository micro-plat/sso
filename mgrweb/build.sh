#!/usr/bin/env bash
cd /root/work/src/mgrweb
npm run build
#go install
#cp ${GOPATH}/bin/mgrweb ./dist
cd dist
cp -r /root/work/img ./static/static/
#scp -r static root@192.168.106.129:/home/mgrweb/mgrweb
