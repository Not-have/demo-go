package verify

import (
	"demo-gin-mvc-verification-code/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type VerifyControllers struct{}

func (con VerifyControllers) Captcha(c *gin.Context) {
	id, b64s, err := models.MakeCaptcha()

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"CaptchaId":    id,
		"CaptchaImage": b64s,
	})
}

func (con VerifyControllers) CheckCaptcha(c *gin.Context) {

	id := c.PostForm("CaptchaId")
	value := c.PostForm("CaptchaValue")

	if models.VerifyCaptcha(id, value) {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "验证码正确",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "验证码错误",
		})
	}
}
