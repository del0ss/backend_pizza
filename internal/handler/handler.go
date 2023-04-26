package handler

import (
	"github.com/gin-gonic/gin"
	"smth/internal/store"
	"smth/internal/store/sqlstore"
	"smth/pkg/auth"
)

type Handler struct {
	store        store.Store
	tokenManager auth.TokenManager
	cors         gin.HandlerFunc
}

func New(store *sqlstore.Store, tokenManager auth.TokenManager, cors gin.HandlerFunc) *Handler {
	return &Handler{
		store:        store,
		tokenManager: tokenManager,
		cors:         cors,
	}
}

func (h *Handler) ConfigureRouter() *gin.Engine {
	router := gin.Default()
	router.Use(h.cors)
	router.GET("/", h.helloPage())
	router.POST("/sing-up", h.registerUser())
	router.POST("/sing-in", h.loginUser())

	authGroup := router.Group("/auth", h.userIdentity)
	{
		authGroup.GET("/pepepe", h.handlePe())
	}
	postsGroup := router.Group("/posts")
	{
		postsGroup.GET("/", h.getPosts())
		postsGroup.POST("/create", h.handlerCreatePost())
		postsGroup.GET("/:id", h.getPost())
		postsGroup.DELETE("/:id", h.DeletePost())
	}
	return router
}
