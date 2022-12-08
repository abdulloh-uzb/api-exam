package v1

import (
	"api-exam/api/models"
	"api-exam/genproto/customer"
	l "api-exam/pkg/logger"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// @Summary admin login
// @Description this func login admin
// @Tags admin
// @Accept json
// @Produce json
// @Param        email  path string true "email"
// @Param        password   path string true "password"
// @Success 200 {object} customer.Admin
// @Router      /v1/admin-login/{email}/{password} [get]
func (h *handlerV1) AdminLogin(c *gin.Context) {
	var (
		email    = c.Param("email")
		password = c.Param("password")
	)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	res, err := h.serviceManager.CustomerService().GetAdminByEmail(ctx, &customer.GetAdminReq{Email: email, Password: password})

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err,
		})
		h.log.Error("error while get admin info", l.Error(err))
		return
	}

	h.jwthandler.Iss = "admin"
	h.jwthandler.Role = "admin"
	h.jwthandler.Aud = []string{"exam-app"}
	h.jwthandler.SigninKey = h.cfg.SignKey
	h.jwthandler.Log = h.log
	tokens, err := h.jwthandler.GenerateAuthJWT()
	accessToken := tokens[0]

	response := &models.Admin{
		Email: res.Email,
	}

	response.AccessToken = accessToken
	c.JSON(http.StatusOK, response)
}
