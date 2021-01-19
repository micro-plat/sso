package system

import (
	"encoding/json"

	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/conf/server/auth/jwt"
)

//WebConfigHandler WebConfigHandler
func WebConfigHandler(ctx hydra.IContext) interface{} {
	configData := map[string]interface{}{}
	jwtConf, err := jwt.GetConf(ctx.APPConf().GetServerConf())
	if err != nil {
		return err
	}

	configData["jwt_name"] = jwtConf.Name
	configData["jwt_source"] = jwtConf.Source
	configData["jwt_authurl"] = jwtConf.AuthURL

	ctx.Response().ContentType("text/plain")
	bytes, _ := json.Marshal(configData)
	return string(bytes)
}
