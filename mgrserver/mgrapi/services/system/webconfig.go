package system 
 
import ( 
  "encoding/json" 
  "fmt" 
  
  "github.com/micro-plat/hydra/conf/server/auth/jwt"
  "github.com/micro-plat/hydra" 
) 
 
//WebConfigHandler WebConfigHandler 
func WebConfigHandler(ctx hydra.IContext) interface{} { 
  configData := map[string]interface{}{} 
  _, err := ctx.APPConf().GetServerConf().GetSubObject("webconf", &configData) 
  if err != nil { 
    return  err
  } 
  jwtConf, err := jwt.GetConf(ctx.APPConf().GetServerConf())
	if err != nil {
		return err
	}

	configData["jwt_name"] = jwtConf.Name
	configData["jwt_source"] = jwtConf.Source
	configData["jwt_authurl"] = jwtConf.AuthURL

  ctx.Response().ContentType("text/plain") 
  bytes, _ := json.Marshal(configData) 
  return fmt.Sprintf("window.globalConfig=%s", string(bytes)) 
} 