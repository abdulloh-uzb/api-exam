package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"api-exam/api/models"
	"api-exam/genproto/customer"
	l "api-exam/pkg/logger"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"google.golang.org/protobuf/encoding/protojson"
)

// Verify customer
// @Summary      Verify customer
// @Description  Verifys customer
// @Tags         Register
// @Accept       json
// @Produce      json
// @Param        email  path string true "email"
// @Param        code   path string true "code"
// @Success      200  {object}  models.Verify
// @Router      /v1/verify/{email}/{code} [get]
func (h *handlerV1) Verification(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	var (
		code  = c.Param("code")
		email = c.Param("email")
	)

	redisData, err := h.redis.Get(email)
	fmt.Println(email, code)
	if err != nil {
		c.JSON(http.StatusGatewayTimeout, gin.H{
			"info":  "Your time has expired",
			"error": err.Error(),
		})
		h.log.Error("Error while getting user from redis", l.Any("redis", err))
		return
	}
	customerRedis := cast.ToString(redisData)
	customerRedisString := models.Customer{}
	err = json.Unmarshal([]byte(customerRedis), &customerRedisString)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("Error while unmarshaling from json to customer string", l.Any("json", err))
		return
	}

	if customerRedisString.Code != code {
		fmt.Println(customerRedisString.Code)
		c.JSON(http.StatusConflict, gin.H{
			"info": "Wrong code",
		})
		return
	}

	body := customer.CustomerRequest{
		FirstName:    customerRedisString.FirstName,
		LastName:     customerRedisString.LastName,
		Email:        customerRedisString.Email,
		Bio:          customerRedisString.Bio,
		PhoneNumber:  customerRedisString.PhoneNumber,
		Password:     customerRedisString.Password,
		RefreshToken: customerRedisString.Refreshtoken,
		Addresses:    customerRedisString.Addresses,
	}
	h.jwthandler.Iss = "user"
	h.jwthandler.Role = "authorized"
	h.jwthandler.Aud = []string{"exam"}
	h.jwthandler.SigninKey = h.cfg.SignKey
	h.jwthandler.Log = h.log
	tokens, err := h.jwthandler.GenerateAuthJWT()
	accessToken := tokens[0]
	refreshToken := tokens[1]
	if err != nil {
		h.log.Error("error occured while generating tokens")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "something went wrong,please try again",
		})
		return
	}

	body.RefreshToken = refreshToken
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()
	res, err := h.serviceManager.CustomerService().Create(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("Error while creating customer", l.Error(err))
		return
	}

	response := models.Verify{
		Id:        int(res.Id),
		FirstName: res.FirstName,
		LastName:  res.LastName,
		Email:     res.Email,
	}

	response.AccessToken = accessToken
	response.RefreshToken = refreshToken

	c.JSON(http.StatusOK, response)

}
