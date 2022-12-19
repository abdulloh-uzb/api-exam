package v1

import (
	"api-exam/genproto/reyting"
	l "api-exam/pkg/logger"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// @Summary create post with info
// @Description this func creates post
// @Tags post
// @Accept json
// @Produce json
// @Security     BearerAuth
// @Param post body reyting.Ranking true "Post"
// @Router /v1/create-reyting [post]
func (h *handlerV1) CreateReyting(c *gin.Context) {
	var (
		reyting reyting.Ranking
	)
	fmt.Println(c)
	err := c.ShouldBindJSON(&reyting)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", l.Error(err))
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()
	response, err := h.serviceManager.ReytingService().CreateRanking(ctx, &reyting)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create reyting", l.Error(err))
		return
	}

	c.JSON(http.StatusCreated, response)
}
