import vue from "vue";

//所有保存的枚举数据
window._EnumList_ = {};
//根据type保存的回调函数
window._EnumCallbackFunc_ = {
  "*" : function(tp){return [];}
}

/*
* 枚举对象使用时须通过引用并进行初始化
* import enums from './enums'
* Vue.use(enums);
*/
export function Enum() {}


//当参数为function时回调
Enum.prototype.callback = function (callback, tp) {
  if (typeof callback == "function"){
    var type = tp || "*"
    window._EnumCallbackFunc_[type] = callback
    return
  }
  throw new Error("无效参数，类型:function(type){return [];}返回一个数组");
};

//数据保存
Enum.prototype.set = function (data, type) {
  if (typeof data == "function"){
    if (!type){
        return
    }
    window._EnumCallbackFunc_[type] = data
    return
  }
  
  if (!Array.isArray(data) || data.length == 0) { 
    return
  }

  if(!checkData(data)){
    throw new Error("数据格式错误，格式:[{type:'zxc',name:'上架',value:'1', pid: 1002},...]其中type与pid参数为非必要参数，name与value为必要参数");
  }

  data.forEach(function (item) {
    var tp = item.type || type
    if(tp){
      if (!window._EnumList_[tp]) {
        window._EnumList_[tp] = [];
      }
      window._EnumList_[tp].push(item);
    }
  });
  return
};

//数据获取
Enum.prototype.get = function (type, pid) {
  if (!type) return [];

  var result = getEnumList(type, pid) //根据存储查询
  if (result.length == 0){
    result = getEnumListByCallback(type, pid)//根据回调获取
  }
  return result
};

//根据value值获取name
Enum.prototype.getName = function (type, value) {
  if (value == "") {
    return "-"
  }

  var enumMap = Enum.prototype.get(type)
  for (var i = 0; i < enumMap.length; i++){
    if (enumMap[i].value == value) {
      return enumMap[i].name
    }
  }
  return value 
}

//对应type数据刷新
Enum.prototype.clear = function (type) {
  window._EnumList_[type] = null;
};

//filter
export const fltrEnum = vue.filter('fltrEnum', (value, enumType) => {
  return new Enum().getName(enumType, value)
})

//数据格式检查
function checkData(data){
  for (var i = 0; i < data.length; i++){
    if(!data[i].hasOwnProperty("name") || !data[i].hasOwnProperty("value")){
      return false
    }
  }
  return true
}

//根据回调函数获取数据
function getEnumListByCallback(type, pid){
  var handle = window._EnumCallbackFunc_[type] || window._EnumCallbackFunc_["*"]
  
  var data = handle.apply(this, [type])
  Enum.prototype.set(data, type)
  return getEnumList(type, pid)
}

//根据type从window._EnumList_中获取相应的数据
function getEnumList(type, pid){
  var result = window._EnumList_[type] || []
  if (!pid){
    return result
  }
  var list = []
  result.forEach((item)=>{
    if (item.pid == pid){
      list.push(item)
    }
  })
  return list
}