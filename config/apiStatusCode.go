package config

const SUCCESSRES int16 = 200    //接口请求成功之后的状态码
const NOTGETDATA int16 = 204    //接口请求失败,不能正常获取数据的都使用此状态码
const TOKENEMPTY int16 = -3     //token不存在,需要令牌
const TOKENEXPIRED int16 = -1   //token已经过期了
const PARAMERROR int16 = -4     //参数错误,传过来的参数不完整或者错误
const INTERNALERROR int16 = 503 //程序内部错误

//支付相关
const CREDENTIALSTOP int16 = -20      //支付证书停用
const CREDENTIALOSTOOMANY int16 = -21 //操作次数过多
const PAYACTIONERROR int16 = -11      //支付失败
const PAYANOALLOW int16 = -22         //非法操作
