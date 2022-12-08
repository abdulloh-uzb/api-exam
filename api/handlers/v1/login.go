package v1

import (
	"api-exam/api/models"
	"api-exam/pkg/etc"
	"api-exam/pkg/logger"
	"context"
	"fmt"
	"net/http"
	"time"

	"api-exam/genproto/customer"

	"github.com/gin-gonic/gin"
)

// Login user
// @Summary      Login customer
// @Tags         Login
// @Accept       json
// @Produce      json
// @Param        email  path string true "email"
// @Param        password   path string true "password"
// @Success         200                   {object}  models.Login
// Failure         500                   {object}  models.Error
// @Router      /v1/login/{email}/{password} [get]
func (h *handlerV1) Login(c *gin.Context) {
	fmt.Println(1)
	var (
		password = c.Param("password")
		email    = c.Param("email")
	)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	res, err := h.serviceManager.CustomerService().GetByEmail(ctx, &customer.LoginReq{
		Email: email,
	})
	fmt.Println(res)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Bu email bilan malumot topilmadi",
		})
		h.log.Error("Error while getting customer by email", logger.Error(err))
		return
	}

	if !etc.CheckPasswordHash(password, res.Password) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Password or Email is wrong",
		})
		return
	}
	h.jwthandler.Iss = "user"
	h.jwthandler.Role = "authorized"
	h.jwthandler.Aud = []string{"exam-app"}
	h.jwthandler.SigninKey = h.cfg.SignKey
	h.jwthandler.Log = h.log
	tokens, err := h.jwthandler.GenerateAuthJWT()
	accessesToken := tokens[0]
	refreshToken := tokens[1]

	if err != nil {
		h.log.Error("error occured while generating tokens")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "something went wrong,please try again",
		})
		return
	}
	response := &models.Login{
		Email:     res.Email,
		FirstName: res.FirstName,
		LastName:  res.LastName,
		Password:  res.Password,
	}

	response.RefreshToken = refreshToken
	response.AccessToken = accessesToken
	response.Password = ""
	c.JSON(http.StatusOK, response)
}
