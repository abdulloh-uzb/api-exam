package v1

import (
	"api-exam/genproto/customer"
	l "api-exam/pkg/logger"
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
)

// @Summary get customer with info
// @Description this func get customer
// @Tags customer
// @Security     BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Success 200 "success"
// @Router /v1/get-customer/{id} [get]
func (h *handlerV1) GetCustomerById(c *gin.Context) {

	string_id := c.Param("id")
	id, err := strconv.ParseInt(string_id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("error while id parseint", l.Error(err))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.CustomerService().GetCustomer(ctx, &customer.CustomerId{
		Id: id,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create product", l.Error(err))
		return
	}

	c.JSON(http.StatusCreated, response)
}

// @Summary get customer list
// @Description this func get list of customers
// @Tags customer
// @Security     BearerAuth
// @Accept json
// @Produce json
// @Success 200 "success"
// @Router /v1/list-customer [get]
func (h *handlerV1) GetCustomerList(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.CustomerService().GetCustomerList(ctx, &customer.Empty{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get list of customer", l.Error(err))
		return
	}

	c.JSON(http.StatusCreated, response)
}

// @Summary delete customer
// @Description this func delete customer by customer id
// @Tags customer
// @Security     BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Success 200 "success"
// @Router /v1/delete-customer/{id} [delete]
func (h *handlerV1) DeleteCustomerById(c *gin.Context) {
	string_id := c.Param("id")
	id, err := strconv.ParseInt(string_id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("error while id parseint", l.Error(err))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.CustomerService().DeleteCustomer(ctx, &customer.CustomerId{
		Id: id,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get post", l.Error(err))
		return
	}

	c.JSON(http.StatusCreated, response)
}

// @Summary update customer
// @Description this func update customer
// @Tags customer
// @Security     BearerAuth
// @Accept json
// @Produce json
// @Param customer body customer.Customer true "Customer"
// @Success 200 {object} customer.Customer
// @Router /v1/update-customer [put]
func (h *handlerV1) UpdateCustomer(c *gin.Context) {
	var (
		body        customer.Customer
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", l.Error(err))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.CustomerService().Create(ctx, &customer.CustomerRequest{
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Bio:       body.Bio,
		Addresses: body.Addresses,
		Email:     body.Email,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create product", l.Error(err))
		return
	}

	c.JSON(http.StatusCreated, response)
}
