package util

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"sort"
	"strings"

	"github.com/micro-plat/hydra"
)

//wxpayVerifySign 微信支付签名验证函数
func VerifySign(ctx hydra.IContext, needVerifyM map[string]interface{}, secret, sign string) bool {

	raw, signCalc := makeSign(needVerifyM, secret)
	ctx.Log().Infof("计算出来的sign: %v\n", signCalc)
	ctx.Log().Infof("请求sign: %v\n", sign)
	if sign == signCalc {
		ctx.Log().Infof("签名校验通过!")
		return true
	}

	ctx.Log().Infof("签名校验失败![%s]", raw)
	return false
}

func makeSign(mReq map[string]interface{}, secret string) (raw string, sign string) {
	fmt.Println("签名计算, API KEY:", secret)
	//1, 对key进行升序排序.
	sortedKeys := make([]string, 0)
	for k, _ := range mReq {
		sortedKeys = append(sortedKeys, k)
	}

	sort.Strings(sortedKeys)

	//2, 对key=value的键值对用&连接起来，略过空值
	var signStrings string
	for _, k := range sortedKeys {
		fmt.Printf("k=%v, v=%v\n", k, mReq[k])
		value := fmt.Sprintf("%v", mReq[k])
		if value != "" {
			signStrings = signStrings + k + "=" + value + "&"
		}
	}

	//3, 在键值对的最后加上key=secret
	if secret != "" {
		signStrings = signStrings + "key=" + secret
	}

	//4, 进行MD5签名并且将所有字符转为大写.
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(signStrings))
	cipherStr := md5Ctx.Sum(nil)
	upperSign := strings.ToUpper(hex.EncodeToString(cipherStr))
	return signStrings, upperSign
}
