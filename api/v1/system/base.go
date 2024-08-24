package sysapi

import (
	"github.com/gin-gonic/gin"
	"github.com/lazybearlee/yuedong-fitness/global"
	"github.com/lazybearlee/yuedong-fitness/model/common/response"
	sysmodel "github.com/lazybearlee/yuedong-fitness/model/system"
	sysrequest "github.com/lazybearlee/yuedong-fitness/model/system/request"
	"github.com/lazybearlee/yuedong-fitness/utils"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
	"time"
)

type BaseApi struct{}

var captchaStore = base64Captcha.DefaultMemStore

// Login
// @Tags Base
// @Summary 用户登录
// @Produce application/json
// @Param data body sysmodel.Login
// @Success 200 {object} response.Response{data=sysresponse.LoginResponse,msg=string} "登录成功"
// @Router /base/login [post]
func (b *BaseApi) Login(c *gin.Context) {
	var login sysrequest.LoginReq
	err := c.ShouldBindJSON(&login)
	if err != nil {
		response.ErrorWithMessage("参数绑定失败", c)
	}
	// 是否开启验证码
	isCaptchaOpen := global.FITNESS_CONFIG.Captcha.OpenCaptcha
	captchaTimeout := global.FITNESS_CONFIG.Captcha.OpenCaptchaTimeOut
	key := c.ClientIP()
	v, ok := global.FITNESS_CACHE.Get(key)
	if !ok {
		global.FITNESS_CACHE.Set(key, 1, time.Second*time.Duration(captchaTimeout)) // 设置缓存超时时间
	}
	var checkCaptcha bool = (isCaptchaOpen != 0 && isCaptchaOpen >= utils.InterfaceToInt(v)) || (login.CaptchaId != "" && login.Captcha != "" && captchaStore.Verify(login.CaptchaId, login.Captcha, true))
	if checkCaptcha {
		// 如果不使用验证码或者验证码验证通过
		// 登陆
		u := sysmodel.SysUser{Username: login.Username, Password: login.Password}
		user, err := userService.UserLogin(u)
		if err != nil {
			global.FITNESS_LOG.Debug("用户登录失败", zap.Any("err", err))
			global.FITNESS_CACHE.Increment(key, 1)
			response.ErrorWithMessage("用户名或密码错误", c)
			return
		}
		if user.Enable != 1 {
			global.FITNESS_LOG.Debug("用户被禁用", zap.Any("user", user))
			global.FITNESS_CACHE.Increment(key, 1)
			response.ErrorWithMessage("用户被禁用", c)
			return
		}

	}
}

// Captcha
// @Tags Base
// @Summary 生成验证码
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=systemRes.SysCaptchaResponse,msg=string} "生成验证码,返回包括随机数id,base64,验证码长度,是否开启验证码"
// @Router /base/captcha [post]
func (b *BaseApi) Captcha(c *gin.Context) {

}
