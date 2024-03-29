package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"smth/internal/model"
	"strconv"
)

func (h *Handler) getPizza() gin.HandlerFunc {
	return func(c *gin.Context) {
		sortType := c.Query("sortType")
		page, _ := strconv.Atoi(c.Query("page"))
		categoryID, _ := strconv.Atoi(c.Query("category"))
		//_, ok := c.Get(userContext)
		//if !ok  {
		//	newErrorMessage(c, http.StatusUnauthorized, "invalid header")
		//	return
		//}
		p, err := h.store.Pizza().GetPizza(sortType, categoryID, page)
		if err != nil {
			newErrorMessage(c, http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, p)
	}
}

func (h *Handler) getCategories() gin.HandlerFunc {
	return func(c *gin.Context) {
		//_, ok := c.Get(userContext)
		//if !ok {
		//	newErrorMessage(c, http.StatusUnauthorized, "invalid header")
		//	return
		//}
		p, err := h.store.Pizza().GetCategories()
		if err != nil {
			newErrorMessage(c, http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, p)
	}
}

func (h *Handler) getCategoryById() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		p, err := h.store.Pizza().GetCategoryById(id)
		if err != nil {
			newErrorMessage(c, http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, p)
	}
}

func (h *Handler) getPizzaById() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		p, err := h.store.Pizza().GetPizzaById(id)
		if err != nil {
			newErrorMessage(c, http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, p)
	}
}

func (h *Handler) deletePizza() gin.HandlerFunc {
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

func (h *Handler) createPizza() gin.HandlerFunc {

	return func(c *gin.Context) {
		var p model.Pizza
		if err := c.BindJSON(&p); err != nil {
			newErrorMessage(c, http.StatusUnauthorized, err.Error())
			return
		}

		//_, ok := c.Get(userContext)
		//if ok == false {
		//	newErrorMessage(c, http.StatusUnauthorized, "invalid header")
		//	return
		//}

		id, err := h.store.Pizza().CreatePizza(p)
		if err != nil {
			logrus.Error(err)
			newErrorMessage(c, http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, id)
	}
}
func (h *Handler) getCountPage() gin.HandlerFunc {
	return func(c *gin.Context) {
		p, err := h.store.Pizza().GetCountPage()
		if err != nil {
			newErrorMessage(c, http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, p)
	}
}
