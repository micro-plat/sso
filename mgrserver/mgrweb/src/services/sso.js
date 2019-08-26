import {httpConfig} from './sso.http';
import {ssoConfig} from './sso.login';

/**
 * @param {前端对应后台的api地址 apiHost} apiHost
 * @param {jwt存储地方, localStorage; 存储在localStorage, sessionStorage:存储在sessionStorage} storagePlace}
 * @param {sso前端地址} ssoHost}
 * @param {子系统标识} sysIdent}
 */
export function ssoHttpConfig(apiHost,storagePlace, ssoHost, sysIdent) {
    if (!apiHost || !ssoHost || !sysIdent) {
        console.log("ssoHttpConfig 输入参数有误");
        return
    }
    
    var httpExtend = httpConfig(apiHost, storagePlace);
    var ssocfg = ssoConfig(ssoHost, sysIdent)
    return {
        http: httpExtend,
        sso: ssocfg
    }
}