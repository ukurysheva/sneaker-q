package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	sneakerq "github.com/ukurysheva/sneaker-q"
)

func (h *Handler) getShopModels(c *gin.Context) {
	shop := c.Param("shop")
	fmt.Println(shop)
	list, err := h.services.ModelList.GetShopModels(shop)

	fmt.Println(list)
	if err != nil {
		fmt.Println(err.Error())
		// newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, list)

}

func (h *Handler) searchModels(c *gin.Context) {
	var input sneakerq.SearchParams

	if err := c.BindJSON(&input); err != nil {
		fmt.Printf("error while binding")
		fmt.Println(err.Error())
		// newErrorResponse
		return
	}
	fmt.Println(input)
	list, err := h.services.ModelList.GetModelsByParams(input)

	fmt.Println(list)
	if err != nil {
		fmt.Println(err.Error())
		// newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, list)
}
