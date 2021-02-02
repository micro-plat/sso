import Vue from "vue"

//相关状态text的颜色样式
var textColor = {
    0: "text-success",
    1: "text-warning",
    10: "text-muted",
    90: "text-danger"
}

//相关状态背景的颜色样式
var bgColor = {
    0: "bg-success",
    1: "bg-warning",
    10: "bg-primary",
    90: "bg-danger"
}

/*
* 枚举对象使用时须通过引用并进行初始化
* import utility from './filter'
* Vue.use(utility);
*/
export function Utility() {
    var that = Vue.prototype

    let defColor = textColor
    Object.assign(defColor, that.$env.conf.textColor || {})
    that.$env.conf.textColor = defColor


    let defBGColor = bgColor
    Object.assign(defBGColor, that.$env.conf.bgColor || {})
    that.$env.conf.bgColor = defBGColor
}

//字符串去空格
Utility.prototype.trim = function(str) {
    if (!str)
        return "";
    return str.toString().replace(/[ |-]/g, "");
}

//是否是手机号码
Utility.prototype.isPhoneNumber = function (phone) {
    let reg = /^1(3[0-9]|4[5,7]|5[0,1,2,3,5,6,7,8,9]|6[2,5,6,7]|7[0,1,7,8]|8[0-9]|9[1,8,9])\d{8}$/;
    return reg.test(phone);
}

//是否是邮箱
Utility.prototype.isEmailNumber = function (email) {
    let reg = /^[a-zA-Z0-9._-]+@[a-zA-Z0-9._-]+\.[a-zA-Z0-9._-]+$/;
    return reg.test(email);
}

//格式化卡号(卡号4位一空格)
Utility.prototype.cardNumberFormat = function (cardNumber) {
    if(!cardNumber){
        return "---"
    }
    if(typeof cardNumber != "string"){
        throw new Error("无效参数，需传入string类型的数据");
    }
    if(/\S{5}/.test(cardNumber)){
        return cardNumber.replace(/\s/g, '').replace(/(.{4})/g, "$1 ");
    }
    return cardNumber
}

//手机号格式化（中间四位显示'*'）
Utility.prototype.phoneFormat = function (phone){
    var reg = /^(\d{3})\d{4}(\d{4})$/;
    return phone.replace(reg, "$1****$2");
}

/*
* 参数说明：
* number：要格式化的数字
* decimals：保留几位小数
* 小数点符号"."
* 千分位符号","
* 四舍五入取值
*/
Utility.prototype.numberFormat = function (number, decimals = 2) {
    if(!number){
        return "---"
    }
    number = (number + '').replace(/[^0-9+-Ee.]/g, '')
    var n = !isFinite(+number) ? 0 : +number
    var prec = !isFinite(+decimals) ? 0 : Math.abs(decimals)
    var toFixedFix = function (n, prec) {
        var k = Math.pow(10, prec)
        return '' + parseFloat(Math['round'](parseFloat((n * k).toFixed(prec * 2))).toFixed(prec * 2)) / k
    }
    var s = (prec ? toFixedFix(n, prec) : '' + Math.round(n)).split('.')
    var re = /(-?\d+)(\d{3})/
    while (re.test(s[0])) {
        s[0] = s[0].replace(re, '$1' + "," + '$2')
    }
  
    if ((s[1] || '').length < prec) {
        s[1] = s[1] || ''
        s[1] += new Array(prec - s[1].length + 1).join('0')
    }
    return s.join(".")
}

//字符串去空格
Vue.filter('fltrTrim', value => {
    return Utility.prototype.trim(value)
})

//格式化卡号(卡号4位一空格)
Vue.filter('fltrCardNumberFormat', value => {
    return Utility.prototype.cardNumberFormat(value)
})

//手机号(手机号中间4位显示*)
Vue.filter('fltrPhoneNumberFormat', value => {
    return Utility.prototype.phoneFormat(value)
})

//要格式化的数字
Vue.filter('fltrNumberFormat', (value, decimals = 2) => {
    return Utility.prototype.numberFormat(value, decimals)
})

//日期格式转换
Vue.filter('fltrDate', (value, format = "yyyy-MM-dd") => {
    if (value === '') {
        return '-'
    } 
    return Utility.prototype.dateFormat(value, format)
})

//完整日期格式转换
Vue.filter('fltrDateTime', (value, format = "yyyy-MM-dd hh:mm") => {
    if (value === '') {
        return '-'
    } 
    return Utility.prototype.dateFormat(value, format)
})
  
//空值时默认显示'---'
Vue.filter('fltrEmpty', value => {
    if (value === '') {
        return '---'
    }
    return value
})
  
// 字符串截取number个字符，超出加'...'
Vue.filter('fltrSubstr', (value, number = 16) => {
    if (!value){
        return '-'
    }
    if (value.length <= number) {
        return value
    }
    return value.slice(0, number - 1) + '...'
})

// 根据value获取text不同颜色的class
Vue.filter('fltrTextColor', value => {
    return Vue.prototype.$env.conf.textColor[value] || "text-info"
})

// 根据value获取背景不同颜色的class
Vue.filter('fltrBgColor', value => {
    return Vue.prototype.$env.conf.bgColor[value] || "bg-info"
})

//日期格式转换
Utility.prototype.dateFormat = function (date, fmt = "yyyy-MM-dd hh:mm") {
    if(!date){
        return ""
    }
    let d = new Date(Date.parse(date))
    var o = {
        'M+': d.getMonth() + 1, // 月份
        'd+': d.getDate(), // 日
        'h+': d.getHours(), // 小时
        'm+': d.getMinutes(), // 分
        's+': d.getSeconds(), // 秒
        'q+': Math.floor((d.getMonth() + 3) / 3), // 季度
        'S': d.getMilliseconds() // 毫秒
    }
    if (/(y+)/.test(fmt)) {
        fmt = fmt.replace(RegExp.$1, (d.getFullYear() + '').substr(4 - RegExp.$1.length))
    }
    for (var k in o) {
        if (new RegExp('(' + k + ')').test(fmt)) {
            fmt = fmt.replace(RegExp.$1, (RegExp.$1.length === 1) ? (o[k]) : (('00' + o[k]).substr(('' + o[k]).length)))
        }
    }
    return fmt
}