package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"smth/internal/model"
	"strconv"
)

func (h *Handler) getPosts() gin.HandlerFunc {
	return func(c *gin.Context) {
		//_, ok := c.Get(userContext)
		//if ok == false {
		//	newErrorMessage(c, http.StatusUnauthorized, "invalid header")
		//	return
		//}
		p, err := h.store.Post().GetPosts()
		if err != nil {
			newErrorMessage(c, http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, p)
	}
}

func (h *Handler) getPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		p, err := h.store.Post().GetPost(id)
		if err != nil {
			newErrorMessage(c, http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, p)
	}
}

func (h *Handler) DeletePost() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		err := h.store.Post().DeletePost(id)
		if err != nil {
			logrus.Error(err)
			newErrorMessage(c, http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, "OK")
	}
}

func (h *Handler) createPost() gin.HandlerFunc {

	return func(c *gin.Context) {
		var p model.Post
		if err := c.BindJSON(&p); err != nil {
			newErrorMessage(c, http.StatusUnauthorized, err.Error())
			return
		}

		//_, ok := c.Get(userContext)
		//if ok == false {
		//	newErrorMessage(c, http.StatusUnauthorized, "invalid header")
		//	return
		//}

		id, err := h.store.Post().CreatePost(p, 1)
		if err != nil {
			logrus.Error(err)
			newErrorMessage(c, http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, id)
	}
}
