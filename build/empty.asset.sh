#!/bin/sh

filepath=$1
asset_name=$2
echo "asset_name:${asset_name}" 
echo "
//+build !none

package main

import (
	_ \"embed\"

	\"github.com/micro-plat/hydra\"
	\"github.com/micro-plat/hydra/conf/server/static\"
)

//go:embed ${asset_name}
var archiveBytes []byte

func init() {
	hydra.OnReady(func() {
		staticOpts = append(staticOpts, static.WithEmbedBytes(Archive, archiveBytes))
		hydra.Conf.GetWeb().Static(staticOpts...)
	})
}
" > $filepath/web.go 