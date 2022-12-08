package v1

import (
	"api-exam/api/models"
	"api-exam/genproto/post"
	l "api-exam/pkg/logger"
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
)

// @Summary create post with info
// @Description this func creates post
// @Tags post
// @Accept json
// @Produce json
// @Security     BearerAuth
// @Param post body post.PostReq true "Post"
// @Success 200 {object} post.Post
// @Router /v1/create-post [post]
func (h *handlerV1) CreatePost(c *gin.Context) {
	var (
		post post.PostReq
	)
	fmt.Println(c)
	err := c.ShouldBindJSON(&post)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", l.Error(err))
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()
	response, err := h.serviceManager.PostService().CreatePost(ctx, &post)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create product", l.Error(err))
		return
	}

	c.JSON(http.StatusCreated, response)
}

// @Summary update post
// @Description this func update post
// @Tags post
// @Accept json
// @Produce json
// @Security     BearerAuth
// @Param update body models.UpdatePost true "Post"
// @Success 200 {object} models.UpdatePostResp
// @Router /v1/update-post [put]
func (h *handlerV1) UpdatePost(c *gin.Context) {
	var (
		post        post.Post
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true
	err := c.ShouldBindJSON(&post)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", l.Error(err))
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()
	response, err := h.serviceManager.PostService().UpdatePost(ctx, &post)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create product", l.Error(err))
		return
	}
	res := &models.UpdatePostResp{
		Id:          int(response.Id),
		Name:        response.Name,
		Description: response.Description,
		CreatedAt:   response.CreatedAt,
		UpdatedAt:   response.UpdatedAt,
		CustomerId:  int(response.CustomerId),
	}
	c.JSON(http.StatusCreated, res)
}

// @Summary get post
// @Description this func get post by id
// @Tags post
// @Security     BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Success 200 "success" {object} post.Post
// @Router /v1/get-post/{id} [get]
func (h *handlerV1) GetPost(c *gin.Context) {
	s_id := c.Param("id")
	id, err := strconv.ParseInt(s_id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("error while id parseint", l.Error(err))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.PostService().GetPost(ctx, &post.Id{
		Id: id,
	})

	fmt.Println(response)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get post", l.Error(err))
		return
	}

	c.JSON(http.StatusCreated, response)
}

// @Summary get post list
// @Description this func get list of posts
// @Tags post
// @Security     BearerAuth
// @Accept json
// @Produce json
// @Success 200 "success" {object} post.Posts
// @Router /v1/list-post [get]
func (h *handlerV1) GetPostList(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.PostService().ListPost(ctx, &post.Empty{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get list of posts", l.Error(err))
		return
	}

	c.JSON(http.StatusCreated, response)
}

// @Summary delete post
// @Description this func delete post by post id
// @Tags post
// @Security     BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Success 200 "success"
// @Router /v1/delete-post/{id} [DELETE]
func (h *handlerV1) DeletePostById(c *gin.Context) {
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

	response, err := h.serviceManager.PostService().DeletePost(ctx, &post.Id{Id: id})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get post", l.Error(err))
		return
	}

	c.JSON(http.StatusCreated, response)
}

// @Summary get post list
// @Description this func get list of customers by post id
// @Tags post
// @Security     BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Success 200 "success" {object} post.Posts
// @Router /v1/getpostlist/{id} [get]
func (h *handlerV1) GetPostByCustomerId(c *gin.Context) {
	s_id := c.Param("id")
	id, err := strconv.ParseInt(s_id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("error while id parseint", l.Error(err))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.PostService().GetPostByCustomerId(ctx, &post.Id{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get list of posts", l.Error(err))
		return
	}

	c.JSON(http.StatusCreated, response)
}
