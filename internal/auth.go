package internal

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckAuth(ctx *gin.Context) error {
	if login, pass, _ := ctx.Request.BasicAuth(); pass != Config.BasicPassword || login != Config.BasicLogin {
		return fmt.Errorf("%v", http.StatusUnauthorized)
	} else {
		return nil
	}
}
