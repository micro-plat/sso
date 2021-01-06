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
			hydra.Conf.GetWeb().Static(static.WithArchiveByEmbed(embed, ext))
		}
	})
}
' > $filepath/web.go 