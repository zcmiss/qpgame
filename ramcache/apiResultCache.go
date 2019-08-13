/*
这里存放api结果缓存
*/
package ramcache

import "sync"

//缓存手机验证码，key是手机号，数组里记录该验证码和过期时间
var PhoneCheckCode sync.Map

//缓存图片验证码，key是，数组里记录该验证码和过期时间
var UserNameCheckCode sync.Map
