package sysservice

import (
	"errors"
	"fmt"
	"github.com/lazybearlee/yuedong-fitness/global"
	"github.com/lazybearlee/yuedong-fitness/utils"
	"gopkg.in/gomail.v2"
)

type EmailService struct{}

// SendValidatorMessage build validator message
func (e *EmailService) SendValidatorMessage(to string) error {
	// 首先查验验证码是否过期，拒绝重复发送验证码
	if _, ok := global.FitnessCache.Get(to); ok {
		return errors.New("验证码已发送，请查看邮箱")
	}
	// 生成6位随机验证码
	code := utils.GenerateVerificationCode()
	// 生成验证码的 body
	body := fmt.Sprintf(`
	  <div>
		<div>
		  尊敬的%s，您好！
		</div>
		<div style="padding: 8px 40px 8px 50px;">
		  <p>本次验证码为<u><strong>%s</strong></u>，为了保证账号安全，验证码有效期为5分钟。请确认为本人操作，切勿向他人泄露，感谢您的理解与使用。</p>
		</div>
		<div>
		  <p>此邮箱为系统邮箱，请勿回复。</p>
		</div>
	  </div>
  	`, to, code)

	// 发送邮件
	m := gomail.NewMessage()
	m.SetHeader("From", global.FitnessConfig.Email.From)
	m.SetHeader("To", to)
	m.SetHeader("Subject", "YueDong Fitness 验证码")
	m.SetBody("text/html", body)

	d := gomail.NewDialer(global.FitnessConfig.Email.Host,
		global.FitnessConfig.Email.Port,
		global.FitnessConfig.Email.From,
		global.FitnessConfig.Email.Secret)

	if err := d.DialAndSend(m); err != nil {
		return err
	}
	// 将验证码存入缓存
	dr, err := utils.ParseDuration(global.FitnessConfig.Email.CodeTimeOut)
	if err != nil {
		return err
	}
	global.FitnessCache.Set(to, code, dr)
	return nil
}
