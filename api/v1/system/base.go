package sysapi

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lazybearlee/yuedong-fitness/global"
	"github.com/lazybearlee/yuedong-fitness/model/common/response"
	sysmodel "github.com/lazybearlee/yuedong-fitness/model/system"
	sysrequest "github.com/lazybearlee/yuedong-fitness/model/system/request"
	sysresponse "github.com/lazybearlee/yuedong-fitness/model/system/response"
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
// @Param data body sysrequest.LoginReq true "用户名, 密码, 验证码"
// @Success 200 {object} response.Response{data=sysresponse.LoginResponse,msg=string} "登录成功"
// @Router /base/login [post]
func (b *BaseApi) Login(c *gin.Context) {
	var login sysrequest.LoginReq
	err := c.ShouldBindJSON(&login)
	if err != nil {
		response.ErrorWithMessage("参数绑定失败", c)
	}
	// 是否开启验证码
	isCaptchaOpen := global.FitnessConfig.Captcha.OpenCaptcha
	captchaTimeout := global.FitnessConfig.Captcha.OpenCaptchaTimeOut
	key := c.ClientIP()
	v, ok := global.FitnessCache.Get(key)
	if !ok {
		global.FitnessCache.Set(key, 1, time.Second*time.Duration(captchaTimeout)) // 设置缓存超时时间
	}
	var checkCaptcha bool = (isCaptchaOpen != 0 && isCaptchaOpen >= utils.InterfaceToInt(v)) || (login.CaptchaId != "" && login.Captcha != "" && captchaStore.Verify(login.CaptchaId, login.Captcha, true))
	if checkCaptcha {
		// 如果不使用验证码或者验证码验证通过
		// 登陆
		u := sysmodel.SysUser{Username: login.Username, Password: login.Password}
		user, err := userService.UserLogin(u)
		if err != nil {
			e := fmt.Sprintf("登录失败: %v", err)
			global.FitnessLog.Debug(e)
			global.FitnessCache.Increment(key, 1)
			response.ErrorWithMessage("用户名或密码错误", c)
			return
		}
		if user.Enable != 1 {
			global.FitnessLog.Debug("用户还未激活或被禁用", zap.Any("user", user))
			global.FitnessCache.Increment(key, 1)
			response.ErrorWithMessage("用户还未激活或被禁用", c)
			return
		}
		// 生成token
		token, claims, err := utils.NewLoginToken(user)
		if err != nil {
			global.FitnessLog.Debug("生成token失败", zap.Any("err", err))
			response.ErrorWithMessage("生成token失败", c)
			return
		}
		// 如果不使用多点
		if !global.FitnessConfig.System.UseMultipoint {
			utils.SetToken(c, token, int(claims.ExpiresAt.Unix()-time.Now().Unix()))
			response.SuccessWithDetailed(sysresponse.LoginResponse{
				User:      *user,
				Token:     token,
				ExpiresAt: claims.ExpiresAt.Unix() * 1000,
			}, "登录成功", c)
			return
		}
		// 否则，需要使用Redis
		// TODO: 使用Redis实现多服务器部署JWT登陆
		return
	}
	global.FitnessCache.Increment(key, 1)
	response.ErrorWithMessage("验证码错误", c)
}

// AdminRegister
// @Tags     SysUser
// @Summary  用户注册账号
// @Produce  application/json
// @Param    data  body      sysrequest.RegisterReq                                            true  "用户名, 昵称, 密码, 角色ID"
// @Success  200   {object}  response.Response{data=sysresponse.UserResponse,msg=string}  "用户注册账号,返回包括用户信息"
// @Router   /user/admin_register [post]
func (b *BaseApi) AdminRegister(c *gin.Context) {
	var register sysrequest.RegisterReq
	err := c.ShouldBindJSON(&register)
	if err != nil {
		response.ErrorWithMessage("参数绑定失败", c)
		return
	}
	var auths []*sysmodel.SysAuthority
	for _, v := range register.AuthorityIds {
		auths = append(auths, &sysmodel.SysAuthority{AuthorityId: v})
	}
	u := sysmodel.SysUser{
		Username:    register.Username,
		NickName:    register.NickName,
		Password:    register.Password,
		HeaderImg:   register.HeaderImg,
		AuthorityId: register.AuthorityId,
		Authorities: auths,
		Phone:       register.Phone,
		Email:       register.Email,
		Enable:      register.Enable,
	}
	ur, err := userService.UserRegister(u)
	if err != nil {
		e := fmt.Sprintf("用户注册失败: %v", err)
		global.FitnessLog.Debug(e)
		response.ErrorWithMessage(e, c)
		return
	}
	response.SuccessWithDetailed(sysresponse.UserResponse{User: ur}, "注册成功", c)
}

