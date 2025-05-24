package routers

import (
	"demo-gin-mvc-verification-code/controllers/verify"

	"github.com/gin-gonic/gin"
)

func RoutersVerify(r *gin.Engine) {
	r.GET("/get-verify", verify.VerifyControllers{}.Captcha)

	r.POST("/check-verify", verify.VerifyControllers{}.CheckCaptcha)
}
