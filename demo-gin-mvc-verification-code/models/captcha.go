package models

import (
	"fmt"
	"image/color"

	"github.com/mojocn/base64Captcha"
)

// 保存在内存中
var store = base64Captcha.DefaultMemStore

// 获取验证码
func MakeCaptcha() (string, string, error) {
	var driver base64Captcha.Driver

	driverString := base64Captcha.DriverString{
		Height:          40,
		Width:           100,
		NoiseCount:      0,
		ShowLineOptions: 2 | 4,
		Length:          4,
		Source:          "1234567890qwertyuioplkjhgfdsazxcvbnm",
		BgColor: &color.RGBA{
			R: 3,
			G: 102,
			B: 214,
			A: 125,
		},
		Fonts: []string{"wqy-microhei.ttc"},
	}

	driver = driverString.ConvertFonts()

	c := base64Captcha.NewCaptcha(driver, store)
	id, b64s, _, err := c.Generate()

	return id, b64s, err
}

// 验证验证码
func VerifyCaptcha(id string, VerifyValue string) bool {

	fmt.Println(id, VerifyValue)

	// 但三个参数代表 验证码ID，验证码，是否清空验证码
	if store.Verify(id, VerifyValue, true) {
		return true
	} else {
		return false
	}
}
