package v1

import (
	"api-exam/api/models"
	"api-exam/genproto/customer"
	"api-exam/pkg/email"
	l "api-exam/pkg/logger"
	"api-exam/pkg/utils"
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"api-exam/pkg/etc"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"golang.org/x/crypto/bcrypt"
)

// Register
// @Summary      Register
// @Description  Registration
// @Tags         Register
// @Accept       json
// @Produce      json
// @Param        customer   body models.CustomerReq     true  "Customers"
// @Success      200  {object}  customer.Customer
// @Router       /v1/register [post]
func (h *handlerV1) Register(c *gin.Context) {
	var (
		body models.Customer
	)
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("error while bind json", l.Error(err))
		return
	}

	err = utils.IsValidMail(body.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid email address",
		})
		return
	}

	// emailni kichkinga harfga o'tqizyapti va spacelarni olib tashlayapti
	body.Email = strings.TrimSpace(body.Email)
	body.Email = strings.ToLower(body.Email)

	ctx, cancel := context.WithTimeout(context.Background(),
		time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	// kiriltgan email oldin ham royxatdan otganmi yoqmi shuni tekshiryapti
	existsEmail, err := h.serviceManager.CustomerService().CheckField(ctx, &customer.CheckFieldReq{
		Field: "email",
		Value: body.Email,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed check email uniques")
		return
	}
	if existsEmail.Exists {
		c.JSON(http.StatusConflict, gin.H{
			"info": "Bunaqa email mavjud",
		})
		return
	}

	// emailni redisdan tekshiryapti
	exists, err := h.redis.Exists(body.Email)

	if cast.ToInt(exists) == 1 {
		c.JSON(http.StatusConflict, gin.H{
			"error": "bunaqa email mavjud",
		})
		return
	}

	hashPass, err := bcrypt.GenerateFromPassword([]byte(body.Password), 14)

	if err != nil {
		h.log.Error("error while hashing password", l.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "something went wrong",
		})
		return
	}
	code := etc.GenerateCode(6)
	body.Password = string(hashPass)
	add := []*models.Address{}
	for _, i := range body.Addresses {
		a := models.Address{
			District: i.District,
			Street:   i.Street,
		}
		add = append(add, &a)
	}
	ref := &models.Customer{
		FirstName:   body.FirstName,
		LastName:    body.LastName,
		Email:       body.Email,
		Username:    body.Username,
		Bio:         body.Bio,
		PhoneNumber: body.PhoneNumber,
		Password:    body.Password,
		Code:        code,
		Addresses:   body.Addresses,
	}
	msg := "Subject: Exam email verification\n Your verification code: " + ref.Code
	err = email.SendMail([]string{ref.Email}, []byte(msg))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
			"info":  "Your Email is not valid, Please recheck it",
		})
		return
	}

	customerJSON, err := json.Marshal(ref)

	if err != nil {
		h.log.Error("error while marshaling customer,inorder to insert it to redis", l.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error while creating customer",
		})
		return
	}

	if err = h.redis.SetWithTTL(string(ref.Email), string(customerJSON), 120); err != nil {
		h.log.Error("error while inserting new customer into redis")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "something went wrong, please try again",
		})
		return
	}

	c.JSON(http.StatusOK, "Emailingizga kod yuborildi ")

}
