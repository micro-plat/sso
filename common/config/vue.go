package config

import (
	"encoding/json"
	"fmt"

	"github.com/micro-plat/hydra"
)

//Vue VueConfig
func Vue(project string) func(ctx hydra.IContext) interface{} {
	return func(ctx hydra.IContext) interface{} {
		varConf := ctx.APPConf().GetVarConf()
		configData := map[string]interface{}{}
		_, err := varConf.GetObject("vues", project, &configData)
		if err != nil {
			return err
		}

		ctx.Response().ContentType("text/plain")
		bytes, _ := json.Marshal(configData)
		return fmt.Sprintf("window.globalConfig=%s", string(bytes))
	}
}
