package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	sneakerq "github.com/ukurysheva/sneaker-q"
)

func (h *Handler) getShopModels(c *gin.Context) {
	shop := c.Param("shop")
	list, err := h.services.Model.GetShopModels(shop)

	fmt.Println(list)
	if err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, err.Error())
		fmt.Println(err.Error())
		return
	}

	c.JSON(http.StatusOK, list)

}

func (h *Handler) searchModels(c *gin.Context) {
	var input sneakerq.SearchParams

	if err := c.BindJSON(&input); err != nil {
		sendErrorResponse(c, http.StatusBadRequest, err.Error())
		fmt.Printf("error while binding")
		fmt.Println(err.Error())
		// newErrorResponse
		return
	}
	list, err := h.services.Model.GetModelsByParams(input)

	fmt.Println(list)
	if err != nil {
		fmt.Println(err.Error())
		sendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, list)
}

func (h *Handler) getModelById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if id < 1 || err != nil {
		sendErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	model, err := h.services.Model.GetModelById(id)
	if err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, err.Error())
		fmt.Printf("error in handler - get model by id")
		fmt.Println(err.Error())
	}
	c.JSON(http.StatusOK, model)
}