// Register 一般用户注册
// @Tags Base
// @Summary 用户注册
// @Produce application/json
// @Param data body sysrequest.RegisterReq true "用户名, 昵称, 密码, 头像, 手机号, 邮箱"
// @Success 200 {object} response.Response{data=sysresponse.UserResponse,msg=string} "用户注册,返回包括用户信息"
// @Router /base/register [post]
func (b *BaseApi) Register(c *gin.Context) {
	var register sysrequest.RegisterReqWithCode
	err := c.ShouldBindJSON(&register)
	if err != nil {
		response.ErrorWithMessage("参数绑定失败", c)
		return
	}
	// 校验Email格式
	if !utils.EmailFormatCheck(register.Email) {
		response.ErrorWithMessage("邮箱格式不正确", c)
		return
	}
	auths := []*sysmodel.SysAuthority{}
	auths = append(auths, &sysmodel.SysAuthority{AuthorityId: global.CommonUser})
	u := sysmodel.SysUser{
		Username:    register.Username,
		NickName:    register.Username,
		Password:    register.Password,
		Phone:       register.Phone,
		Email:       register.Email,
		AuthorityId: global.CommonUser, // 默认为普通用户
		Authorities: auths,
		Enable:      2, // 默认为2，需要发送邮件或短信进行激活
	}
	ur, err := userService.UserRegister(u)
	if err != nil {
		e := fmt.Sprintf("用户注册失败: %v", err)
		global.FitnessLog.Debug(e)
		response.ErrorWithMessage(e, c)
		return
	}
	// 发送激活邮件
	err = emailService.SendValidatorMessage(register.Email)
	if err != nil {
		e := fmt.Sprintf("发送验证码失败: %v", err)
		global.FitnessLog.Debug(e)
		response.ErrorWithMessage(e, c)
		return
	}
	response.SuccessWithDetailed(sysresponse.UserResponse{User: ur}, "请前往邮箱激活账号", c)
}

// VerificationCode
// @Tags Base
// @Summary 发送验证码
// @Produce application/json
// @Param data body sysrequest.VerificationCodeReq true "邮箱"
// @Success 200 {object} response.Response{data=string,msg=string} "发送验证码"
// @Router /base/verification_code [post]
func (b *BaseApi) VerificationCode(c *gin.Context) {
	var v sysrequest.VerificationCodeReq
	err := c.ShouldBindJSON(&v)
	if err != nil {
		response.ErrorWithMessage("参数绑定失败", c)
		return
	}
	// 校验Email格式
	if !utils.EmailFormatCheck(v.Email) {
		response.ErrorWithMessage("邮箱格式不正确", c)
		return
	}
	err = emailService.SendValidatorMessage(v.Email)
	if err != nil {
		e := fmt.Sprintf("发送验证码失败: %v", err)
		global.FitnessLog.Debug(e)
		response.ErrorWithMessage(e, c)
		return
	}
	response.SuccessWithMessage("发送验证码成功", c)
}

// RegisterWithCode
// @Tags Base
// @Summary 验证邮箱验证码，完成注册
// @Produce application/json
// @Param data body sysrequest.RegisterReqWithCode true "邮箱, 验证码"
// @Success 200 {object} response.Response{data=string,msg=string} "验证邮箱验证码"
// @Router /base/register_with_code [post]
func (b *BaseApi) RegisterWithCode(c *gin.Context) {
	var v sysrequest.RegisterReqWithCode
	err := c.ShouldBindJSON(&v)
	if err != nil {
		response.ErrorWithMessage("参数绑定失败", c)
		return
	}
	// 校验Email格式
	if !utils.EmailFormatCheck(v.Email) {
		response.ErrorWithMessage("邮箱格式不正确", c)
		return
	}
	// 验证验证码
	code, ok := global.FitnessCache.Get(v.Email)
	if !ok {
		response.ErrorWithMessage("验证码无效", c)
		return
	}
	if code != v.Code {
		response.ErrorWithMessage("验证码错误", c)
		return
	}
	// 修改用户状态
	var auths []*sysmodel.SysAuthority
	auths = append(auths, &sysmodel.SysAuthority{AuthorityId: global.CommonUser})
	u := sysmodel.SysUser{
		Username:    v.Username,
		NickName:    v.Username,
		Password:    v.Password,
		Phone:       v.Phone,
		Email:       v.Email,
		AuthorityId: global.CommonUser, // 默认为普通用户
		Authorities: auths,
		Enable:      1, // 已激活，正常用户
	}
	_, err = userService.UserRegister(u)
	if err != nil {
		e := fmt.Sprintf("用户注册失败: %v", err)
		global.FitnessLog.Debug(e)
		response.ErrorWithMessage(e, c)
		return
	}
	response.SuccessWithMessage("验证成功", c)
}

// Captcha
// @Tags Base
// @Summary 生成验证码
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=sysresponse.CaptchaResponse,msg=string} "生成验证码,返回包括随机数id,base64,验证码长度,是否开启验证码"
// @Router /base/captcha [post]
func (b *BaseApi) Captcha(c *gin.Context) {
	// 判断验证码是否开启
	openCaptcha := global.FitnessConfig.Captcha.OpenCaptcha               // 是否开启防爆次数
	openCaptchaTimeOut := global.FitnessConfig.Captcha.OpenCaptchaTimeOut // 缓存超时时间
	key := c.ClientIP()
	v, ok := global.FitnessCache.Get(key)
	if !ok {
		global.FitnessCache.Set(key, 1, time.Second*time.Duration(openCaptchaTimeOut))
	}
	var oc bool = openCaptcha == 0 || openCaptcha < utils.InterfaceToInt(v)
	// 生成默认数字的driver
	driver := base64Captcha.NewDriverDigit(global.FitnessConfig.Captcha.ImgHeight, global.FitnessConfig.Captcha.ImgWidth, global.FitnessConfig.Captcha.KeyLong, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, captchaStore)
	id, b64s, _, err := cp.Generate()
	if err != nil {
		global.FitnessLog.Error("验证码获取失败", zap.Error(err))
		response.ErrorWithMessage("验证码获取失败", c)
		return
	}
	response.SuccessWithDetailed(sysresponse.CaptchaResponse{
		CaptchaId:     id,
		PicPath:       b64s,
		CaptchaLength: global.FitnessConfig.Captcha.KeyLong,
		OpenCaptcha:   oc,
	}, "验证码获取成功", c)
}
