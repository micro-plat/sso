package system 
 
import ( 
  "encoding/json" 
  "fmt" 
 
  "github.com/micro-plat/hydra" 
) 
 
//WebConfigHandler WebConfigHandler 
func WebConfigHandler(ctx hydra.IContext) interface{} { 
  configData := map[string]interface{}{} 
  _, err := ctx.APPConf().GetServerConf().GetSubObject("webconf", &configData) 
  if err != nil { 
    return  err
  } 
  ctx.Response().ContentType("text/plain") 
  bytes, _ := json.Marshal(configData) 
  return fmt.Sprintf("window.globalConfig=%s", string(bytes)) 
} 