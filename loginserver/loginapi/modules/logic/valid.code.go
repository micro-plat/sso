package logic

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/lib4dev/vcs"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/components/uuid"
	"github.com/micro-plat/lib4go/errs"
	"github.com/micro-plat/lib4go/types"
	"github.com/micro-plat/sso/loginserver/loginapi/modules/access/system"
	"github.com/micro-plat/sso/loginserver/loginapi/modules/model"
	"github.com/micro-plat/sso/sso/errorcode"

	"net/http"
)

//ValidCodeLogic 登录验证码
type ValidCodeLogic struct {
	sysDB system.IDbSystem
}

//NewValidCodeLogic 创建登录对象
func NewValidCodeLogic() *ValidCodeLogic {
	return &ValidCodeLogic{
		sysDB: system.NewDbSystem(),
	}
}

//SendSmsCode SendSmsCode
func (l *ValidCodeLogic) SendSmsCode(userInfo types.XMap, ident, validCode string) (err error) {
	loginCfg := model.GetLoginConf()

	if len(userInfo.GetString("mobile")) == 0 {
		return errs.NewError(errorcode.ERR_USER_NOTBIND_PHONE, fmt.Errorf("用户没有绑定手机号"))
	}

	reqID := uuid.GetSUUID("validcode").Get().ToString("SSO")
	params := &vcs.SendRequest{
		ReqID:       reqID,
		Ident:       ident,
		UserAccount: userInfo.GetString("mobile"),
		TemplateID:  loginCfg.SMSTemplateID,
		Keywords:    validCode,
	}

	if _, err := vcs.SendSmsCode(params, hydra.G.GetPlatName()); err != nil {
		return errs.NewError(errorcode.ERR_SYS_ERROR, err)
	}
	return nil
}

//SendWechatCode 发送微信验证码
func (l *ValidCodeLogic) SendWechatCode(userInfo types.XMap, ident, validCode string) error {
	token, err := l.GetAccessToken()
	if err != nil {
		return err
	}
	if len(userInfo.GetString("wx_openid")) == 0 {
		return errs.NewError(errorcode.ERR_USER_NOTBINDWX, fmt.Errorf("用户没有绑定微信"))
	}
	sysTitle := ""
	if !strings.EqualFold(ident, "") {
		fmt.Println("ident:", ident)
		system, _ := l.sysDB.QuerySysInfoByIdent(ident)
		if system != nil {
			sysTitle = system.GetString("name")
		}
	}

	loginCfg := model.GetLoginConf()
	reqID := uuid.GetSUUID("validcode").Get().ToString("SSO")
	params := &vcs.SendRequest{
		ReqID:       reqID,
		Ident:       ident,
		UserAccount: userInfo.GetString("wx_openid"),
		TemplateID:  loginCfg.SMSTemplateID,
		ExtParams: types.XMap{
			"access_token": token,
			"system_title": sysTitle,
		},
		Keywords: validCode,
	}

	if _, err := vcs.SendSmsCode(params, hydra.G.GetPlatName()); err != nil {
		return errs.NewError(errorcode.ERR_SYS_ERROR, err)
	}

	return nil
}

//GetAccessToken 动态获取token
func (l *ValidCodeLogic) GetAccessToken() (string, error) {
	conf := model.GetLoginConf()
	url := fmt.Sprintf("%s/%s/wechat/token/get", conf.WechatTokenHost, conf.WechatAppID)
	client := hydra.C.HTTP().GetRegularClient()
	body, statusCode, err := client.Get(url)
	if err != nil {
		err = errs.NewErrorf(errorcode.ERR_WECHAT_TOKEN_GET_ERROR, "获取wechat/token失败：%v", err)
		return "", err
	}
	if statusCode != 200 {
		return "", errs.NewErrorf(statusCode, "获取token信息失败,HttpStatus:%d, body:%s", statusCode, body)
	}

	data := make(map[string]interface{})
	err = json.Unmarshal([]byte(body), &data)
	if err != nil {
		return "", fmt.Errorf("字符串转json发生错误，err：%v", err)
	}

	if errcode, ok := data["errcode"]; ok && types.GetInt(errcode) != 0 {
		return "", errs.NewError(http.StatusNotExtended, fmt.Errorf("获取token失败: %s", types.GetString(data["errmsg"])))
	}

	return types.GetString(data["access_token"]), nil
}

// //构造发送验证码的实体数据
// func (m *LoginLogic) constructSendData(openID, validCode, ident, templateID string) model.TemplateMsg {
// 	sendTitle := "用户系统"
// 	if !strings.EqualFold(ident, "") {
// 		system, _ := m.sysDB.QuerySysInfoByIdent(ident)
// 		if system != nil {
// 			sendTitle = system.GetString("name")
// 		}
// 	}

// 	data := model.TemplateMsg{
// 			Touser:     openID,
// 			TemplateID: templateID,
// 		Data: &model.TemplateData{
// 			First: model.KeyWordData{
// 				Value: fmt.Sprintf("你好,登录到[%s]需要进行验证码验证,请无泄露", sendTitle),
// 				Color: "#173177",
// 			},
// 			Keyword1: model.KeyWordData{
// 				Value: validCode,
// 				Color: "#FF0000",
// 			},
// 			Keyword2: model.KeyWordData{
// 				Value: "5分钟",
// 				Color: "#173177",
// 			},
// 			Keyword3: model.KeyWordData{
// 				Value: time.Now().Format("2006-01-02 15:04:05"),
// 				Color: "#173177",
// 			},
// 			Remark: model.KeyWordData{
// 				Value: "若非本人操作,可能你的账号存在安全风险,请及时修改密码",
// 				Color: "#173177",
// 			},
// 		},
// 	}
// 	return data
// }
