#!/bin/sh

filepath=$1
echo '
package main

import (
	"path"

	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/conf/server/static"
)

func init() {
	hydra.OnReady(func() {			
		for _, v := range AssetNames() {
			ext := path.Ext(v)
			embed, _ := Asset(v)
			staticOpts = append(staticOpts,static.WithArchiveByEmbed(embed, ext))			
			hydra.Conf.GetWeb().Static(staticOpts...)
		}
	})
}
' > $filepath/web.go 