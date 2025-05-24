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
